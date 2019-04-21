package gocodecheck
import (
	"testing"
)

var dataPath = "./data"

func BenchmarkCalculateAnagramSingle(b *testing.B) {
	b.ReportAllocs()
	CalculateAnagram(dataPath + "/single.txt")
}

func BenchmarkCalculateAnagramSmall(b *testing.B) {
	b.ReportAllocs()
	CalculateAnagram(dataPath + "/small.txt")
}

func BenchmarkCalculateAnagramList(b *testing.B) {
	b.ReportAllocs()
	CalculateAnagram(dataPath + "/wordlist.txt")
}

func TestCalculateAnagramSingle(t *testing.T) {
	count, largest := CalculateAnagram(dataPath + "/single.txt")
	if count != 1 {
		t.Errorf("Error anagram single %d, result %v", count, largest)
	}
}

func TestCalculateAnagramSmall(t *testing.T) {
	count, largest := CalculateAnagram(dataPath + "/small.txt")
	if count != 7 {
		t.Errorf("Error anagram small %d, result %v", count, largest)
	}
}

func TestCalculateAnagramList(t *testing.T) {
	count, largest := CalculateAnagram(dataPath + "/wordlist.txt")
	if largest[0] != "alerts" {
		t.Errorf("Error anagram word list count:%d, result %v", count, largest)
	}
}
