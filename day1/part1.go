package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() int {
	dat, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")

	slice1 := []int{}
	slice2 := []int{}

	for _, item := range items {
		var fst int
		var snd int
		_, e := fmt.Sscan(item, &fst, &snd)
		check(e)

		slice1 = append(slice1, fst)
		slice2 = append(slice2, snd)
	}

	slices.Sort(slice1)
	slices.Sort(slice2)

	result := 0

	for i, e := range slice1 {
		result += abs(e - slice2[i])
	}

	return result
}
