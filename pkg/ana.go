package wordplay

func findAnagrams(words WordList, indices []int, soFar []Word, remaining LetterBag, out chan WordList, depth int) {
	if depth == 0 {
		defer close(out)
	}
	newIndices := make([]int, 0)
	for i := range indices {
		if words[indices[i]].letterCounts.IsSubset(remaining) {
			newIndices = append(newIndices, indices[i])
		}
	}

	for i, index := range newIndices {
		w := words[index]
		soFar[depth] = w
		if w.letterCounts == remaining {
			out <- append(WordList{}, soFar...)
		} else {
			findAnagrams(words,
				newIndices[i:],
				soFar,
				remaining.Minus(w.letterCounts),
				out,
				depth+1)
		}
	}
}

func FindAnagrams(s string, words WordList) chan WordList {
	out := make(chan WordList)
	indices := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		indices[i] = i
	}
	go findAnagrams(words, indices, make(WordList, 10), LetterBagFromString(s), out, 0)
	return out
}
