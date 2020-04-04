package crandint

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	rand "math/rand"
)

// IntGenerator ...
type IntGenerator struct {
	rand *rand.Rand
}

//CryptoSource ...
type CryptoSource struct{}

// Seed for rand
func (s CryptoSource) Seed(seed int64) {}

//Int63 for rand...
func (s CryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

// Uint64 for rand...
func (s CryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// GetRandomInt generate random (crypto/rand) int
func (i IntGenerator) GetRandomInt(max int) int {
	return i.rand.Intn(max)
}

// New return *IntGenerator
func New() *IntGenerator {
	return &IntGenerator{rand: rand.New(CryptoSource{})}
}

// Get return random int
func Get(max int) int {
	return New().GetRandomInt(max)
}
