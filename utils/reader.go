package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadLines open and file an return an array of it lines
func ReadLines(fileName string) []string {
	file, err := os.Open(fileName)
	var words []string

	if err != nil {
		log.Fatalf("Fail opening files %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	file.Close()
	return words
}
