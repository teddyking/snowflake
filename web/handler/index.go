package handler

import (
	"net/http"
	"path/filepath"

	"github.com/alecthomas/template"
)

type FlakerService interface{}

type IndexHandler struct {
	templateDirPath string
	flakerService   FlakerService
}

func NewIndexHandler(templateDirPath string, flakerService FlakerService) *IndexHandler {
	return &IndexHandler{
		templateDirPath: templateDirPath,
		flakerService:   flakerService,
	}
}

func (h *IndexHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(h.templateDirPath, "index.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}
