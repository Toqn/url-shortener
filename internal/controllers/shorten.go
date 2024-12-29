package controllers

import (
	"html/template"
	"net/http"
	"strings"
)

func ShowShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(originalURL, "http") && !strings.HasPrefix(originalURL, "https") {
		originalURL = "https://" + originalURL
	}

	// implement shorten url

	data := map[string]string{
		"ShortURL": originalURL,
	}

	tmpl, err := template.ParseFiles("internal/template/shorten.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
