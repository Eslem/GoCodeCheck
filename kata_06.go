package gocodecheck

import (
	"log"
	"time"
)

// CalculateAnagram returns the total count and the largest anagram in a file
func CalculateAnagram(fileName string) (int, []string) {
	defer TimeTrack(time.Now(), "calculate anagram "+fileName)
	words := ReadWords(fileName)
	anagrams := make(map[string][]string)

	var largest string
	largestLen := 0
	count := 0

	for _, word := range words {
		sorted := SortAndLowercase(word)
		_, ok := anagrams[sorted]
		if !ok {
			anagrams[sorted] = []string{word}
		} else {
			if len(anagrams[sorted]) == 1 {
				count++
			}
			anagrams[sorted] = append(anagrams[sorted], word)
		}

		if len(anagrams[sorted]) > largestLen {
			largest = sorted
			largestLen = len(anagrams[sorted])
		}
	}

	log.Printf("count %d \n", count)
	log.Println("Largest", anagrams[largest][0])
	return count, anagrams[largest]
}
