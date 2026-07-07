package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, f := range []string{"standard.txt", "shadow.txt", "thinkertoy.txt", "templates/index.html"} {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			fmt.Println("ERROR: missing file:", f, "\nRun `go run .` from inside the ascii-art-web folder.")
			os.Exit(1)
		}
	}

	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", AsciiArt)

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}