package inmem

import (
	"time"

	"github.com/direct-fuel-injection/url-shortener/internal/domain"
)

type Url struct {
	id string
	shortUrl string
	longUrl string

	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func (u *Url) Copy() *Url {
	if u == nil {
		return nil
	}

	return &Url{
		id: u.id,
		shortUrl: u.shortUrl,
		longUrl: u.longUrl,

		createdAt: u.createdAt,
		updatedAt: u.updatedAt,
		deletedAt: u.deletedAt,
	}
}

func urlStoreToDomain(url *Url) (*domain.Url, error) {
	if url == nil {
		return nil, domain.ErrNil
	}

	return domain.NewUrl(url.id, url.longUrl)
}

func urlDomainToStore(url *domain.Url) (*Url, error)  {
	return &Url{
		id: url.ID(),
		shortUrl: url.ShortUrl(),
		longUrl: url.LongUrl(),

		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}