package shrt

import (
	"math/rand"
	"strings"
)

type Shrt struct {
	Alphabet string
	Size     int
}

type Shortener interface {
	Shorten(url string) string
}

const defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func New() *Shrt {
	return &Shrt{
		Alphabet: defaultAlphabet,
	}
}

func (s *Shrt) generate(size int) string {
	sb := strings.Builder{}
	sb.Grow(size)
	for i := 0; i < size; i++ {
		sb.WriteByte(s.Alphabet[rand.Intn(len(s.Alphabet))])
	}

	return sb.String()
}

func (s *Shrt) Generate(n int) []string {
	ids := make([]string, n)

	for i := 0; i < n; i++ {
		s := s.generate(5)
		ids[i] = s
	}

	return ids
}
