package web

import (
	"net/http"

	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/template/index.html")
	t.Execute(w, nil)
}
