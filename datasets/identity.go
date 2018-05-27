package datasets

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const (
	HashSize = md5.Size * 2
)

type Hash [HashSize]byte

func (h Hash) Lo() []byte {
	return h[:HashSize/2]
}

func (h *Hash) FromSlice(s []byte) {
	copy(h[:], s)
}

func (h *Hash) FromHex(src []byte) error {
	_, err := hex.Decode(h[:], src)
	return err
}

func (h Hash) ToHex() []byte {
	return []byte(hex.EncodeToString(h[:]))
}

func (h Hash) ToMapKey() string {
	return string(h[:])
}

func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(h.ToHex()))
}

type Identity interface {
	MakeHash() Hash
}

//go:generate msgp
type MovieIdentity struct {
	Title string
	Year  int
}

func (m *MovieIdentity) SearchIndexValue() []byte {
	return bytes.ToLower([]byte(m.Title))
}

func (m *MovieIdentity) MakeHash() Hash {
	var hash Hash
	copy(hash[0:md5.Size], makeHash(m.Title, m.Year))
	return hash
}

func makeHash(title string, year int) []byte {
	h := md5.Sum([]byte(fmt.Sprintf("%s(%d)", title, year)))
	return h[:]
}

//go:generate msgp
type EpisodeIdentity struct {
	Title         string
	Year          int
	EpisodeTitle  string
	EpisodeNumber int
	SeassonNumber int
}

func (e *EpisodeIdentity) MakeHash() Hash {
	episodeSum := md5.Sum([]byte(fmt.Sprintf("%s(%d) {%s(%d.%d)}", e.Title, e.Year, e.EpisodeTitle, e.SeassonNumber, e.EpisodeNumber)))

	var hash Hash
	copy(hash[:md5.Size], makeHash(e.Title, e.Year))
	copy(hash[md5.Size:], episodeSum[:])
	return hash
}

func (e *EpisodeIdentity) MakeParentHash() Hash {
	var hash Hash
	copy(hash[0:md5.Size], makeHash(e.Title, e.Year))
	return hash
}
