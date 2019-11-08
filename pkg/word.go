package wordplay

import "bytes"

type Word struct {
    string
    count map[rune]int
}

func NewWord (s string) Word {
    return Word{string:s}
}

func (w Word) ToString() string {
    return w.string
}
func (w Word) Length() int {
    return len(w.string)
}
func (w *Word) LetterCounts() map[rune]int {
    if (w.count == nil) {
        w.count = make(map[rune]int)
        for _, c := range(w.string) {
            w.count[c] += 1
        }
    }
    return w.count
}

func (w *Word) LetterCount(r rune) int {
    return w.LetterCounts()[r]

}

func (this Word) IsSubset(that Word) bool {
    for ru, count := range(this.LetterCounts()) {
        if (that.LetterCount(ru) < count) {
            return false
        }
    }
    return true
}

func (this Word) Minus(that Word) Word {
    var out bytes.Buffer
    toRemove := that.LetterCounts()
    for _, r := range(this.ToString()) {
        if (toRemove[r] > 0) {
            toRemove[r] -= 1
        } else {
            out.WriteRune(r)
        }
    }
    return NewWord(out.String())

}

func (this Word) Intersects(that Word) bool {
    for ru, count := range(this.LetterCounts()) {
        if (count > 0 && that.LetterCount(ru) == 0) {
            return false
        }
    }
    return true

}

func (this Word) Equals(that Word) bool {
    return this.string == that.string
}

// func main() {
//     w := NewWord(os.Args[1])
//     y := NewWord("hello and goodbye")
//     fmt.Println(w.IsSubset(y))
//     fmt.Println(y.IsSubset(w))
// }