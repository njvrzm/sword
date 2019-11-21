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

	for i, w := range c.words {
		if w.IsSubset(c.remaining) {
			remaining := c.remaining.Minus(w)

			soFar := make([]Word, len(c.soFar)+1)
			copy(soFar, append(c.soFar, w))
			if remaining.IsEmpty() {
				c.out <- soFar
			} else {
				findAnagrams(anagramContext{c.words[i:], soFar, remaining, c.out, false})
			}
		}
	}
}
func FindAnagrams(s string, wl WordList) chan []Word {
	out := make(chan []Word)
	go findAnagrams(anagramContext{wl, make([]Word, 0), NewWord(s), out, true})
	return out
}
