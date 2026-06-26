package main

import (
	"errors"
	"html/template"
	"net/http"
)

type PageData struct {
    Input  string
    Banner string
    Result string
}

var bannerNames = []string{"standard", "shadow", "thinkertoy"}

var (
    ErrInvalidInput = errors.New("contains unprintable characters")
    ErrUnknownBanner = errors.New("unknown banner")
)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	
	tmpl, err := template.ParseFiles("templates/index.html") 
		if err != nil {
			http.Error(w, "templates not found", http.StatusNotFound)
			return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "No text added", http.StatusBadRequest)
		return
	}

	//handling errors
	asciiText, err := GenerateASCII(text,banner) 

	if err != nil {
		switch {
		case errors.Is(err,ErrInvalidInput), errors.Is(err,ErrUnknownBanner):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	data := PageData {
		Input: text,
		Banner: banner,
		Result: asciiText,
	}
	tmpl, err := template.ParseFiles("templates/index.html") 
		if err != nil {
			http.Error(w, "templates not found", http.StatusNotFound)
			return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}