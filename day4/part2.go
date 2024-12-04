package main

import (
	"os"
	"strings"
)

func part2() int {
	dat, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")

	count := 0

	for i, row := range items {
		for j := range row {
			l1 := safe_access(items, i-1, j-1)
			l2 := safe_access(items, i-1, j+1)
			l3 := safe_access(items, i+1, j-1)
			l4 := safe_access(items, i+1, j+1)
			l5 := safe_access(items, i, j)

			if l5 == 'A' && ((l1 == 'M' && l4 == 'S') || (l1 == 'S' && l4 == 'M')) && ((l2 == 'M' && l3 == 'S') || (l2 == 'S' && l3 == 'M')) {
				count += 1
			}
		}
	}

	return count
}
