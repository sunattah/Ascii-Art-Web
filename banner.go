package main

import (
	"fmt"
	"os"
	"strings"
)

func loadbanner(input string) (map[rune][]string, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return nil, fmt.Errorf("empty file read %s", err)
	}
	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	bannermap := map[rune][]string{}
	currentchar := rune(32)

	for i := 1; i+8 < len(lines); i += 9 {
		bannermap[currentchar] = lines[i : i+8]
		currentchar++
	}
	return nil, err
}
func render(input string, banner map[rune][]string) []string {
	bannerslice := make([]string, 8)
	for i := 0; i < 8; i++ {
		for _, ch := range input {
			bannerslice = banner[ch]
		}
	}
	return bannerslice

}
func slice(input string) []string {
	line := strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(line, "\n")
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage go run . [TEXT] [BANNER]")
	}
	input := os.Args[1]
	banner := "standard"
	if len(os.Args) == 3 {
		banner = os.Args[2] + "txt"
	}
	s, _ := loadbanner(banner)
	// if err !=nil {
	// 	fmt.Println(s)
	// }
	lines := slice(input)
	for _, ch := range lines {
		if ch == "" {
			fmt.Println()
		}
		renderline := render(input, s)
		for _, char := range renderline {
			fmt.Println(char)
		}
	}
}
