package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	calc(dat)
}
