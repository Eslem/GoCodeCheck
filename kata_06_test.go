package kata

import (
	"testing"
)

func BenchmarkCalculateAnagramSingle(b *testing.B) {
	b.ReportAllocs()
	CalculateAnagram("data/single.txt")
}

func BenchmarkCalculateAnagramSmall(b *testing.B) {
	b.ReportAllocs()
	CalculateAnagram("data/small.txt")
}

func BenchmarkCalculateAnagramList(b *testing.B) {
	b.ReportAllocs()
	CalculateAnagram("data/wordlist.txt")
}

func TestCalculateAnagramSingle(t *testing.T) {
	count, largest := CalculateAnagram("data/single.txt")
	if count != 1 {
		t.Errorf("Error anagram single %d, result %v", count, largest)
	}
}

func TestCalculateAnagramSmall(t *testing.T) {
	count, largest := CalculateAnagram("data/small.txt")
	if count != 7 {
		t.Errorf("Error anagram small %d, result %v", count, largest)
	}
}

func TestCalculateAnagramList(t *testing.T) {
	count, largest := CalculateAnagram("data/wordlist.txt")
	if largest[0] != "alerts" {
		t.Errorf("Error anagram word list count:%d, result %v", count, largest)
	}
}
