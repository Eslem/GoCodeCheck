package kata

import (
	"testing"
)

func BenchmarkCalculateAnagramSingle(b *testing.B) {
	b.ReportAllocs()
	calculateAnagram("data/single.txt")
}

func BenchmarkCalculateAnagramSmall(b *testing.B) {
	b.ReportAllocs()
	calculateAnagram("data/small.txt")
}

func BenchmarkCalculateAnagramList(b *testing.B) {
	b.ReportAllocs()
	calculateAnagram("data/wordlist.txt")
}
