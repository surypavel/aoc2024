package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

func ToInts(ts []string) []int {
	output := []int{}
	for _, t := range ts {
		num, err := strconv.Atoi(t)
		check(err)
		output = append(output, num)
	}
	return output
}

func is_valid(measurements []int) bool {
	differences := []int{}

	for i := 0; i < len(measurements)-1; i++ {
		differences = append(differences, measurements[i]-measurements[i+1])
	}

	fn1 := func(b int) bool { return b > 0 && b < 4 }
	fn2 := func(b int) bool { return b < 0 && b > -4 }
	return All(differences, fn1) || All(differences, fn2)
}

func part1() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")
	result := 0

	for _, item := range items {
		measurements := ToInts(strings.Split(item, " "))
		if is_valid(measurements) {
			result += 1
		}
	}

	return result
}
