package goserbench

// Wrappers around serializers to standardize their interface.

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"code.google.com/p/goprotobuf/proto"

	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/binary"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/philhofer/msgp/msgp"
	"github.com/ugorji/go/codec"
	vmihailenco "github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2/bson"
)

type Serializer interface {
	Marshal(o interface{}) ([]byte, error)
	Unmarshal(d []byte, o interface{}) error
	String() string
}

type MsgpSerializer struct{}

func (m MsgpSerializer) Marshal(o interface{}) ([]byte, error) {
	// The byte array input to MarshalMsg must have enough capacity to fit the msg.
	b := make([]byte, 0, o.(msgp.Sizer).Msgsize())
	return o.(msgp.Marshaler).MarshalMsg(b)
}

func (m MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(msgp.Unmarshaler).UnmarshalMsg(d)
	return err
}

func (m MsgpSerializer) String() string {
	return "msgp"
}

type JsonSerializer struct{}

func (m JsonSerializer) Marshal(o interface{}) ([]byte, error) {
	d, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (m JsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func (j JsonSerializer) String() string {
	return "json"
}

type GobSerializer struct{}

func (g GobSerializer) Marshal(o interface{}) ([]byte, error) {
	b := &bytes.Buffer{}
	e := gob.NewEncoder(b)
	err := e.Encode(o)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (g GobSerializer) Unmarshal(d []byte, o interface{}) error {
	b := bytes.NewBuffer(d)
	e := gob.NewDecoder(b)
	err := e.Decode(o)
	return err
}

func (g GobSerializer) String() string {
	return "gob"
}

type ProtobufSerializer struct{}

func (p ProtobufSerializer) Marshal(o interface{}) ([]byte, error) {
	return proto.Marshal(o.(proto.Message))
}

func (p ProtobufSerializer) Unmarshal(d []byte, o interface{}) error {
	return proto.Unmarshal(d, o.(proto.Message))
}

func (p ProtobufSerializer) String() string {
	return "protobuf"
}

type VmihailencoMsgpackSerializer struct{}

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) ([]byte, error) {
	d, err := vmihailenco.Marshal(o)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func (m VmihailencoMsgpackSerializer) String() string {
	return "vmihailenco-msgpack"
}

type BsonSerializer struct{}

func (m BsonSerializer) Marshal(o interface{}) ([]byte, error) {
	d, err := bson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (m BsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return bson.Unmarshal(d, o)
}

func (j BsonSerializer) String() string {
	return "bson"
}

type XdrSerializer struct{}

func (x XdrSerializer) Marshal(o interface{}) ([]byte, error) {
	d, err := xdr.Marshal(o)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (x XdrSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func (x XdrSerializer) String() string {
	return "xdr"
}

type UgorjiCodecSerializer struct {
	name string
	h    codec.Handle
}

func NewUgorjiCodecSerializer(name string, h codec.Handle) *UgorjiCodecSerializer {
	return &UgorjiCodecSerializer{
		name: name,
		h:    h,
	}
}

func (u *UgorjiCodecSerializer) Marshal(o interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := codec.NewEncoder(buf, u.h)
	err := enc.Encode(o)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (u *UgorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	buf := bytes.NewReader(d)
	dec := codec.NewDecoder(buf, u.h)
	return dec.Decode(o)
}

func (u *UgorjiCodecSerializer) String() string {
	return "ugorjicodec-" + u.name
}

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o interface{}) ([]byte, error) {
	d, err := sereal.Marshal(o)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func (s SerealSerializer) String() string {
	return "sereal"
}

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o interface{}) ([]byte, error) {
	d, err := binary.Marshal(o)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func (b BinarySerializer) String() string {
	return "binary"
}
