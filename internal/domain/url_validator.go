package domain

import (
	"fmt"
	"net/url"
)

func validateLongUrl(longUrl string) error {
	parsedURL, err := url.Parse(longUrl)

	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("%w: longUrl must be a valid URL with a scheme and host", ErrInvalid)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("%w: longUrl must use http or https scheme", ErrInvalid)
	}

	return nil
}