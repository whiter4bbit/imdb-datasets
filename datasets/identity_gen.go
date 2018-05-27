package datasets

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *EpisodeIdentity) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Title":
			z.Title, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Year":
			z.Year, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "EpisodeTitle":
			z.EpisodeTitle, err = dc.ReadString()
			if err != nil {
				return
			}
		case "EpisodeNumber":
			z.EpisodeNumber, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "SeassonNumber":
			z.SeassonNumber, err = dc.ReadInt()
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
func (z *EpisodeIdentity) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Title"
	err = en.Append(0x85, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Title)
	if err != nil {
		return
	}
	// write "Year"
	err = en.Append(0xa4, 0x59, 0x65, 0x61, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Year)
	if err != nil {
		return
	}
	// write "EpisodeTitle"
	err = en.Append(0xac, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.EpisodeTitle)
	if err != nil {
		return
	}
	// write "EpisodeNumber"
	err = en.Append(0xad, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt(z.EpisodeNumber)
	if err != nil {
		return
	}
	// write "SeassonNumber"
	err = en.Append(0xad, 0x53, 0x65, 0x61, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt(z.SeassonNumber)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *EpisodeIdentity) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Title"
	o = append(o, 0x85, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "Year"
	o = append(o, 0xa4, 0x59, 0x65, 0x61, 0x72)
	o = msgp.AppendInt(o, z.Year)
	// string "EpisodeTitle"
	o = append(o, 0xac, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.EpisodeTitle)
	// string "EpisodeNumber"
	o = append(o, 0xad, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.EpisodeNumber)
	// string "SeassonNumber"
	o = append(o, 0xad, 0x53, 0x65, 0x61, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.SeassonNumber)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EpisodeIdentity) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Year":
			z.Year, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeTitle":
			z.EpisodeTitle, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeNumber":
			z.EpisodeNumber, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "SeassonNumber":
			z.SeassonNumber, bts, err = msgp.ReadIntBytes(bts)
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
func (z *EpisodeIdentity) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Title) + 5 + msgp.IntSize + 13 + msgp.StringPrefixSize + len(z.EpisodeTitle) + 14 + msgp.IntSize + 14 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Hash) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes((z)[:])
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Hash) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes((z)[:])
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Hash) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (z)[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Hash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (z)[:])
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Hash) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (HashSize * (msgp.ByteSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MovieIdentity) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Title":
			z.Title, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Year":
			z.Year, err = dc.ReadInt()
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
func (z MovieIdentity) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Title"
	err = en.Append(0x82, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Title)
	if err != nil {
		return
	}
	// write "Year"
	err = en.Append(0xa4, 0x59, 0x65, 0x61, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Year)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MovieIdentity) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Title"
	o = append(o, 0x82, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "Year"
	o = append(o, 0xa4, 0x59, 0x65, 0x61, 0x72)
	o = msgp.AppendInt(o, z.Year)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MovieIdentity) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Year":
			z.Year, bts, err = msgp.ReadIntBytes(bts)
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
func (z MovieIdentity) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Title) + 5 + msgp.IntSize
	return
}
