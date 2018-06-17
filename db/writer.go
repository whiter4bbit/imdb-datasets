package db

import (
	"encoding/binary"
	"errors"
	bolt "github.com/coreos/bbolt"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"log"
	"reflect"
	"time"
)

type Writer interface {
	WriteMovie(*datasets.MovieIdentity) error
	WriteEpisode(*datasets.EpisodeIdentity) error
	WriteAttribute(datasets.Hash, datasets.Attribute) error
}

type HashAndAttribute struct {
	Hash datasets.Hash
	Attr datasets.Attribute
}

type NoOpWriter struct {
	movies     int
	episodes   int
	attributes int
}

func (n *NoOpWriter) WriteMovie(*datasets.MovieIdentity) error {
	n.movies += 1
	return nil
}

func (n *NoOpWriter) WriteEpisode(*datasets.EpisodeIdentity) error {
	n.episodes += 1
	return nil
}

func (n *NoOpWriter) WriteAttribute(datasets.Hash, datasets.Attribute) error {
	n.attributes += 1
	return nil
}

func (n *NoOpWriter) ReportStats() {
	log.Printf("Movies:\t%d\nEpisodes:\t%d\nAttributes:\t%d", n.movies, n.episodes, n.attributes)
}

type MockWriter struct {
	Attrs    []*HashAndAttribute
	Movies   []*datasets.MovieIdentity
	Episodes []*datasets.EpisodeIdentity
}

func (w *MockWriter) WriteMovie(m *datasets.MovieIdentity) error {
	if w.findMovie(m.MakeHash()) {
		return nil
	}

	w.Movies = append(w.Movies, m)
	return nil
}

func (w *MockWriter) findMovie(hash datasets.Hash) (found bool) {
	for _, m := range w.Movies {
		if reflect.DeepEqual(m.MakeHash(), hash) {
			found = true
		}
	}
	return
}

func (w *MockWriter) WriteEpisode(e *datasets.EpisodeIdentity) error {
	if !w.findMovie(e.MakeParentHash()) {
		return NotFoundErr
	}
	w.Episodes = append(w.Episodes, e)
	return nil
}

func (w *MockWriter) WriteAttribute(hash datasets.Hash, attr datasets.Attribute) error {
	w.Attrs = append(w.Attrs, &HashAndAttribute{
		Hash: hash,
		Attr: attr,
	})
	return nil
}

var (
	seqIdBucketKey       = []byte("seqId")
	moviesBucketKey      = []byte("movies")
	attrsBucketKey       = []byte("attrs")
	searchIndexBucketKey = []byte("search")

	NotFoundErr = errors.New("movie not found")
)

func itob32(i uint32) []byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], i)
	return b[:]
}

func btoi32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

const attrKeySize = 9 //4 + 1 + 4

func makeAttrKey(movieSeqId, attrSeqId uint32, attrType datasets.AttrType) []byte {
	return append(itob32(movieSeqId), append([]byte{byte(attrType)}, itob32(attrSeqId)...)...)
}

func attrKeyGetType(key []byte) datasets.AttrType {
	return datasets.AttrType(key[4])
}

type clock struct {
	t time.Time
}

func (c *clock) reset() time.Duration {
	d := time.Now().Sub(c.t)
	c.t = time.Now()
	return d
}

type BoltdbWriter struct {
	db                *bolt.DB
	tx                *bolt.Tx
	writes            int
	maxWritesPerTx    int
	seqIdBucket       *bolt.Bucket
	moviesBucket      *bolt.Bucket
	attrsBucket       *bolt.Bucket
	searchIndexBucket *bolt.Bucket
}

func NewWriter(dbPath string, maxWritesPerTx int) (*BoltdbWriter, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.NoSync = true
	db.NoGrowSync = true

	for _, bucketKey := range [][]byte{seqIdBucketKey, moviesBucketKey, attrsBucketKey, searchIndexBucketKey} {
		if err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(bucketKey)
			return err
		}); err != nil {
			return nil, err
		}
	}

	return &BoltdbWriter{
		db:             db,
		maxWritesPerTx: maxWritesPerTx,
	}, nil
}

