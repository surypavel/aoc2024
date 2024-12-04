package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func safe_access(items []string, x int, y int) byte {
	if 0 <= x && x < len(items) && 0 <= y && y < len(items[x]) {
		return items[x][y]
	}
	return '0'
}

func lookup(items []string, i int, j int, dx int, dy int, max int) string {
	str := make([]byte, 0)
	for z := 0; z < max; z += 1 {
		str = append(str, safe_access(items, i+z*dx, j+z*dy))
	}
	return string(str)
}

func part1() int {
	dat, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")
	dirs := [3]int{-1, 0, 1}

	count := 0

	for i, row := range items {
		for j := range row {
			for _, dx := range dirs {
				for _, dy := range dirs {
					l := lookup(items, i, j, dx, dy, 4)
					if l == "XMAS" {
						count += 1
					}
				}
			}
		}
	}

	return count
}
