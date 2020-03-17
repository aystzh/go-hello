package main

import (
	"fmt"
	"github.com/aystzh/go-hello/study/example/calc"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())  // Ahoy, world!
	fmt.Println(calc.Add(10, 3))

}