package wordplay

import (
	"bufio"
	"os"
	"sort"
)

type WordList []Word

func NewWordList(path string) WordList {
	d := make(WordList, 0)
	d.load(path)
	// Sorted by length to speed anagram search
	sort.Slice(d, func(i, j int) bool {
		return d[i].Len() > d[j].Len()
	})
	return d
}

func (d *WordList) load(path string) {
	fh, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	sc := bufio.NewScanner(fh)
	for sc.Scan() {
		*d = append(*d, NewWord(sc.Text()))
	}
}

type Filter func(Word) bool

func (d WordList) FilterOut(f Filter) (out WordList) {
	for _, w := range d {
		if !f(w) {
			out = append(out, w)
		}
	}
	return
}
