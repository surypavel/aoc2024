package main

import (
	"os"
	"strings"
)

func FindWrongIndex[T any](ts []T, pred func(T) bool) int {
	for i, t := range ts {
		if !pred(t) {
			return i
		}
	}
	return -1
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func is_valid_2(measurements []int) bool {
	differences := []int{}

	for i := 0; i < len(measurements)-1; i++ {
		differences = append(differences, measurements[i]-measurements[i+1])
	}

	fn1 := func(b int) bool { return b > 0 && b < 4 }
	fn2 := func(b int) bool { return b < 0 && b > -4 }

	// Finds the first wrong index for every direction
	wrongIndex1 := FindWrongIndex(differences, fn1)
	wrongIndex2 := FindWrongIndex(differences, fn2)

	// Check if it is either ok or it is ok if the first suspicious index gets removed
	// There are two options which to remove (the wrong one or the preceding)
	return wrongIndex1 == -1 || wrongIndex2 == -1 || is_valid(RemoveIndex(measurements, wrongIndex1)) || is_valid(RemoveIndex(measurements, wrongIndex2)) || is_valid(RemoveIndex(measurements, wrongIndex1+1)) || is_valid(RemoveIndex(measurements, wrongIndex2+1))
}

func part2() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")
	result := 0

	for _, item := range items {
		measurements := ToInts(strings.Split(item, " "))
		if is_valid_2(measurements) {
			result += 1
		}
	}

	return result
}
