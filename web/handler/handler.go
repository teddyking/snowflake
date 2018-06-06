package handler

import (
	"net/http"
)

func New(templateDirPath, staticDirPath string) http.Handler {
	mux := http.NewServeMux()

	indexHandler := NewIndexHandler(templateDirPath, nil)
	fs := http.FileServer(http.Dir(staticDirPath))

	mux.HandleFunc("/", indexHandler.HandleIndex)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
