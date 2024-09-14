package route

import (
	"github.com/gorilla/mux"
	"go-short-me/internal/infrastructure/middleware"
	"go-short-me/internal/representation/handler"
	"net/http"
)

func NewRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/shortener/", handler.ShortUrl).Methods(http.MethodPost)
	r.HandleFunc("/{url}", handler.UrlRedirect).Methods(http.MethodGet)
	r.Use(middleware.LoggingMiddleware, middleware.RecoveryMiddleware, middleware.ContentTypeMiddleware)

	return r
}
