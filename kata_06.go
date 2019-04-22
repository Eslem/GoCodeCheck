package gocodecheck

import (
	"time"

	"github.com/Eslem/GoCodeCheck/utils"
)

// CalculateAnagram returns the total count and the largest anagram in a file
func CalculateAnagram(fileName string) (int, []string) {
	defer utils.TimeTrack(time.Now(), "calculate anagram "+fileName)
	words := utils.ReadLines(fileName)
	anagrams := make(map[string][]string)

	var largest string
	largestLen := 0
	count := 0

	for _, word := range words {
		sorted := utils.SortAndLowercase(word)
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

	return count, anagrams[largest]
}
