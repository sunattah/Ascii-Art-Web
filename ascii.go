package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(name string) ([]string, error) {
	data, err := os.ReadFile(name + ".txt")
	if err != nil {
		return nil, fmt.Errorf("banner not found: %s", name)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n"), nil
}

func Render(text string, banner []string) (string, error) {
	var out strings.Builder
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			out.WriteString("\n")
			continue
		}
		rows := make([]string, 8)
		for _, ch := range line {
			if ch < 32 || ch > 126 {
				return "", fmt.Errorf("character out of range: %q", ch)
			}
			start := (int(ch)-32)*9 + 1
			for i := range rows {
				rows[i] += banner[start+i]
			}
		}
		for _, row := range rows {
			out.WriteString(row)
			out.WriteString("\n")
		}
	}
	return out.String(), nil
}
