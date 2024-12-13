package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Machine struct {
	A     Pair
	B     Pair
	Prize Pair
}

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

func parse_input(input []byte) []Machine {
	machines := make([]Machine, 0)
	sections_array := strings.Split(string(input), "\n\n")

	for _, section_string := range sections_array {
		lines_array := strings.Split(section_string, "\n")
		parse := regexp.MustCompile(`(\d+)`)

		pairs := make([]Pair, 0)
		for _, line_string := range lines_array {
			match := parse.FindAllStringSubmatch(line_string, -1)
			x := to_int(match[0][0])
			y := to_int(match[1][0])
			pair := Pair{X: x, Y: y}
			pairs = append(pairs, pair)
		}

		machines = append(machines, Machine{A: pairs[0], B: pairs[1], Prize: pairs[2]})
	}

	return machines
}

func find_m(n int, primary_pair Pair, secondary_pair Pair, prize Pair) int {
	is_not_too_big := primary_pair.X*n <= prize.X && primary_pair.Y*n <= prize.Y
	is_divisible := (prize.X-primary_pair.X*n)%secondary_pair.X == 0 && (prize.Y-primary_pair.Y*n)%secondary_pair.Y == 0
	is_equal_part := (prize.Y-primary_pair.Y*n)/secondary_pair.Y == (prize.X-primary_pair.X*n)/secondary_pair.X

	if is_not_too_big && is_divisible && is_equal_part {
		return (prize.Y - primary_pair.Y*n) / secondary_pair.Y
	}

	return -1
}

func determine_n(primary_pair Pair, secondary_pair Pair, prize Pair) (int, int) {
	for n := 100; n >= 0; n-- {
		m := find_m(n, primary_pair, secondary_pair, prize)

		if m != -1 {
			return n, m
		}
	}

	// No solution
	return 0, 0
}

func calc_tokens(machine Machine) int {
	sumA := machine.A.X + machine.A.Y
	sumB := machine.B.X + machine.B.Y

	// Pushing A costs 3 tokens
	// That means, pushing B is more effective unless it has too little impact
	shouldPushB := sumB < 3*sumA

	if shouldPushB {
		b, a := determine_n(machine.B, machine.A, machine.Prize)
		return b + 3*a
	} else {
		a, b := determine_n(machine.A, machine.B, machine.Prize)
		return b + 3*a
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	machines := parse_input(input)

	sum_1 := 0
	for _, machine := range machines {
		sum_1 += calc_tokens(machine)
	}

	fmt.Print("part 1 - ", sum_1, "\n")
}
