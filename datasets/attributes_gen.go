package datasets

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Actors) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Actor":
			z.Actor, err = dc.ReadString()
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
func (z Actors) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Actor"
	err = en.Append(0x81, 0xa5, 0x41, 0x63, 0x74, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Actor)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Actors) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Actor"
	o = append(o, 0x81, 0xa5, 0x41, 0x63, 0x74, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Actor)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Actors) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Actor":
			z.Actor, bts, err = msgp.ReadStringBytes(bts)
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
func (z Actors) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Actor)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *AttrType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 byte
		zb0001, err = dc.ReadByte()
		if err != nil {
			return
		}
		(*z) = AttrType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z AttrType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteByte(byte(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z AttrType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendByte(o, byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AttrType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 byte
		zb0001, bts, err = msgp.ReadByteBytes(bts)
		if err != nil {
			return
		}
		(*z) = AttrType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z AttrType) Msgsize() (s int) {
	s = msgp.ByteSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Attributes) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Plot":
			z.Plot, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "Votes":
			z.Votes, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Actors":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Actors) >= int(zb0002) {
				z.Actors = (z.Actors)[:zb0002]
			} else {
				z.Actors = make([]string, zb0002)
			}
			for za0001 := range z.Actors {
				z.Actors[za0001], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "Genres":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0003) {
				z.Genres = (z.Genres)[:zb0003]
			} else {
				z.Genres = make([]string, zb0003)
			}
			for za0002 := range z.Genres {
				z.Genres[za0002], err = dc.ReadString()
				if err != nil {
					return
				}
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
func (z *Attributes) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Plot"
	err = en.Append(0x85, 0xa4, 0x50, 0x6c, 0x6f, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Plot)
	if err != nil {
		return
	}
	// write "Rating"
	err = en.Append(0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Rating)
	if err != nil {
		return
	}
	// write "Votes"
	err = en.Append(0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Votes)
	if err != nil {
		return
	}
	// write "Actors"
	err = en.Append(0xa6, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Actors)))
	if err != nil {
		return
	}
	for za0001 := range z.Actors {
		err = en.WriteString(z.Actors[za0001])
		if err != nil {
			return
		}
	}
	// write "Genres"
	err = en.Append(0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Genres)))
	if err != nil {
		return
	}
	for za0002 := range z.Genres {
		err = en.WriteString(z.Genres[za0002])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Attributes) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Plot"
	o = append(o, 0x85, 0xa4, 0x50, 0x6c, 0x6f, 0x74)
	o = msgp.AppendString(o, z.Plot)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendFloat64(o, z.Rating)
	// string "Votes"
	o = append(o, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Votes)
	// string "Actors"
	o = append(o, 0xa6, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Actors)))
	for za0001 := range z.Actors {
		o = msgp.AppendString(o, z.Actors[za0001])
	}
	// string "Genres"
	o = append(o, 0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Genres)))
	for za0002 := range z.Genres {
		o = msgp.AppendString(o, z.Genres[za0002])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Attributes) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Plot":
			z.Plot, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "Votes":
			z.Votes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Actors":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Actors) >= int(zb0002) {
				z.Actors = (z.Actors)[:zb0002]
			} else {
				z.Actors = make([]string, zb0002)
			}
			for za0001 := range z.Actors {
				z.Actors[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Genres":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0003) {
				z.Genres = (z.Genres)[:zb0003]
			} else {
				z.Genres = make([]string, zb0003)
			}
			for za0002 := range z.Genres {
				z.Genres[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
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
func (z *Attributes) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Plot) + 7 + msgp.Float64Size + 6 + msgp.IntSize + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Actors {
		s += msgp.StringPrefixSize + len(z.Actors[za0001])
	}
	s += 7 + msgp.ArrayHeaderSize
	for za0002 := range z.Genres {
		s += msgp.StringPrefixSize + len(z.Genres[za0002])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Genres) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Genre":
			z.Genre, err = dc.ReadString()
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
func (z Genres) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Genre"
	err = en.Append(0x81, 0xa5, 0x47, 0x65, 0x6e, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Genre)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Genres) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Genre"
	o = append(o, 0x81, 0xa5, 0x47, 0x65, 0x6e, 0x72, 0x65)
	o = msgp.AppendString(o, z.Genre)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Genres) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Genre":
			z.Genre, bts, err = msgp.ReadStringBytes(bts)
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
func (z Genres) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Genre)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Plots) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Plot":
			z.Plot, err = dc.ReadString()
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
func (z Plots) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Plot"
	err = en.Append(0x81, 0xa4, 0x50, 0x6c, 0x6f, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Plot)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Plots) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Plot"
	o = append(o, 0x81, 0xa4, 0x50, 0x6c, 0x6f, 0x74)
	o = msgp.AppendString(o, z.Plot)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Plots) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Plot":
			z.Plot, bts, err = msgp.ReadStringBytes(bts)
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
func (z Plots) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Plot)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Ratings) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Votes":
			z.Votes, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, err = dc.ReadFloat64()
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
func (z Ratings) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Votes"
	err = en.Append(0x82, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Votes)
	if err != nil {
		return
	}
	// write "Rating"
	err = en.Append(0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Rating)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Ratings) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Votes"
	o = append(o, 0x82, 0xa5, 0x56, 0x6f, 0x74, 0x65, 0x73)
	o = msgp.AppendInt(o, z.Votes)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendFloat64(o, z.Rating)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Ratings) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Votes":
			z.Votes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadFloat64Bytes(bts)
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
func (z Ratings) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 7 + msgp.Float64Size
	return
}
