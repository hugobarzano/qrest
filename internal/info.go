package internal

import (
	"fmt"
	"net/http"
	"text/template"
)

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}
	t.Execute(w, "")
}
