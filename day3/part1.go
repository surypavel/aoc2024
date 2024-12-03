package main

import (
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() int {
	dat, err := os.ReadFile("./example.txt")
	check(err)

	result := 0

	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := r.FindAllStringSubmatch(string(dat), -1)
	for _, v := range matches {
		num1, err1 := strconv.Atoi(v[1])
		num2, err2 := strconv.Atoi(v[2])

		check(err1)
		check(err2)

		result += num1 * num2
	}

	return result
}
