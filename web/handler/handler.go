package handler

import (
	"net/http"
)

func New(staticDirPath string) http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(staticDirPath))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
