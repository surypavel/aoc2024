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

func expand(items map[string]int, number int) map[string]int {
	if number == 0 {
		return items
	}

	new_items := make(map[string]int)

	for item, qty := range items {
		if item == "0" {
			new_items["1"] += qty
		} else if len(item)%2 == 0 {
			mid := len(item) / 2
			new_items[item[:mid]] += qty
			new_items[trim_zeroes(item[mid:])] += qty
		} else {
			new_items[trim_zeroes(strconv.Itoa(to_int(item)*2024))] += qty
		}
	}

	return expand(new_items, number-1)
}

func trim_zeroes(input string) string {
	r, _ := regexp.Compile("^0*(0|[1-9][0-9]*)$")
	return r.FindStringSubmatch(input)[1]
}

func create_map(items []string) map[string]int {
	m := make(map[string]int)
	for _, item := range items {
		m[item] = 1
	}
	return m
}

func sum_map(m map[string]int) int {
	sum := 0
	for _, count := range m {
		sum += count
	}
	return sum
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(input), " ")

	fmt.Print("part 1 - ", sum_map(expand(create_map(items), 25)), "\n")
	fmt.Print("part 2 - ", sum_map(expand(create_map(items), 75)), "\n")
}
