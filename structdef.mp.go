package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"github.com/philhofer/msgp/msgp"
)


// MarshalMsg implements the msgp.Marshaler interface
func (z *Primitive) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 4)

	o = msgp.AppendString(o, "String")

	o = msgp.AppendString(o, z.String)

	o = msgp.AppendString(o, "Int64")

	o = msgp.AppendInt64(o, z.Int64)

	o = msgp.AppendString(o, "Bool")

	o = msgp.AppendBool(o, z.Bool)

	o = msgp.AppendString(o, "Float64")

	o = msgp.AppendFloat64(o, z.Float64)

	return
}
// UnmarshalMsg unmarshals a Primitive from MessagePack, returning any extra bytes
// and any errors encountered
func (z *Primitive) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "String":

			z.String, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "Int64":

			z.Int64, bts, err = msgp.ReadInt64Bytes(bts)

			if err != nil {
				return
			}

		case "Bool":

			z.Bool, bts, err = msgp.ReadBoolBytes(bts)

			if err != nil {
				return
			}

		case "Float64":

			z.Float64, bts, err = msgp.ReadFloat64Bytes(bts)

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

// Msgsize implements the msgp.Sizer interface
func (z *Primitive) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 6

	s += msgp.StringPrefixSize + len(z.String)
	s += msgp.StringPrefixSize + 5

	s += msgp.Int64Size
	s += msgp.StringPrefixSize + 4

	s += msgp.BoolSize
	s += msgp.StringPrefixSize + 7

	s += msgp.Float64Size

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *Primitive) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "String":

			z.String, err = dc.ReadString()

			if err != nil {
				return
			}

		case "Int64":

			z.Int64, err = dc.ReadInt64()

			if err != nil {
				return
			}

		case "Bool":

			z.Bool, err = dc.ReadBool()

			if err != nil {
				return
			}

		case "Float64":

			z.Float64, err = dc.ReadFloat64()

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

// EncodeMsg implements the msgp.Encodable interface
func (z *Primitive) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(4)
	if err != nil {
		return
	}

	err = en.WriteString("String")
	if err != nil {
		return
	}

	err = en.WriteString(z.String)

	if err != nil {
		return
	}

	err = en.WriteString("Int64")
	if err != nil {
		return
	}

	err = en.WriteInt64(z.Int64)

	if err != nil {
		return
	}

	err = en.WriteString("Bool")
	if err != nil {
		return
	}

	err = en.WriteBool(z.Bool)

	if err != nil {
		return
	}

	err = en.WriteString("Float64")
	if err != nil {
		return
	}

	err = en.WriteFloat64(z.Float64)

	if err != nil {
		return
	}

	return
}

// MarshalMsg implements the msgp.Marshaler interface
func (z *Data) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 2)

	o = msgp.AppendString(o, "Metadata")

	o = msgp.AppendString(o, z.Metadata)

	o = msgp.AppendString(o, "Data")

	o = msgp.AppendBytes(o, z.Data)

	return
}
// UnmarshalMsg unmarshals a Data from MessagePack, returning any extra bytes
// and any errors encountered
func (z *Data) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Metadata":

			z.Metadata, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "Data":

			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)

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

// Msgsize implements the msgp.Sizer interface
func (z *Data) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 8

	s += msgp.StringPrefixSize + len(z.Metadata)
	s += msgp.StringPrefixSize + 4

	s += msgp.BytesPrefixSize + len(z.Data)

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *Data) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Metadata":

			z.Metadata, err = dc.ReadString()

			if err != nil {
				return
			}

		case "Data":

			z.Data, err = dc.ReadBytes(z.Data)

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

// EncodeMsg implements the msgp.Encodable interface
func (z *Data) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}

	err = en.WriteString("Metadata")
	if err != nil {
		return
	}

	err = en.WriteString(z.Metadata)

	if err != nil {
		return
	}

	err = en.WriteString("Data")
	if err != nil {
		return
	}

	err = en.WriteBytes(z.Data)

	if err != nil {
		return
	}

	return
}

// MarshalMsg implements the msgp.Marshaler interface
func (z *Map) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	o = msgp.AppendMapHeader(o, 1)

	o = msgp.AppendString(o, "Map")

	o = msgp.AppendMapHeader(o, uint32(len(z.Map)))
	for xvk, bzg := range z.Map {
		o = msgp.AppendString(o, xvk)

		o = msgp.AppendString(o, bzg)

	}

	return
}
// UnmarshalMsg unmarshals a Map from MessagePack, returning any extra bytes
// and any errors encountered
func (z *Map) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Map":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Map == nil && msz > 0 {
				z.Map = make(map[string]string, int(msz))
			} else if len(z.Map) > 0 {
				for key, _ := range z.Map {
					delete(z.Map, key)
				}
			}
			for inx := uint32(0); inx < msz; inx++ {
				var xvk string
				var bzg string
				xvk, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}

				bzg, bts, err = msgp.ReadStringBytes(bts)

				if err != nil {
					return
				}

				z.Map[xvk] = bzg
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

// Msgsize implements the msgp.Sizer interface
func (z *Map) Msgsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 3

	s += msgp.MapHeaderSize
	if z.Map != nil {
		for xvk, bzg := range z.Map {
			_ = bzg
			s += msgp.StringPrefixSize + len(xvk)

			s += msgp.StringPrefixSize + len(bzg)

		}
	}

	return
}

// DecodeMsg implements the msgp.Decodable interface
func (z *Map) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Map":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Map == nil && msz > 0 {
				z.Map = make(map[string]string, int(msz))
			} else if len(z.Map) > 0 {
				for key, _ := range z.Map {
					delete(z.Map, key)
				}
			}
			for inx := uint32(0); inx < msz; inx++ {
				var xvk string
				var bzg string
				xvk, err = dc.ReadString()
				if err != nil {
					return
				}

				bzg, err = dc.ReadString()

				if err != nil {
					return
				}

				z.Map[xvk] = bzg
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

// EncodeMsg implements the msgp.Encodable interface
func (z *Map) EncodeMsg(en *msgp.Writer) (err error) {

	err = en.WriteMapHeader(1)
	if err != nil {
		return
	}

	err = en.WriteString("Map")
	if err != nil {
		return
	}

	err = en.WriteMapHeader(uint32(len(z.Map)))
	if err != nil {
		return
	}

	for xvk, bzg := range z.Map {
		err = en.WriteString(xvk)
		if err != nil {
			return
		}

		err = en.WriteString(bzg)

		if err != nil {
			return
		}

	}

	return
}
