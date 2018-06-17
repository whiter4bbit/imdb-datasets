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
	db             *bolt.DB
	tx             *bolt.Tx
	writes         int
	maxWritesPerTx int
	buckets        map[string]*bolt.Bucket
}

func NewWriter(dbPath string, maxWritesPerTx int) (*BoltdbWriter, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.NoSync = true
	db.NoGrowSync = true

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
		w.buckets = make(map[string]*bolt.Bucket)
	}

	return nil
}

func (w *BoltdbWriter) bucket(name []byte) (*bolt.Bucket, error) {
	bucket := w.buckets[string(name)]
	if bucket == nil {
		b, err := w.tx.CreateBucketIfNotExists(name)
		if err != nil {
			return nil, err
		}

		w.buckets[string(name)] = b

		bucket = b
	}

	return bucket, nil
}

func (w *BoltdbWriter) writeSeqId(hash datasets.Hash) (uint64, error) {
	b, err := w.bucket(seqIdBucketKey)
	if err != nil {
		return 0, err
	}

	next, err := b.NextSequence()
	if err != nil {
		return 0, err
	}

	if err := b.Put(itob32(uint32(next)), hash[:]); err != nil {
		return 0, err
	}

	w.writes += 1

	return next, nil
}

func (w *BoltdbWriter) writeSearchIndexEntry(movie *datasets.MovieIdentity, seqId uint64) error {
	b, err := w.bucket(searchIndexBucketKey)
	if err != nil {
		return err
	}

	if err := b.Put(movie.SearchIndexValue(), itob32(uint32(seqId))); err != nil {
		return err
	}

	w.writes += 1

	return nil
}

func (w *BoltdbWriter) writeNewEntry(hash datasets.Hash, entry *DBEntry) error {
	b, err := w.bucket(moviesBucketKey)
	if err != nil {
		return err
	}

	value, err := entry.MarshalMsg([]byte{})
	if err != nil {
		return err
	}

	if err := b.Put(hash[:], value); err != nil {
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

func (w *BoltdbWriter) updateEntry(hash []byte, f func(*DBEntry)) error {
	if err := w.ensureTxn(); err != nil {
		return err
	}

	b, err := w.bucket(moviesBucketKey)
	if err != nil {
		return err
	}

	value := b.Get(hash)
	if value == nil {
		return NotFoundErr
	}

	entry := &DBEntry{}

	if _, err := entry.UnmarshalMsg(value); err != nil {
		return err
	}

	f(entry)

	if value, err = entry.MarshalMsg([]byte{}); err != nil {
		return err
	}

	if err := b.Put(hash, value); err != nil {
		return err
	}

	w.writes += 1

	return nil
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
	movies, err := w.bucket(moviesBucketKey)
	if err != nil {
		return 0, err
	}

	value := movies.Get(hash[:])
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

	attributes, err := w.bucket(attrsBucketKey)
	if err != nil {
		return err
	}

	attrSeqId, err := attributes.NextSequence()
	if err != nil {
		return err
	}

	var attrValue []byte

	if attrValue, err = attr.MarshalMsg([]byte{}); err != nil {
		return err
	}

	err = attributes.Put(makeAttrKey(movieSeqId, uint32(attrSeqId), attr.GetType()), attrValue)

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
