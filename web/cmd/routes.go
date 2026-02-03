package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/qoute", qoute)

	return mux
}
