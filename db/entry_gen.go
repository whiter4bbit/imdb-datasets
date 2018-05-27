package db

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
	"github.com/whiter4bbit/imdb-datasets/datasets"
)

// DecodeMsg implements msgp.Decodable
func (z *DBEntry) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Movie":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(datasets.MovieIdentity)
				}
				err = z.Movie.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Episode":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Episode = nil
			} else {
				if z.Episode == nil {
					z.Episode = new(datasets.EpisodeIdentity)
				}
				err = z.Episode.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "SeqId":
			z.SeqId, err = dc.ReadUint32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DBEntry) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Movie"
	err = en.Append(0x83, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if err != nil {
		return
	}
	if z.Movie == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Movie.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Episode"
	err = en.Append(0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	if z.Episode == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Episode.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "SeqId"
	err = en.Append(0xa5, 0x53, 0x65, 0x71, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.SeqId)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DBEntry) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Movie"
	o = append(o, 0x83, 0xa5, 0x4d, 0x6f, 0x76, 0x69, 0x65)
	if z.Movie == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Movie.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Episode"
	o = append(o, 0xa7, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65)
	if z.Episode == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Episode.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "SeqId"
	o = append(o, 0xa5, 0x53, 0x65, 0x71, 0x49, 0x64)
	o = msgp.AppendUint32(o, z.SeqId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DBEntry) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Movie":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Movie = nil
			} else {
				if z.Movie == nil {
					z.Movie = new(datasets.MovieIdentity)
				}
				bts, err = z.Movie.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Episode":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Episode = nil
			} else {
				if z.Episode == nil {
					z.Episode = new(datasets.EpisodeIdentity)
				}
				bts, err = z.Episode.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "SeqId":
			z.SeqId, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DBEntry) Msgsize() (s int) {
	s = 1 + 6
	if z.Movie == nil {
		s += msgp.NilSize
	} else {
		s += z.Movie.Msgsize()
	}
	s += 8
	if z.Episode == nil {
		s += msgp.NilSize
	} else {
		s += z.Episode.Msgsize()
	}
	s += 6 + msgp.Uint32Size
	return
}
