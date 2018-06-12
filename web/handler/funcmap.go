package handler

import (
	"html/template"
	"strings"
	"time"
)

var CustomTemplateFuncs = template.FuncMap{
	"humanizeTime": humanizeTime,
	"nl2br":        nl2br,
}

func humanizeTime(i int64) string {
	t := time.Unix(i, 0)
	return t.Format(time.RFC1123)
}

func nl2br(s string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(s), "\n", "<br />", -1))
}
