package shrt

import (
	"testing"
)

func TestShortenReturnsShort(t *testing.T) {
	s := New()
	s.Size = 3
	short := s.Shorten("http://example.com")
	if len(short) != 3 {
		t.Errorf("Shorten() = %s; want 3 characters", short)
	}

	t.Log(short)
}
