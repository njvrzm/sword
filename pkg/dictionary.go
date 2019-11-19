package wordplay

import (
	"bufio"
	"os"
	"sort"
)

type Dictionary struct {
	path  string
	Words []Word
}

func NewDictionary(path string) Dictionary {
	d := Dictionary{path: path, Words: make([]Word, 0)}
	d.Reset()
	sort.Slice(d.Words, func(i, j int) bool {
		return d.Words[i].Len() > d.Words[j].Len()
	})
	return d
}

func (d *Dictionary) Reset() {
	fh, err := os.Open(d.path)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	sc := bufio.NewScanner(fh)
	for sc.Scan() {
		d.Words = append(d.Words, NewWord(sc.Text()))
	}
}

type Filter func(Word) bool

func (d *Dictionary) FilterOut(f Filter) {
	toScan := d.Words
	d.Words = d.Words[:0]
	for _, w := range toScan {
		if !f(w) {
			d.Words = append(d.Words, w)
		}
	}
}

func (d Dictionary) Index(word Word) int {
	for i, w := range d.Words {
		if word.Equals(w) {
			return i
		}
	}
	return -1
}

func (d Dictionary) Len() int {
	return len(d.Words)
}
