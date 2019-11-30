package wordplay

type anagramContext struct {
	words     WordList
	indices   []int
	soFar     []Word
	remaining LetterBag
	out       chan []Word
	depth     int
}

func findAnagrams(c anagramContext) {
	if c.depth == 0 {
		defer close(c.out)
	}
	newIndices := make([]int, 0)
	for i := range c.indices {
		if c.words[c.indices[i]].letterCounts.IsSubset(c.remaining) {
			newIndices = append(newIndices, c.indices[i])
		}
	}

	for i, index := range newIndices {
		w := c.words[index]
		remaining := c.remaining.Minus(w.letterCounts)

		c.soFar[c.depth] = w
		if remaining.IsEmpty() {
			c.out <- append(WordList{}, c.soFar...)
		} else {
			findAnagrams(anagramContext{c.words, newIndices[i:], c.soFar, remaining, c.out, c.depth + 1})
		}
	}
}

func FindAnagrams(s string, wl WordList) chan []Word {
	out := make(chan []Word)
	indices := make([]int, len(wl))
	for i := 0; i < len(wl); i++ {
		indices[i] = i
	}
	go findAnagrams(anagramContext{wl, indices, make([]Word, 10), LetterBagFromString(s), out, 0})
	return out
}
