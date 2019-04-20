package kata

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func readWords(fileName string) []string {
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

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s  \n", name, elapsed)
}

func calculateAnagram(fileName string) {
	defer timeTrack(time.Now(), "calculate anagram "+fileName)
	words := readWords(fileName)
	anagrams := make(map[string][]string)

	var largest string
	largestLen := 0

	count := 0

	for _, word := range words {
		sorted := strings.ToLower(SortStringByCharacter(word))
		_, ok := anagrams[sorted]
		if !ok {
			anagrams[sorted] = make([]string, 1)
			anagrams[sorted][0] = word
		} else {
			anagrams[sorted] = append(anagrams[sorted], word)
			count++
		}

		if len(anagrams[sorted]) > largestLen {
			largest = sorted
			largestLen = len(anagrams[sorted])
		}
	}

	log.Printf("count %d \n", count)
	log.Println("Largest", anagrams[largest])
}