func (w *BoltdbWriter) ensureTxn() error {
	if w.tx != nil && w.writes >= w.maxWritesPerTx {
		if err := w.tx.Commit(); err != nil {
			return err
		}
		w.tx = nil
	}

	if w.tx == nil {
		tx, err := w.db.Begin(true)
		if err != nil {
			return err
		}
		w.writes = 0
		w.tx = tx

		w.seqIdBucket = w.tx.Bucket(seqIdBucketKey)
		w.moviesBucket = w.tx.Bucket(moviesBucketKey)
		w.attrsBucket = w.tx.Bucket(attrsBucketKey)
		w.searchIndexBucket = w.tx.Bucket(searchIndexBucketKey)
	}

	return nil
}

func (w *BoltdbWriter) writeSeqId(hash datasets.Hash) (uint64, error) {
	next, err := w.seqIdBucket.NextSequence()
	if err != nil {
		return 0, err
	}

	if err := w.seqIdBucket.Put(itob32(uint32(next)), hash[:]); err != nil {
		return 0, err
	}

	w.writes += 1

	return next, nil
}

func (w *BoltdbWriter) writeSearchIndexEntry(movie *datasets.MovieIdentity, seqId uint64) error {

	indexValue := movie.SearchIndexValue()

	value := append(w.searchIndexBucket.Get(indexValue), itob32(uint32(seqId))...)

	if err := w.searchIndexBucket.Put(indexValue, value); err != nil {
		return err
	}

	w.writes += 1

	return nil
}

func (w *BoltdbWriter) writeNewEntry(hash datasets.Hash, entry *DBEntry) error {
	value, err := entry.MarshalMsg([]byte{})
	if err != nil {
		return err
	}

	if err := w.moviesBucket.Put(hash[:], value); err != nil {
		return err
	}

	w.writes += 1

	return nil
}

func (w *BoltdbWriter) WriteMovie(movie *datasets.MovieIdentity) error {
	if err := w.ensureTxn(); err != nil {
		return err
	}

	hash := movie.MakeHash()

	seqId, err := w.writeSeqId(hash)
	if err != nil {
		return err
	}

	if err := w.writeNewEntry(hash, &DBEntry{
		SeqId: uint32(seqId),
		Movie: movie,
	}); err != nil {
		return err
	}

	return w.writeSearchIndexEntry(movie, seqId)
}

func (w *BoltdbWriter) WriteEpisode(episode *datasets.EpisodeIdentity) error {
	if err := w.ensureTxn(); err != nil {
		return err
	}

	_, err := w.getSeqId(episode.MakeParentHash())
	if err != nil {
		return err
	}

	hash := episode.MakeHash()

	seqId, err := w.writeSeqId(hash)
	if err != nil {
		return err
	}

	return w.writeNewEntry(hash, &DBEntry{
		SeqId:   uint32(seqId),
		Episode: episode,
	})
}

func (w *BoltdbWriter) getSeqId(hash datasets.Hash) (uint32, error) {
	value := w.moviesBucket.Get(hash[:])
	if value == nil {
		return 0, NotFoundErr
	}

	var entry DBEntry
	if _, err := entry.UnmarshalMsg(value); err != nil {
		return 0, err
	}

	return entry.SeqId, nil
}

func (w *BoltdbWriter) WriteAttribute(hash datasets.Hash, attr datasets.Attribute) error {
	if err := w.ensureTxn(); err != nil {
		return err
	}

	movieSeqId, err := w.getSeqId(hash)
	if err != nil {
		return err
	}

	attrSeqId, err := w.attrsBucket.NextSequence()
	if err != nil {
		return err
	}

	var attrValue []byte

	if attrValue, err = attr.MarshalMsg([]byte{}); err != nil {
		return err
	}

	err = w.attrsBucket.Put(makeAttrKey(movieSeqId, uint32(attrSeqId), attr.GetType()), attrValue)

	w.writes += 1

	return err
}

func (w *BoltdbWriter) Close() error {
	if w.tx != nil {
		if err := w.tx.Commit(); err != nil {
			return err
		}
	}
	return w.db.Close()
}
