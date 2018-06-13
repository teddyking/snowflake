package handler

import (
	"html/template"
	"strings"
	"time"
)

var CustomTemplateFuncs = template.FuncMap{
	"humanizeTime":           humanizeTime,
	"nl2br":                  nl2br,
	"codebaseFromImportPath": codebaseFromImportPath,
	"trimCommit":             trimCommit,
}

func humanizeTime(i int64) string {
	t := time.Unix(i, 0)
	return t.Format(time.RFC1123)
}

func nl2br(s string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(s), "\n", "<br />", -1))
}

func codebaseFromImportPath(s string) string {
	segments := strings.Split(s, "/")
	maxSegments := 3

	if len(segments) < maxSegments {
		maxSegments = len(segments)
	}

	return strings.Join(segments[:maxSegments], "/")
}

func trimCommit(s string) string {
	maxChars := 8

	if len(s) < maxChars {
		maxChars = len(s)
	}

	return s[:maxChars]
}
