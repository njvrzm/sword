package wordplay

type anagramContext struct {
	words     WordList
	soFar     []Word
	remaining Word
	out       chan []Word
	top       bool
}

func findAnagrams(c anagramContext) {
	if c.top {
		defer close(c.out)
	}
	words := c.words.FilterOut(func(w Word) bool {
		return !w.IsSubset(c.remaining)
	})

	for i, w := range words {
		remaining := c.remaining.Minus(w)

		soFar := make(WordList, len(c.soFar), len(c.soFar)+1)
		copy(soFar, c.soFar)
		soFar = append(soFar, w)
		if remaining.IsEmpty() {
			c.out <- soFar
		} else {
			findAnagrams(anagramContext{words[i:], soFar, remaining, c.out, false})
		}
	}
}

func FindAnagrams(s string, wl WordList) chan []Word {
	out := make(chan []Word)
	go findAnagrams(anagramContext{wl, make([]Word, 0), NewWord(s), out, true})
	return out
}
