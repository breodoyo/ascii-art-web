package main

func GenerateASCII(text, banner string) (string, error) {
	file := "/banner" + banner + ".txt"

	asciiMap, error := ReadBanner(file)
}