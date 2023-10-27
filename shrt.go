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
	return &Shrt{}
}

func (s *Shrt) Shorten(url string) string {
	if s.Alphabet == "" {
		s.Alphabet = defaultAlphabet
	}

	if s.Size == 0 {
		s.Size = 5
	}

	sb := strings.Builder{}
	sb.Grow(s.Size)
	for i := 0; i < s.Size; i++ {
		sb.WriteByte(s.Alphabet[rand.Intn(len(s.Alphabet))])
	}

	return sb.String()
}
