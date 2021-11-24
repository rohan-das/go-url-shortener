package main

import (
	"go-url-shortner/handlers"
	"go-url-shortner/stores"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	store := stores.New()
	handler := handlers.New(store)

	r := mux.NewRouter()

	r.HandleFunc("/short", handler.GetShortURL)

	http.ListenAndServe(":8080", r)
}
