package handlers

import (
	"encoding/json"
	"go-url-shortner/models"
	"go-url-shortner/stores"
	"log"
	"net/http"
)

type urlHandler struct {
	store stores.URL
}

func New(store stores.URL) *urlHandler {
	return &urlHandler{store: store}
}

func (u *urlHandler) GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.MyURL

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	shortUrl, err := u.store.GetShortURL(url)
	if err != nil {
		log.Println(err.Error())
		writeResponse(w, http.StatusInternalServerError, "some unexpected error occured")
		return
	}

	writeResponse(w, http.StatusOK, shortUrl)
}

func writeResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
