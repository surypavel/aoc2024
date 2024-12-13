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

// This is working only if the resulting matrix is singular but it is always the case
func calc_solution(m Machine) (int, int) {
	// m.A.X + m.B.X = m.Price.X
	// m.A.Y + m.B.Y = m.Price.Y
	// (A.X B.X) (Price.X)
	// (A.Y B.Y) (Price.Y)
	// Calculate using cramer's rule

	det := (m.A.X*m.B.Y - m.A.Y*m.B.X)
	x1 := (m.Prize.X*m.B.Y - m.Prize.Y*m.B.X)
	x2 := (m.A.X*m.Prize.Y - m.A.Y*m.Prize.X)

	if x1%det == 0 && x2%det == 0 {
		return x1 / det, x2 / det
	}

	return 0, 0
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	machines := parse_input(input)

	sum_1 := 0
	for _, machine := range machines {
		x1, x2 := calc_solution(machine)
		sum_1 += 3*x1 + x2
	}

	fmt.Print("part 1 - ", sum_1, "\n")

	offset := 10000000000000
	sum_2 := 0
	for _, machine := range machines {
		x1, x2 := calc_solution(Machine{A: machine.A, B: machine.B, Prize: Pair{X: machine.Prize.X + offset, Y: machine.Prize.Y + offset}})
		sum_2 += 3*x1 + x2
	}

	fmt.Print("part 2 - ", sum_2, "\n")

}
