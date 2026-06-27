package main

import (
	"os"
	"fmt"
	"strings"
)

var validBanners = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
}

func LoadBanner(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) < 855 {
		return nil, fmt.Errorf("invalid banner file")
	}

	var banner [][]string

	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}
		banner = append(banner, lines[i:i+8])
	}

	return banner, nil
}

func RenderWord(word string, banner [][]string) (string, error) {
	var result strings.Builder

	for row := 0; row < 8; row++ {
		for _, ch := range word {
			if ch < 32 || ch > 126 {
				return "", ErrInvalidInput
			}

			result.WriteString(banner[ch-32][row])
		}
		result.WriteString("\n")
	}

	return result.String(), nil
}

func GenerateASCII(text, banner string) (string, error) {
	// 1. Reject unknown banners before doing anything else.
	if !validBanners[banner] {
		return "", ErrUnknownBanner
	}

	// Handles non-ASCII characters
	for _, r := range text {
		if r != '\n' && (r < 32 || r > 126) {
			return "", ErrInvalidInput
		}
	}

	// 3. Load the banner 
	bannerFile := "banners/" + banner + ".txt"
	bannerMap, err := LoadBanner(bannerFile)
	if err != nil {
		return "", err
	}

	// 4. Render line by line.
	lines := strings.Split(text, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		ascii, err := RenderWord(line, bannerMap)
		if err != nil {
			return "", err
		}
		result.WriteString(ascii)
	}

	return result.String(), nil
}