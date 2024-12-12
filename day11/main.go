package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func to_int(s string) int {
	num, err := strconv.Atoi(s)
	check(err)
	return num
}

func expand(items []string, number int) []string {
	if number == 0 {
		return items
	}

	new_items := make([]string, 0)

	for _, item := range items {
		if item == "0" {
			new_items = append(new_items, "1")
		} else if len(item)%2 == 0 {
			mid := len(item) / 2
			new_items = append(new_items, item[:mid])
			new_items = append(new_items, trim_zeroes(item[mid:]))
		} else {
			new_items = append(new_items, trim_zeroes(strconv.Itoa(to_int(item)*2024)))
		}
	}

	return expand(new_items, number-1)
}

func trim_zeroes(input string) string {
	r, _ := regexp.Compile("^0*(0|[1-9][0-9]*)$")
	return r.FindStringSubmatch(input)[1]
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(input), " ")

	fmt.Print("part 1 - ", len(expand(items, 40)), "\n")
}
