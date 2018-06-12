package handler

import (
	"html/template"
	"strings"
)

var CustomTemplateFuncs = template.FuncMap{
	"nl2br": nl2br,
}

func nl2br(s string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(s), "\n", "<br />", -1))
}
