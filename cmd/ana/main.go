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
		for i, w := range words {
			if i > 0 {
				b.WriteString(" ")
			}
			b.WriteString(w.ToString())
		}
		fmt.Println(b.String())
		b.Reset()
	}

}
