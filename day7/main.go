package main

import (
	"fmt"
	"os"
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

func to_ints(ts []string) []int {
	output := []int{}
	for _, t := range ts {
		num, err := strconv.Atoi(t)
		check(err)
		output = append(output, num)
	}
	return output
}

type checker func(int, []int) bool

func is_attainable_1(result int, numbers []int) bool {
	if numbers[0] > result {
		return false
	}

	if len(numbers) > 1 {
		first, second, rest := numbers[0], numbers[1], numbers[2:]
		return is_attainable_1(result, append([]int{first + second}, rest...)) || is_attainable_1(result, append([]int{first * second}, rest...))
	}
	return result == numbers[0]
}

func is_attainable_2(result int, numbers []int) bool {
	if numbers[0] > result {
		return false
	}

	if len(numbers) > 1 {
		first, second, rest := numbers[0], numbers[1], numbers[2:]
		return is_attainable_2(result, append([]int{to_int(strconv.Itoa(first) + strconv.Itoa(second))}, rest...)) || is_attainable_2(result, append([]int{first + second}, rest...)) || is_attainable_2(result, append([]int{first * second}, rest...))
	}
	return result == numbers[0]
}
func sum(items []string, is_attainable checker) int {
	sum := 0
	for _, item := range items {
		split := strings.Split(item, ": ")
		numbers := to_ints(strings.Split(split[1], " "))
		result := to_int(split[0])

		if is_attainable(result, numbers) {
			sum += result
		}
	}
	return sum
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(dat), "\n")

	fmt.Print("part 1 - ", sum(items, is_attainable_1), "\n")
	fmt.Print("part 2 - ", sum(items, is_attainable_2), "\n")
}
