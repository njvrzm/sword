package main

import (
	"fmt"

	wp "github.com/njvrzm/wordplay/pkg"
)

func main() {
	fmt.Println(wp.NewWord("capitalize scene").Minus(wp.NewWord("capsize")))

}
