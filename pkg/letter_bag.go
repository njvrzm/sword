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
	for i, count := range this {
		if that[i] < count {
			return false
		}
	}
	return true
}

func (this LetterBag) Minus(that LetterBag) (lb LetterBag) {
	for i, count := range this {
		lb[i] = count - that[i]
	}
	return lb
}

func (this LetterBag) IsEmpty() bool {
	for _, count := range this {
		if count > 0 {
			return false
		}
	}
	return true
}
