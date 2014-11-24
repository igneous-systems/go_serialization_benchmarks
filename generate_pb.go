package goserbench

// Protobuf analogs of functions in generate.go

import (
	"math/rand"

	"code.google.com/p/goprotobuf/proto"
)

func genPBPrimitive() interface{} {
	return &PBPrimitive{
		Str:     proto.String(randString(16)),
		Int64:   proto.Int64(rand.Int63()),
		Bool:    proto.Bool(randBool()),
		Float64: proto.Float64(rand.Float64()),
	}
}

func genPBData() interface{} {
	return &PBData{
		Metadata: proto.String(randString(8)),
		Data:     randBytes(16384),
	}
}

func genPBMap() interface{} {
	mapSize := 100
	keySize := 100
	valueSize := 100
	m := &PBMap{Kvps: []*PBMap_KVP{}}
	for i := 0; i < mapSize; i++ {
		key := randString(keySize)
		value := randString(valueSize)
		newKvp := &PBMap_KVP{
			Key:   proto.String(key),
			Value: proto.String(value),
		}
		m.Kvps = append(m.Kvps, newKvp)
	}
	return m
}
