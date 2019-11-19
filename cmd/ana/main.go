package main

import (
	"fmt"
	"os"
	"strings"

	wp "github.com/njvrzm/wordplay/pkg"
)

func main() {
	wl := wp.NewWordList(os.Args[1]).
		FilterOut(wp.Word.OneLetter).
		FilterOut(wp.Word.NotJustLetters)
	var b strings.Builder
	for words := range wp.FindAnagrams(os.Args[2], wl) {
		for _, w := range words {
			b.WriteString(w.ToString())
			b.WriteString(" ")
		}
		fmt.Println(b.String())
		b.Reset()
	}

}
