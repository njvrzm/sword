package wordplay

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func findAnagrams(words WordList, indices []int, soFar []Word, remaining LetterBag, out chan WordList, depth int, threads int, modulus int) {
	if depth == 0 {
		defer wg.Done()
	}
	newIndices := make([]int, 0)
	for i := range indices {
		if words[indices[i]].letterCounts.IsSubset(remaining) {
			newIndices = append(newIndices, indices[i])
		}
	}

	for i, index := range newIndices {
		if depth == 0 && (index%threads != modulus) {
			continue
		}
		w := words[index]
		soFar[depth] = w
		if w.letterCounts == remaining {
			out <- append(WordList{}, soFar[:depth+1]...)
		} else {
			findAnagrams(words,
				newIndices[i:],
				soFar,
				remaining.Minus(w.letterCounts),
				out,
				depth+1,
				threads,
				modulus,
			)
		}
	}
}

func FindAnagrams(s string, words WordList) chan WordList {
	out := make(chan WordList)
	indices := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		indices[i] = i
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go findAnagrams(words, indices, make(WordList, 10), LetterBagFromString(s), out, 0, runtime.NumCPU(), i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
