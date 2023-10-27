package shrt

type Shrt struct {
}

type Shortener interface {
	Shorten(url string) string
}

func New() *Shrt {
	return &Shrt{}
}

func (s *Shrt) Shorten(url string) string {
	return ""
}
