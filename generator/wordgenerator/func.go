package wordgenerator

import (
	"github.com/regemdeu/algorithms/generator/crandint"
)

// New create WordGenerator instance
func New(minwl, maxwl int) *WordGenerator {
	return &WordGenerator{
		MaxWordLen:   maxwl,
		MinWordLen:   minwl,
		Alphabets:    GetAlphabets(),
		IntGenerator: crandint.New(),
	}
}

// GetAlphabets return alphabets
func GetAlphabets() []rune {
	len := 26

	alphabets := make([]rune, len)

	for i := 0; i < len; i++ {
		alphabets[i] = rune(97 + i)
	}

	return alphabets
}
