package wordplay

import (
	"bytes"
	"strings"
)

type Word struct {
	string
	letterCounts []int8
}

func NewWord(s string) Word {
	var w Word
	w.string = s
	w.letterCounts = make([]int8, 26)
	for _, c := range w.string {
		if c >= 'a' && c <= 'z' {
			w.letterCounts[c-'a'] += 1
		}
	}
	return w
}

func (w Word) ToString() string {
	return w.string
}

func (w Word) Len() int {
	return len(w.string)
}

func (w Word) IsEmpty() bool {
	for _, c := range w.string {
		if c != ' ' {
			return false
		}
	}
	return true
}

func (w *Word) LetterCount(r rune) int8 {
	if 'a' <= r && r <= 'z' {
		return w.letterCounts[r-'a']
	}
	return 0
}

func (this Word) IsSubset(that Word) bool {
	for i, c := range this.letterCounts {
		if that.letterCounts[i] < c {
			return false
		}
	}
	return true
}

func (this Word) Minus(that Word) Word {
	var out bytes.Buffer
	toRemove := make([]int8, 26)
	copy(toRemove, that.letterCounts)
	for _, r := range this.ToString() {
		if 'a' <= r && r <= 'z' {
			if toRemove[r-'a'] > 0 {
				toRemove[r-'a'] -= 1
			} else {
				out.WriteRune(r)
			}
		}
	}
	return NewWord(out.String())
}

func (this Word) NotJustLetters() bool {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, r := range this.ToString() {
		if !strings.ContainsRune(letters, r) {
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
