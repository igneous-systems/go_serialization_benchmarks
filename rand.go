package goserbench

// Utilities for generating cryptographically random test data.

import (
	cryptorand "crypto/rand"
	"math/big"
	mathrand "math/rand"
)

// This exposes the convenience methods of the math/rand package,
// but with the cryptographic randomness of the crypto/rand package.
var crand = NewCryptoRand()

// Wrapper around math/rand.Rand to add utility methods
type CryptoRand struct {
	*mathrand.Rand
}

func NewCryptoRand() *CryptoRand {
	return &CryptoRand{
		mathrand.New(CryptoSource{}),
	}
}

func (r CryptoRand) Bytes(l int) []byte {
	b := make([]byte, l)
	_, err := cryptorand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func (r CryptoRand) Bool() bool {
	return r.Intn(2) == 1
}

type CryptoSource struct{}

func (s CryptoSource) Int63() int64 {
	max := &big.Int{}
	max.SetUint64(1 << 63)
	resultBig, err := cryptorand.Int(cryptorand.Reader, max)
	if err != nil {
		panic(err)
	}
	return resultBig.Int64()
}

// CryptoSource does not require a seed for randomness,
// so do nothing with this piece of the Source interface.
func (s CryptoSource) Seed(seed int64) {}
