package goserbench

import (
	"math/rand"
	"reflect"
	"testing"
)

var validate = false

func benchMarshal(
	b *testing.B,
	s Serializer,
	generate func() interface{}, // Generate a pointer to a struct to marshal
) {
	// Generate data.
	dataLen := 1000
	data := make([]interface{}, dataLen)
	for i := 0; i < dataLen; i++ {
		data[i] = generate()
	}

	// Benchmark marshaling.
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Randomize selection of the struct to serialize in order to avoid cache hits and compiler optimization.
		n := rand.Intn(len(data))
		_, err := s.Marshal(data[n])
		if err != nil {
			b.Fatalf("%s failed to marshal: %s (%v)", s, err, data[n])
		}
	}
}

func benchUnmarshal(
	b *testing.B,
	s Serializer, // The serializer that will be used to marshal or unmarshal
	generate func() interface{}, // Generate a pointer to a struct to marshal and unmarshal
) {
	// Generate data.
	dataLen := 1000
	var err error
	data := make([]interface{}, dataLen)
	for i := 0; i < dataLen; i++ {
		data[i] = generate()
	}

	// Marshal the data.
	ser := make([][]byte, dataLen)
	for i, d := range data {
		ser[i], err = s.Marshal(d)
		if err != nil {
			b.Fatalf("%s failed to marshal: %s (%v)", s, err, d)
		}
	}

	// Generate random indices to unmarshal.
	randIndex := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		randIndex[i] = rand.Intn(dataLen)
	}

	// Generate a slice of pointers to structs to unmarshal into.
	dataPtrType := reflect.TypeOf(data[0])
	if dataPtrType.Kind() != reflect.Ptr {
		b.Fatalf("Generate must return a pointer type, got %v")
	}
	dataType := dataPtrType.Elem()
	deser := make([]interface{}, b.N)
	for i := 0; i < b.N; i++ {
		deser[i] = reflect.New(dataType).Interface()
	}

	// Benchmark unmarshaling.
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Randomize selection of the struct to deserialize in order to avoid cache hits and compiler optimization.
		err := s.Unmarshal(ser[randIndex[i]], deser[i])
		if err != nil {
			b.Fatalf("%s failed to unmarshal %s: %s", s, ser[randIndex[i]], err)
		}
	}
	b.StopTimer()

	// Validate unmarshalled data.
	if validate {
		for i := 0; i < b.N; i++ {
			if !reflect.DeepEqual(data[randIndex[i]], deser[i]) {
				b.Fatalf("unmarshaled object differed:\n\tpre-marshal:\t%#v\n\tpost-marshal:\t%#v\n", data[randIndex[i]], deser[i])
			}
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
	benchUnmarshal(b, JsonSerializer{}, genPrimitive)
}

func BenchmarkJsonUnmarshalData(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{}, genData)
}

func BenchmarkJsonUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{}, genMap)
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
	benchUnmarshal(b, GobSerializer{}, genPrimitive)
}

func BenchmarkGobUnmarshalData(b *testing.B) {
	benchUnmarshal(b, GobSerializer{}, genData)
}

func BenchmarkGobUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, GobSerializer{}, genMap)
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
	benchUnmarshal(b, MsgpSerializer{}, genPrimitive)
}

func BenchmarkMsgpUnmarshalData(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{}, genData)
}

func BenchmarkMsgpUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, MsgpSerializer{}, genMap)
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
	benchUnmarshal(b, ProtobufSerializer{}, genPBPrimitive)
}

func BenchmarkProtobufUnmarshalData(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{}, genPBData)
}

func BenchmarkProtobufUnmarshalMap(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{}, genPBMap)
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
