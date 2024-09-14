package handler

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func UrlRedirect(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)["url"]
	logrus.Println(url)

	http.Redirect(w, r, url, http.StatusFound)
}
