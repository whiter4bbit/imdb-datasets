package db

import (
	"bytes"
	bolt "github.com/coreos/bbolt"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"sort"
)

const maxMoviesToReturn = 100

type Reader struct {
	db *bolt.DB
}

type Filter struct {
	Hashes []datasets.Hash
}

func NewReader(dbPath string) (*Reader, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &Reader{
		db: db,
	}, nil
}

func (r *Reader) bulkGet(tx *bolt.Tx, bucket []byte, keys [][]byte) [][]byte {
	sort.Slice(keys, func(i, j int) bool {
		return bytes.Compare(keys[i], keys[j]) < 0
	})

	b := tx.Bucket(bucket)

	c := b.Cursor()

	var values [][]byte

	for _, key := range keys {
		foundKey, value := c.Seek(key)

		if bytes.Equal(foundKey, key) {
			values = append(values, value)
		}
	}

	return values
}

type entry struct {
	key   []byte
	value []byte
}

func (r *Reader) bulkGetByPrefix(tx *bolt.Tx, bucket, prefix []byte, size int, skipFirst bool) []*entry {
	b := tx.Bucket(bucket)

	c := b.Cursor()

	startKey := append(prefix, make([]byte, size-len(prefix))...)

	var entries []*entry

	key, value := c.Seek(startKey)

	if skipFirst {
		key, value = c.Next()
	}

	for ; key != nil && bytes.Equal(key[:len(prefix)], prefix); key, value = c.Next() {
		entries = append(entries, &entry{
			key:   key,
			value: value,
		})
	}

	return entries
}

func (r *Reader) getAttrs(tx *bolt.Tx, seqId uint32) (*datasets.Attributes, error) {
	attrs := &datasets.Attributes{}
	for _, entry := range r.bulkGetByPrefix(tx, attrsBucketKey, itob32(seqId), attrKeySize, false) {
		attrType := attrKeyGetType(entry.key)

		attr, err := datasets.UnmarshalAttribute(attrType, entry.value)
		if err != nil {
			return nil, err
		}

		attr.Update(attrs)
	}
	return attrs, nil
}

func (r *Reader) ForEachTitleIndex(f func(title []byte, seqId uint32) error) error {
	tx, err := r.db.Begin(false)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket(searchIndexBucketKey)

	return b.ForEach(func(key, value []byte) error {
		return f(key, btoi32(value))
	})
}

func (r *Reader) getByIds(tx *bolt.Tx, ids []uint32, filter *Filter) ([]*datasets.Movie, error) {
	var keys [][]byte
	for _, id := range ids {
		keys = append(keys, itob32(id))
	}

	var hashesToSearch map[string]bool
	if filter != nil {
		hashesToSearch = make(map[string]bool)

		for _, hash := range filter.Hashes {
			hashesToSearch[hash.ToMapKey()] = true
		}
	}

	var hashes []datasets.Hash
	for _, value := range r.bulkGet(tx, seqIdBucketKey, keys) {
		if filter != nil && !hashesToSearch[string(value)] {
			continue
		}

		var hash datasets.Hash
		copy(hash[:], value)

		hashes = append(hashes, hash)
	}

	return r.GetByHash(hashes)
}

func (r *Reader) GetByTitleIndex(titles [][]byte) ([]*datasets.Movie, error) {
	tx, err := r.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var ids []uint32
	for _, value := range r.bulkGet(tx, searchIndexBucketKey, titles) {
		for len(value) > 0 {
			ids = append(ids, btoi32(value))
			value = value[4:]
		}
	}

	return r.getByIds(tx, ids, nil)
}

func (r *Reader) GetByTitlePrefix(prefix []byte, filter *Filter) ([]*datasets.Movie, error) {
	tx, err := r.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	startKey := bytes.ToLower(prefix)

	endKey := append(startKey, 0xFF)

	cursor := tx.Bucket(searchIndexBucketKey).Cursor()

	var ids []uint32
	for key, value := cursor.Seek(startKey); key != nil && bytes.Compare(key, endKey) < 0; key, value = cursor.Next() {
		for len(value) > 0 {
			ids = append(ids, btoi32(value))
			value = value[4:]
		}
	}

	if filter == nil && len(ids) > maxMoviesToReturn {
		ids = ids[:maxMoviesToReturn]
	}

	return r.getByIds(tx, ids, filter)
}

func (r *Reader) GetByHash(hashes []datasets.Hash) ([]*datasets.Movie, error) {
	tx, err := r.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var (
		movies []*datasets.Movie
		keys   [][]byte
	)

	for i, _ := range hashes {
		keys = append(keys, hashes[i][:])
	}

	for _, movieValue := range r.bulkGet(tx, moviesBucketKey, keys) {
		var (
			movieEntry DBEntry
			movie      datasets.Movie
		)

		if _, err := movieEntry.UnmarshalMsg(movieValue); err != nil {
			return nil, err
		}

		attrs, err := r.getAttrs(tx, movieEntry.SeqId)
		if err != nil {
			return nil, err
		}

		movie.Attributes = attrs
		movie.MovieIdentity = movieEntry.Movie
		movie.Id = movie.MovieIdentity.MakeHash()

		for _, episodeEntry := range r.bulkGetByPrefix(tx, moviesBucketKey, movie.Id.Lo(), datasets.HashSize, true) {
			var (
				episodeDBEntry DBEntry
				episode        datasets.Episode
			)

			if _, err := episodeDBEntry.UnmarshalMsg(episodeEntry.value); err != nil {
				return nil, err
			}

			episode.EpisodeIdentity = episodeDBEntry.Episode
			episode.Id = episode.EpisodeIdentity.MakeHash()

			attrs, err := r.getAttrs(tx, episodeDBEntry.SeqId)
			if err != nil {
				return nil, err
			}

			episode.Attributes = attrs

			movie.Episodes = append(movie.Episodes, &episode)
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (r *Reader) Close() error {
	return r.db.Close()
}
