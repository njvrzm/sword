package wordplay

type LetterBag [26]int8

func LetterBagFromString(s string) (lb LetterBag) {
	for _, c := range s {
		if 'A' <= c && c <= 'Z' {
			lb[c-'A']++
		} else if 'a' <= c && c <= 'z' {
			lb[c-'a']++
		}
	}
	return
}

func (this LetterBag) IsSubset(that LetterBag) bool {
	for i := 0; i < 26; i++ {
		if that[i] < this[i] {
			return false
		}
	}
	return true
}

func (this LetterBag) Minus(that LetterBag) (lb LetterBag) {
	for i := 0; i < 26; i++ {
		lb[i] = this[i] - that[i]
	}
	return lb
}
