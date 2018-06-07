package handler

import (
	"net/http"
)

func New(templateDirPath, staticDirPath string, flakerService FlakerService) http.Handler {
	mux := http.NewServeMux()

	indexHandler := NewIndexHandler(templateDirPath, flakerService)
	fs := http.FileServer(http.Dir(staticDirPath))

	mux.HandleFunc("/", indexHandler.HandleIndex)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
