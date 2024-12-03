package inmem

import (
	"context"
	"sync"

	"github.com/direct-fuel-injection/url-shortener/internal/domain"
)

type UrlStore struct {
	data map[string]*Url
	mu sync.RWMutex
}

func NewUrlStore() *UrlStore {
	return &UrlStore{
		data: make(map[string]*Url),
	}
}

func (s *UrlStore) GetById(_ context.Context, id string) (*domain.Url, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, exists := s.data[id]

	if !exists {
		return nil, domain.ErrNotFound
	}

	return urlStoreToDomain(url)
}

func (s *UrlStore) GetByShortUrl(_ context.Context, shortUrl string) (*domain.Url, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// TODO: make index by shortUrl
	for _, url := range s.data {
		if url.shortUrl == shortUrl {
			return urlStoreToDomain(url)
		}
	}

	return nil, domain.ErrNotFound
}

func (s *UrlStore) Create(_ context.Context, url *domain.Url) (*domain.Url, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if newUrl, err := urlDomainToStore(url); err != nil {
		return nil, err
	} else {
		s.data[newUrl.id] = newUrl
		return urlStoreToDomain(newUrl)
	}
}