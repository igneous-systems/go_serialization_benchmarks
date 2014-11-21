package goserbench

import "math/rand"

func randBytes(l int) []byte {
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

func randBool() bool {
	return rand.Intn(2) == 1
}
