package handler

import (
	"encoding/json"
	url2 "go-short-me/internal/component/gsm/core/domain/url"
	"go-short-me/internal/infrastructure/db"
	"go-short-me/internal/infrastructure/models"
	"go-short-me/internal/infrastructure/reason"
	"go-short-me/internal/infrastructure/responder"
	"net/http"
)

func ShortUrl(w http.ResponseWriter, r *http.Request) {
	var url models.Url
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = r.Body.Close()
	if err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, reason.GoesWrong)
		return
	}

	newUrl, err := url2.NewUrl(url.Url)
	if err != nil {
		return
	}

	adapter := &db.MockDB{}
	_, err = newUrl.CreateShortURL(adapter)
	if err != nil {
		return
	}

	responder.RespondWithJSON(w, http.StatusOK, url)
}
