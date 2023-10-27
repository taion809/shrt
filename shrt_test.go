package shrt

import (
	"testing"
)

func TestShortenReturnsShort(t *testing.T) {
	s := New()
	short := s.Generate(10)

	t.Log(short)
}

func BenchmarkGenerate(b *testing.B) {
	s := New()
	for n := 0; n < b.N; n++ {
		s.Generate(10)
	}
}
