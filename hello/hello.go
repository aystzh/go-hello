package main

import "fmt"
import "github.com/aystzh/go-hello/hello/morestrings"
import "github.com/google/go-cmp/cmp"

func main() {
	fmt.Println("update")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
