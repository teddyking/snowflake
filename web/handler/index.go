package handler

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
)

//go:generate counterfeiter . FlakerService
type FlakerService interface {
	List(ctx context.Context, in *api.FlakerListReq, opts ...grpc.CallOption) (*api.FlakerListRes, error)
}

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
	flakerListRes, err := h.flakerService.List(context.Background(), &api.FlakerListReq{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles(filepath.Join(h.templateDirPath, "index.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, flakerListRes.Flakes)
}
