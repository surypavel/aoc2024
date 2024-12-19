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

func count_attainable(pattern string, towels []string, sum int, cache map[string]int) int {
	if pattern == "" {
		return 1
	} else {
		count := -1
		if cache[pattern] != 0 {
			count = cache[pattern]
		} else {
			count = 0
			for _, towel := range towels {
				if strings.HasPrefix(pattern, towel) {
					count += count_attainable(pattern[len(towel):], towels, 1, cache)
				}
			}
			cache[pattern] = count
		}

		return sum * count
	}
}

func is_attainable(pattern string, towels []string) bool {
	if pattern == "" {
		return true
	} else {
		for _, towel := range towels {
			if strings.HasPrefix(pattern, towel) {
				if is_attainable(pattern[len(towel):], towels) {
					return true
				}
			}
		}
		return false
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	input_parts := strings.Split(string(input), "\n\n")
	towels := strings.Split(input_parts[0], ", ")
	patterns := strings.Split(input_parts[1], "\n")

	sum_1 := 0
	for _, pattern := range patterns {
		if is_attainable(pattern, towels) {
			sum_1++
		}
	}

	fmt.Print("part 1 - ", sum_1, "\n")

	sum_2 := 0
	cache := make(map[string]int)
	for _, pattern := range patterns {
		sum_2 += count_attainable(pattern, towels, 1, cache)
	}

	fmt.Print("part 1 - ", sum_2, "\n")
}
