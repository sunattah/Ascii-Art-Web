package main

import (
	"html/template"
	"net/http"
	"strings"
)

var tmpl *template.Template

func main() {
	// Parse templates at startup. If it fails, panic early.
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		panic("Failed to load templates: " + err.Error())
	}

	// Register the required HTTP Endpoints
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleASCII)

	// Start the server
	println("Server starting on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// GET / Handler
func handleHome(w http.ResponseWriter, r *http.Request) {
	// Strict check: If they type a random path like /abc, return 404 Not Found
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // 200 OK
	tmpl.Execute(w, nil)
}

// POST /ascii-art Handler
func handleASCII(w http.ResponseWriter, r *http.Request) {
	// Check 1: Enforce POST method only. Incorrect methods get 400 Bad Request
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request - Method Not Allowed", http.StatusBadRequest)
		return
	}

	// Read form values from HTML tags
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Check 2: Validate banner input data to prevent directory traversal attacks
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "400 Bad Request - Invalid Banner Selection", http.StatusBadRequest)
		return
	}

	// Process the text using your old logic
	asciiResult, err := GenerateASCII(text, banner)
	if err != nil {
		// If banner file is completely missing, it's a 404 or 500 error
		if strings.Contains(err.Error(), "no such file") {
			http.Error(w, "404 Banner Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Prepare data to send back to the home page view template
	data := PageData{
		InputText: text,
		Result:    asciiResult,
	}

	w.WriteHeader(http.StatusOK) // 200 OK
	tmpl.Execute(w, data)
}
