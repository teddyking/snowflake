package handler

import (
	"net/http"
	"path/filepath"
)

func New(staticDirPath string, flakerService FlakerService) http.Handler {
	mux := http.NewServeMux()

	templatesDirPath := filepath.Join(staticDirPath, "templates")
	indexHandler := NewIndexHandler(templatesDirPath, flakerService)
	fs := http.FileServer(http.Dir(staticDirPath))

	mux.HandleFunc("/", indexHandler.HandleIndex)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
