package transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/direct-fuel-injection/url-shortener/internal/domain"
	"github.com/gorilla/mux"
)

type UrlService interface {
	GetUrl(ctx context.Context, shortUrl string) (*domain.Url, error)
	CreateUrl(ctx context.Context, longUrl string) (*domain.Url, error)
}

type HttpServer struct {
	service UrlService
}

func NewHttpServer(service UrlService) HttpServer {
	return HttpServer{
		service: service,
	}
}

func (h HttpServer) GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["hash"]

	url, err := h.service.GetUrl(r.Context(), shortUrl)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			NotFound("not-found", err, w, r)
			return
		}
		
		fmt.Printf("error: %s ", err)
		RespondWithError(err, w, r)
		return
	}
	
	RedirectPermanent(url.LongUrl(), w, r)
}

func (h HttpServer) CreateUrl(w http.ResponseWriter, r *http.Request) {
	url, err := h.service.CreateUrl(r.Context(), r.URL.Query().Get("url"))

	if err != nil {
		if errors.Is(err, domain.ErrInvalid) {
			BadRequest("incorrect-input", err, w, r)
			return
		}
		
		fmt.Printf("error: %s ", err)
		RespondWithError(err, w, r)
		return
	}
	
	RespondOK(map[string]string{"url": url.ShortUrl() }, w, r)
}