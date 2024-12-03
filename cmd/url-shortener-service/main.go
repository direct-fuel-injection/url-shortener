package main

import (
	"log"
	"net/http"
	"os"

	"github.com/direct-fuel-injection/url-shortener/internal/config"
	"github.com/direct-fuel-injection/url-shortener/internal/repository/inmem"
	"github.com/direct-fuel-injection/url-shortener/internal/services"
	"github.com/direct-fuel-injection/url-shortener/internal/transport"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func run() error {
	cfg := config.Read()

	urlStorage := inmem.NewUrlStore()
	urlService := services.NewUrlService(urlStorage)
	httpServer := transport.NewHttpServer(urlService)

	router := mux.NewRouter()
	router.HandleFunc("/create", httpServer.CreateUrl).Methods("GET")
	router.HandleFunc("/{hash}", httpServer.GetUrl).Methods("GET")

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("Server started on http://localhost:%s", cfg.Port)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	return nil
}
