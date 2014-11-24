package goserbench

import (
	"fmt"
	"math/rand"
)

// Generate a random hex string of length l.
func randString(l int) string {
	// Each byte produces 2 characters.
	numBytes := (l + 1) / 2
	bytes := make([]byte, numBytes)
	for i := 0; i < numBytes; i++ {
		bytes[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", bytes)[:l]
}

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
