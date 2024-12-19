package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func attainable(pattern string, towels []string) bool {
	if pattern == "" {
		return true
	} else {
		for _, towel := range towels {
			if strings.HasPrefix(pattern, towel) {
				if attainable(pattern[len(towel):], towels) {
					return true
				}
			}
		}
		return false
	}
}

func main() {
	input, err := os.ReadFile("example.txt")
	check(err)

	input_parts := strings.Split(string(input), "\n\n")
	towels := strings.Split(input_parts[0], ", ")
	patterns := strings.Split(input_parts[1], "\n")

	sum := 0
	for _, pattern := range patterns {
		if attainable(pattern, towels) {
			sum++
		}
	}

	fmt.Print("part 1 - ", sum, "\n")
}
