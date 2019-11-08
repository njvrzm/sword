package wordplay
 
import "testing"

func TestWordIsSubset(t *testing.T) {
    yes_cases := []struct {
        sub, sup string
    }{
        {"hell", "hello"},
        {"yes", "heresy"},
        {"stamp", "antipostmodern"},
    }
    for _, c := range yes_cases {
        if ! NewWord(c.sub).IsSubset(NewWord(c.sup)) {
            t.Errorf("%q should be a subset of %q", c.sub, c.sup)
        }
    }
}
func TestWordIsNotSubset(t *testing.T) {
    no_cases := []struct {
        sub, sup string
    }{
        {"diamond", "hello"},
        {"aa", "heresy"},
        {"postmoderity", "antipostmodern"},
    }
    for _, c := range no_cases {
        if NewWord(c.sub).IsSubset(NewWord(c.sup)) {
            t.Errorf("%q should not be a subset of %q", c.sub, c.sup)
        }
    }
} 

func TestWordMinus(t *testing.T) {
    cases := []struct {
        word, subtrahend, desired string
    }{
        {"word", "do", "wr"},
        {"stipulation", "pistol", "uatin"},
        {"susurrus", "rusts", "urus"},
    }
    for _, c := range cases {
        result := NewWord(c.word).Minus(NewWord(c.subtrahend)).ToString()
        if result != c.desired {
            t.Errorf("%q - %q should be %q but got %q", c.word, c.subtrahend, c.desired, result)
        }
    }
}