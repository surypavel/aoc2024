package main

import (
	"os"
	"regexp"
	"strconv"
)

func part2() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	result := 0
	enabled := true

	r := regexp.MustCompile(`(don't)\(\)|(do)\(\)|(mul)\(([0-9]+),([0-9]+)\)`)
	matches := r.FindAllStringSubmatch(string(dat), -1)
	for _, v := range matches {
		if v[1] == "don't" {
			enabled = false
		}
		if v[2] == "do" {
			enabled = true
		}
		if v[3] == "mul" && enabled {
			num1, err1 := strconv.Atoi(v[4])
			num2, err2 := strconv.Atoi(v[5])

			check(err1)
			check(err2)

			result += num1 * num2
		}
	}

	return result
}
