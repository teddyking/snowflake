package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
)

//go:generate counterfeiter . FlakerService
type FlakerService interface {
	List(ctx context.Context, in *api.FlakerListReq, opts ...grpc.CallOption) (*api.FlakerListRes, error)
}

type IndexHandler struct {
	templates     *template.Template
	flakerService FlakerService
}

func NewIndexHandler(templateDirPath string, flakerService FlakerService) *IndexHandler {
	templatesGlob := fmt.Sprintf("%s/*.html", templateDirPath)
	parsedTemplates := template.Must(
		template.New("").Funcs(CustomTemplateFuncs).ParseGlob(templatesGlob),
	)

	return &IndexHandler{
		templates:     parsedTemplates,
		flakerService: flakerService,
	}
}

func (h *IndexHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	flakerListRes, err := h.flakerService.List(context.Background(), &api.FlakerListReq{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.templates.ExecuteTemplate(w, "index.html", flakerListRes.Flakes)
}
