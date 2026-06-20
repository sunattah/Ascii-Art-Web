package main

import (
	"os"
	"strings"
)

type PageData struct {
	InputText string
	Result    string
	Error     string
}

func GenerateASCII(text, banner string) (string, error) {
	filePath := "banners/" + banner + ".txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(content), "\n")
	_ = lines

	return "Sample ASCII output for: " + text, nil
}
