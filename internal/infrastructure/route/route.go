package route

import (
	"github.com/gorilla/mux"
	"go-short-me/internal/representation/handler"
	"net/http"
)

func NewRoute() {
	r := mux.NewRouter()

	r.HandleFunc("/{url}", handler.UrlRedirect).Methods(http.MethodGet)
}
