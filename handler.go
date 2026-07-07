package main

import (
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Result, Error, Text, Banner string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}
	render(w, PageData{Banner: "standard"})
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text, banner := r.FormValue("text"), r.FormValue("banner")
	if banner == "" {
		banner = "standard"
	}

	valid := map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}
	if !valid[banner] {
		http.Error(w, "400 - Bad Request", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(text) == "" {
		render(w, PageData{Error: "Please enter some text.", Banner: banner})
		return
	}

	bannerLines, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, "404 - Banner Not Found", http.StatusNotFound)
		return
	}

	result, err := Render(text, bannerLines)
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}

	render(w, PageData{Result: result, Text: text, Banner: banner})
}

func render(w http.ResponseWriter, data PageData) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "404 - Template Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}