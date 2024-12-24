package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type item interface {
	eval(map[string]item) bool
}

type constant struct {
	value bool
}

type equation struct {
	a  string
	b  string
	op string
}

func (r constant) eval(m map[string]item) bool {
	return r.value
}

func (e equation) eval(m map[string]item) bool {
	if e.op == "AND" {
		return m[e.a].eval(m) && m[e.b].eval(m)
	} else if e.op == "OR" {
		return m[e.a].eval(m) || m[e.b].eval(m)
	} else if e.op == "XOR" {
		return m[e.a].eval(m) != m[e.b].eval(m)
	} else {
		panic("Something is wrong.")
	}
}

func int_pow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func format_bit(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	assignments := make(map[string]item)
	result_bits := make([]string, 0)
	blocks := strings.Split(string(input), "\n\n")
	for _, line := range strings.Split(blocks[0], "\n") {
		items := strings.Split(line, ": ")
		assignments[items[0]] = constant{value: items[1] == "1"}
	}

	for _, line := range strings.Split(blocks[1], "\n") {
		items := strings.Split(line, " ")
		assignments[items[4]] = equation{a: items[0], b: items[2], op: items[1]}

		if strings.HasPrefix(items[4], "z") {
			result_bits = append(result_bits, items[4])
		}
	}

	slices.Sort(result_bits)

	total := 0
	for i, bit := range result_bits {
		result_bit := assignments[bit].eval(assignments)
		fmt.Print(format_bit(result_bit))
		if result_bit {
			total += int_pow(2, i)
		}
	}

	fmt.Print("\n")

	fmt.Print("part 1 - ", total, "\n")
}
