package goserbench

import "math/rand"

func genPrimitive() interface{} {
	return &Primitive{
		String:  string(randBytes(16)),
		Int64:   rand.Int63(),
		Bool:    randBool(),
		Float64: rand.Float64(),
	}
}

func genData() interface{} {
	return &Data{
		Metadata: string(randBytes(8)),
		Data:     randBytes(16384),
	}
}

func genMap() interface{} {
	mapSize := 100
	keySize := 100
	valueSize := 100
	m := &Map{Map: map[string]string{}}
	for i := 0; i < mapSize; i++ {
		key := string(randBytes(keySize))
		value := string(randBytes(valueSize))
		m.Map[key] = value
	}
	return m
}
