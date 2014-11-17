package goserbench

import (
	"os"
	"reflect"
	"testing"
)

var (
	validate = os.Getenv("VALIDATE")
)

func benchMarshal(
	b *testing.B,
	s Serializer,
	generate func() interface{}, // Generate a struct to marshal
) {
	dataLen := 1000
	data := make([]interface{}, dataLen)
	for i := 0; i < dataLen; i++ {
		data[i] = generate()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Randomize selection of the struct to serialize in order to avoid cache hits and compiler optimization.
		n := crand.Intn(len(data))
		_, err := s.Marshal(data[n])
		if err != nil {
			b.Fatalf("%s failed to marshal: %s (%v)", s, err, data[n])
		}
	}
}

func benchUnmarshal(
	b *testing.B,
	s Serializer, // The serializer that will be used to marshal or unmarshal
	generate func() interface{}, // Generate a struct to marshal and unmarshal
	o interface{}, // The struct to unmarshal into. Must be same type as return value of generate().
) {
	dataLen := 1000
	var err error
	data := make([]interface{}, dataLen)
	for i := 0; i < dataLen; i++ {
		data[i] = generate()
	}
	ser := make([][]byte, dataLen)
	for i, d := range data {
		ser[i], err = s.Marshal(d)
		if err != nil {
			b.Fatalf("%s failed to marshal: %s (%v)", s, err, d)
		}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Randomize selection of the struct to deserialize in order to avoid cache hits and compiler optimization.
		n := crand.Intn(len(ser))
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			// Don't measure validation.
			b.StopTimer()
			if !reflect.DeepEqual(data[n], o) {
				b.Fatalf("unmarshaled object differed:\n\tpre-marshal:\t%v\n\tpost-marshal:\t%v\n", data[n], o)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkJsonMarshalPrimitive(b *testing.B) {
	benchMarshal(b, JsonSerializer{}, genPrimitive)
}

func BenchmarkJsonMarshalData(b *testing.B) {
	benchMarshal(b, JsonSerializer{}, genData)
}

func BenchmarkJsonMarshalMap(b *testing.B) {
	benchMarshal(b, JsonSerializer{}, genMap)
}

func BenchmarkJsonUnmarshalPrimitive(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{}, genPrimitive, &Primitive{})
}

func BenchmarkJsonUnmarshalData(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{}, genData, &Data{})
}

func BenchmarkJsonUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{}, genMap, &Map{})
}

func BenchmarkGobMarshalPrimitive(b *testing.B) {
	benchMarshal(b, GobSerializer{}, genPrimitive)
}

func BenchmarkGobMarshalData(b *testing.B) {
	benchMarshal(b, GobSerializer{}, genData)
}

func BenchmarkGobMarshalMap(b *testing.B) {
	benchMarshal(b, GobSerializer{}, genMap)
}

func BenchmarkGobUnmarshalPrimitive(b *testing.B) {
	benchUnmarshal(b, GobSerializer{}, genPrimitive, &Primitive{})
}

func BenchmarkGobUnmarshalData(b *testing.B) {
	benchUnmarshal(b, GobSerializer{}, genData, &Data{})
}

func BenchmarkGobUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, GobSerializer{}, genMap, &Map{})
}

func BenchmarkMsgpMarshalPrimitive(b *testing.B) {
	benchMarshal(b, MsgpSerializer{}, genPrimitive)
}

func BenchmarkMsgpMarshalData(b *testing.B) {
	benchMarshal(b, MsgpSerializer{}, genData)
}

func BenchmarkMsgpMarshalMap(b *testing.B) {
	benchMarshal(b, MsgpSerializer{}, genMap)
}

func BenchmarkMsgpUnmarshalPrimitive(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{}, genPrimitive, &Primitive{})
}

func BenchmarkMsgpUnmarshalData(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{}, genData, &Data{})
}

func BenchmarkMsgpUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{}, genMap, &Map{})
}

func BenchmarkProtobufMarshalPrimitive(b *testing.B) {
	benchMarshal(b, ProtobufSerializer{}, genPBPrimitive)
}

func BenchmarkProtobufMarshalData(b *testing.B) {
	benchMarshal(b, ProtobufSerializer{}, genPBData)
}

func BenchmarkProtobufMarshalMap(b *testing.B) {
	benchMarshal(b, ProtobufSerializer{}, genPBMap)
}

func BenchmarkProtobufUnmarshalPrimitive(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{}, genPBPrimitive, &PBPrimitive{})
}

func BenchmarkProtobufUnmarshalData(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{}, genPBData, &PBData{})
}

func BenchmarkProtobufUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{}, genPBMap, &PBMap{})
}

/*
func BenchmarkBsonMarshalPrimitive(b *testing.B) {
	benchMarshal(b, BsonSerializer{}, genPrimitive)
}

func BenchmarkBsonUnmarshalPrimitive(b *testing.B) {
	benchUnmarshal(b, BsonSerializer{}, genPrimitive, &Primitive{})
}
*/

/*
func BenchmarkVmihailencoMsgpackMarshal(b *testing.B) {
	benchMarshal(b, VmihailencoMsgpackSerializer(0))
}

func BenchmarkVmihailencoMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, VmihailencoMsgpackSerializer(0))
}

func BenchmarkXdrMarshal(b *testing.B) {
	benchMarshal(b, XdrSerializer(0))
}

func BenchmarkXdrUnmarshal(b *testing.B) {
	benchUnmarshal(b, XdrSerializer(0))
}

func BenchmarkUgorjiCodecMsgpackMarshal(b *testing.B) {
	s := NewUgorjiCodecSerializer("msgpack", &codec.MsgpackHandle{})
	benchMarshal(b, s)
}

func BenchmarkUgorjiCodecMsgpackUnmarshal(b *testing.B) {
	s := NewUgorjiCodecSerializer("msgpack", &codec.MsgpackHandle{})
	benchUnmarshal(b, s)
}

func BenchmarkUgorjiCodecBincMarshal(b *testing.B) {
	s := NewUgorjiCodecSerializer("binc", &codec.BincHandle{})
	benchMarshal(b, s)
}

func BenchmarkUgorjiCodecBincUnmarshal(b *testing.B) {
	s := NewUgorjiCodecSerializer("binc", &codec.BincHandle{})
	benchUnmarshal(b, s)
}

func BenchmarkSerealMarshal(b *testing.B) {
	benchMarshal(b, SerealSerializer(0))
}

func BenchmarkSerealUnmarshal(b *testing.B) {
	benchUnmarshal(b, SerealSerializer(0))
}

func BenchmarkBinaryMarshal(b *testing.B) {
	benchMarshal(b, BinarySerializer(0))
}

func BenchmarkBinaryUnmarshal(b *testing.B) {
	benchUnmarshal(b, BinarySerializer(0))
}

func BenchmarkMsgpMarshal(b *testing.B) {
	benchMarshal(b, MsgpSerializer{})
}

func BenchmarkMsgpUnmarshal(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{})
}
*/
