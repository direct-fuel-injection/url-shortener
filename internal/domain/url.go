package domain

import (
	"github.com/deatil/go-encoding/base62"
)

type Url struct {
	id string
	shortUrl string
	longUrl string
}


func NewUrl(id, longUrl string) (*Url, error) {
	url := &Url{
		id: id,
	}

	if err := url.SetLongUrl((longUrl)); err != nil {
		return nil, err
	}
	
	return url, nil
}


func (url *Url) ID() string {
	return url.id
}

func (url *Url) ShortUrl() string {
	return url.shortUrl
}

func (url *Url) LongUrl() string {
	return url.longUrl
}


func (url *Url) SetLongUrl(longUrl string) error {
	if err := validateLongUrl(longUrl); err != nil {
		return err
	}

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	url.longUrl = longUrl
	url.shortUrl = base62.NewEncoding(alphabet).EncodeToString([]byte(longUrl))

	return nil
}