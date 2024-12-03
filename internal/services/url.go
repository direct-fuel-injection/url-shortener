package services

import (
	"context"

	"github.com/direct-fuel-injection/url-shortener/internal/domain"
	"github.com/google/uuid"
)

type UrlRepository interface {
	GetById(ctx context.Context, id string) (*domain.Url, error) 
	GetByShortUrl(ctx context.Context, shortUrl string) (*domain.Url, error) 
	Create(_ context.Context, url *domain.Url) (*domain.Url, error) 
}

type UrlService struct {
	repository UrlRepository
}

func NewUrlService(repository UrlRepository) UrlService {
	return UrlService{
		repository: repository,
	}
}

func (s UrlService) GetUrl(ctx context.Context, shortUrl string) (*domain.Url, error) {
	return s.repository.GetByShortUrl(ctx, shortUrl)
}

func (s UrlService) CreateUrl(ctx context.Context, longUrl string) (*domain.Url,error) {
	var url *domain.Url
	var err error

	if url, err = domain.NewUrl(uuid.NewString(), longUrl); err != nil {
		return nil, err
	}
	
	if _, err = s.repository.Create(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}