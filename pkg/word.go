package wordplay

type Word struct {
	string
	letterCounts LetterBag
}

func NewWord(s string) Word {
	return Word{s, LetterBagFromString(s)}
}

func (w Word) ToString() string {
	return w.string
}

func (w Word) Len() int {
	return len(w.string)
}

func (this Word) NotJustLetters() bool {
	for _, r := range this.ToString() {
		if !('a' <= r && r <= 'z') {
			return true
		}
	}
	return false
}

func (this Word) OneLetter() bool {
	word := this.ToString()
	if len(word) == 1 && word != "a" && word != "i" {
		return true
	}
	return false
}
