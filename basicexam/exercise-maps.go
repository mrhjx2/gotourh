package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s struct) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}