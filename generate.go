package goserbench

func genPrimitive() interface{} {
	return &Primitive{
		String:  string(crand.Bytes(16)),
		Int64:   crand.Int63(),
		Bool:    crand.Bool(),
		Float64: crand.Float64(),
	}
}

func genData() interface{} {
	return &Data{
		Metadata: string(crand.Bytes(8)),
		Data:     crand.Bytes(16384),
	}
}

func genMap() interface{} {
	mapSize := 100
	keySize := 100
	valueSize := 100
	m := &Map{Map: map[string]string{}}
	for i := 0; i < mapSize; i++ {
		key := string(crand.Bytes(keySize))
		value := string(crand.Bytes(valueSize))
		m.Map[key] = value
	}
	return m
}
