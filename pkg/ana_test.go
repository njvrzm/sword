package wordplay

import (
	"testing"
)

func BenchmarkFindAnagrams(b *testing.B) {
	words := NewWordList("testdata/basic_850.txt")
	target := "benchmark studies"
	count := 0
	for i := 0; i <= b.N; i++ {
		count = 0
		for range FindAnagrams(target, words) {
			count += 1
		}
	}
}
