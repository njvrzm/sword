package main
import (
    "github.com/njvrzm/wordplay/pkg"
    "os"
    "fmt"
    "flag"
)

func oneLetter(w wordplay.Word) bool {
    word := w.ToString()
    if (len(word) == 1 && word != "a" && word != "i") {
        return true
    }
    return false
}

func main() {
    d := wordplay.NewDictionary(os.Args[1])
    minusString := flag.String("minus", "", "a string to remove before looking for words-in")
    flag.CommandLine.Parse(os.Args[3:])
    d.FilterOut(oneLetter) // every letter is in my usual dictionary
    word := wordplay.NewWord(os.Args[2])
    if *minusString != "" {
        word = word.Minus(wordplay.NewWord(*minusString))
    }
    for _, w := range d.Words {
        if w.IsSubset(word) {
            fmt.Println( w.ToString(), word.Minus(w).ToString())
        }
    }
}
