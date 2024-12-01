package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func part2() int {
	dat, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")

	slice1 := []int{}
	slice2 := make(map[int]int)

	for _, item := range items {
		var fst int
		var snd int
		_, e := fmt.Sscan(item, &fst, &snd)
		check(e)

		slice1 = append(slice1, fst)
		slice2[snd] += 1
	}

	slices.Sort(slice1)

	result := 0

	for _, e := range slice1 {
		result += e * slice2[e]
	}

	return result
}
