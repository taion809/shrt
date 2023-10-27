package shrt

import (
	"context"
	"math/rand"
	"strings"
	"sync"

	redis "github.com/redis/go-redis/v9"
)

type Shrt struct {
	Alphabet string
	Cache    *redis.Client

	lock      sync.Mutex
	pool      []string
	poolSize  int
	poolCount int
}

type Shortener interface {
	Generate(size int) []string
	Take(size int) ([]string, error)
}

const defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func New() *Shrt {
	return &Shrt{
		Alphabet: defaultAlphabet,
		pool:     make([]string, 1000),
		poolSize: 1000,
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

func (s *Shrt) Take(size int) ([]string, error) {
	ids := make([]string, size)
	for i := 0; i < size; i++ {
		if s.poolCount == 0 {
			if err := s.updateLocalPool(context.TODO()); err != nil {
				return nil, err
			}
		}

		ids[i] = s.pool[s.poolCount-1]
		s.poolCount--
	}

	return ids, nil
}

func (s *Shrt) updateLocalPool(ctx context.Context) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	list, err := s.Cache.LPopCount(ctx, "pool", s.poolSize).Result()
	if err != nil {
		return err
	}

	s.pool = list
	s.poolCount = len(list)

	return nil
}

func (s *Shrt) UpdateRemotePool(ctx context.Context) error {
	ids := s.Generate(s.poolSize)
	if err := s.Cache.RPush(ctx, "pool", ids).Err(); err != nil {
		return err
	}

	return nil
}
