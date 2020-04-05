package wordgenerator

import (
	"strings"

	"github.com/regemdeu/algorithms/generator/crandint"
)

// WordGenerator ...
type WordGenerator struct {
	Alphabets    []rune
	MinWordLen   int
	MaxWordLen   int
	IntGenerator *crandint.IntGenerator
}

// GetRandomInt generate random (crypto/rand) int
func (wg *WordGenerator) GetRandomInt(max int) int {
	return wg.IntGenerator.GetRandomInt(max)
}

// GenWord generate word
func (wg *WordGenerator) GenWord(l int) string {
	var r string

	for i := 0; i < l; i++ {
		r += string(wg.Alphabets[wg.GetRandomInt(len(wg.Alphabets))])
	}
	return r
}

// GenWords generate words
func (wg *WordGenerator) GenWords(count int) string {
	var s []string

	for i := 0; i < count; i++ {
		wl := wg.GetRandomInt(wg.MaxWordLen)
		if wl < wg.MinWordLen {
			wl = wg.MinWordLen
		}
		s = append(s, wg.GenWord(wl))
	}

	return strings.Join(s, " ")
}
