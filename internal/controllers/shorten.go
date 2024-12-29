package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type ShortenedUrlData struct {
	ShortURL    string
	OriginalURL string
}

type Service struct {
	mapping map[string]string
}

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

	svc := NewService()
	_ = svc.Shorten(originalURL)

	data := ShortenedUrlData{
		ShortURL:    svc.mapping[originalURL],
		OriginalURL: originalURL,
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

func NewService() *Service {
	return &Service{
		mapping: make(map[string]string),
	}
}

func (s *Service) Shorten(originalURL string) error {
	originalURL = strings.TrimSpace(originalURL)
	hashed := hashURL(originalURL)
	shortenedURL := fmt.Sprintf(
		"%s/%s",
		"cust.om",
		hashed,
	)

	s.mapping[originalURL] = shortenedURL
	return nil
}

func hashURL(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))[:8] // Use first 8 characters
}
