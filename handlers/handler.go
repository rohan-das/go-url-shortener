package handlers

import (
	"encoding/json"
	"go-url-shortner/models"
	"go-url-shortner/stores"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type urlHandler struct {
	store stores.URL
}

func New(store stores.URL) *urlHandler {
	return &urlHandler{store: store}
}

func (u *urlHandler) GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.MyURL

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		writeResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	// to check if input URL is valid or not
	if !govalidator.IsURL(url.LongURL) {
		writeResponse(w, http.StatusBadRequest, "invalid url")
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
