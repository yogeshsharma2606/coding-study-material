package main

import (
	"fmt"
	"strings"
)

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		words[i] = strings.Title(w)
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(titleCase("hello world")) // "Hello World"
	fmt.Println(titleCase("sd CSCSCS"))   //
}
