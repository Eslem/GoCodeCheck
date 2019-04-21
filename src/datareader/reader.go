package datareader

import (
	"bufio"
	"log"
	"os"
)

func ReadWords(fileName string) []string {
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
