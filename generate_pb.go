package goserbench

// Protobuf analogs of functions in generate.go

import "code.google.com/p/goprotobuf/proto"

func genPBPrimitive() interface{} {
	return &PBPrimitive{
		Str:     proto.String(string(crand.Bytes(16))),
		Int64:   proto.Int64(crand.Int63()),
		Bool:    proto.Bool(crand.Bool()),
		Float64: proto.Float64(crand.Float64()),
	}
}

func genPBData() interface{} {
	return &PBData{
		Metadata: proto.String(string(crand.Bytes(8))),
		Data:     crand.Bytes(16384),
	}
}

func genPBMap() interface{} {
	mapSize := 100
	keySize := 100
	valueSize := 100
	m := &PBMap{MapSS: &PBMapSS{Kvps: []*PBKVPSS{}}}
	for i := 0; i < mapSize; i++ {
		key := string(crand.Bytes(keySize))
		value := string(crand.Bytes(valueSize))
		newKvp := &PBKVPSS{
			Key:   proto.String(key),
			Value: proto.String(value),
		}
		m.MapSS.Kvps = append(m.MapSS.Kvps, newKvp)
	}
	return m
}
