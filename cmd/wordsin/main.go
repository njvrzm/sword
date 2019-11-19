package main

import (
	"flag"
	"fmt"
	"os"

	wordplay "github.com/njvrzm/wordplay/pkg"
)

func main() {
	minusString := flag.String("minus", "", "a string to remove before looking for words-in")
	flag.CommandLine.Parse(os.Args[3:])

	d := wordplay.NewDictionary(os.Args[1])
	d.FilterOut(wordplay.Word.OneLetter) // every letter is in my usual dictionary
	d.FilterOut(wordplay.Word.NotJustLetters)

	word := wordplay.NewWord(os.Args[2])
	if *minusString != "" {
		word = word.Minus(wordplay.NewWord(*minusString))
	}
	for _, w := range d.Words {
		if w.IsSubset(word) {
			fmt.Println(w.ToString(), word.Minus(w).ToString())
		}
	}
}
