package wordplay

type anagramContext struct {
	words     WordList
	soFar     []Word
	remaining Word
	out       chan []Word
	depth     int
}

func findAnagrams(c anagramContext) {
	if c.depth == 0 {
		defer close(c.out)
	}
	words := c.words.FilterOut(func(w Word) bool {
		return !w.IsSubset(c.remaining)
	})

	for i, w := range words {
		remaining := c.remaining.Minus(w)

		c.soFar[c.depth] = w
		if remaining.IsEmpty() {
			anagram := make(WordList, c.depth+1)
			copy(anagram, c.soFar)
			c.out <- anagram
		} else {
			findAnagrams(anagramContext{words[i:], c.soFar, remaining, c.out, c.depth + 1})
		}
	}
}

func FindAnagrams(s string, wl WordList) chan []Word {
	out := make(chan []Word)
	go findAnagrams(anagramContext{wl, make([]Word, 10), NewWord(s), out, 0})
	return out
}
