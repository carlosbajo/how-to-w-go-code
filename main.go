package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
