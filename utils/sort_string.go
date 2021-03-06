package utils

import (
	"sort"
	"strings"
)

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// SortAndLowercase Sort and string and change it to lowercase
func SortAndLowercase(text string) string {
	return strings.ToLower(sortStringByCharacter(text))
}
