package main

import (
	"fmt"
	"os"
	"strings"
)

type Schema struct {
	heights []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_schema(lines []string) Schema {
	heights := []int{-1, -1, -1, -1, -1}
	for _, line := range lines {
		for i, item := range line {
			if item == '#' {
				heights[i]++
			}
		}
	}
	return Schema{heights: heights}
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	locks := make([]Schema, 0)
	keys := make([]Schema, 0)

	for _, lock_or_key := range strings.Split(string(input), "\n\n") {
		lines := strings.Split(lock_or_key, "\n")
		schema := parse_schema(lines)
		if lines[0] == "#####" {
			keys = append(keys, schema)
		} else {
			locks = append(locks, schema)

		}
	}

	sum := 0
	for _, k := range keys {
		for _, l := range locks {
			fits := true
			for i := range k.heights {
				if k.heights[i]+l.heights[i] > 5 {
					fits = false
				}
			}
			if fits {
				sum++
			}
		}
	}

	fmt.Print("part 1 - ", sum, "\n")
}
