package main

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
)

type item interface {
	eval(map[string]item) bool
	symbolic_eval(map[string]item) []string
}

type constant struct {
	name  string
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

func (e equation) symbolic_eval(m map[string]item) []string {
	new_deps := make([]string, 0)
	new_deps = append(new_deps, e.op)
	sym_eval_a := m[e.a].symbolic_eval(m)
	sym_eval_b := m[e.b].symbolic_eval(m)
	if len(sym_eval_a) > len(sym_eval_b) || (len(sym_eval_a) == 1 && sym_eval_a[0][0] == 'x') {
		new_deps = append(new_deps, sym_eval_a...)
		new_deps = append(new_deps, sym_eval_b...)
	} else {
		new_deps = append(new_deps, sym_eval_b...)
		new_deps = append(new_deps, sym_eval_a...)
	}
	return new_deps
}

func (r constant) symbolic_eval(m map[string]item) []string {
	return []string{r.name}
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

func pad(value int) string {
	return fmt.Sprintf("%02d", value)
}

func get_correct_symbolic_eval(n int) []string {
	result := make([]string, 0)
	result = append(result, "XOR")

	if n == 0 {
		result = append(result, "x00", "y00")
	}

	if n > 0 {
		for i := 1; i < n; i++ {
			result = append(result, "OR", "AND")
		}

		for i := 0; i < n; i++ {
			result = append(result, "AND", "x"+pad(i), "y"+pad(i), "XOR", "x"+pad(i+1), "y"+pad(i+1))
		}
	}

	return result
}

func find_item_by_symbolic_eval(assignments map[string]item, all_bits []string, search []string) string {
	for _, bit := range all_bits {
		if reflect.DeepEqual(assignments[bit].symbolic_eval(assignments), search) {
			return bit
		}
	}
	panic("Not found.")
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	// Parse input

	assignments := make(map[string]item)
	result_bits := make([]string, 0)
	all_bits := make([]string, 0)
	blocks := strings.Split(string(input), "\n\n")
	for _, line := range strings.Split(blocks[0], "\n") {
		items := strings.Split(line, ": ")
		assignments[items[0]] = constant{name: items[0], value: items[1] == "1"}
	}

	for _, line := range strings.Split(blocks[1], "\n") {
		items := strings.Split(line, " ")
		assignments[items[4]] = equation{a: items[0], b: items[2], op: items[1]}

		all_bits = append(all_bits, items[4])
		if strings.HasPrefix(items[4], "z") {
			result_bits = append(result_bits, items[4])
		}
	}

	slices.Sort(result_bits)

	// Part 1

	total := 0
	for i, bit := range result_bits {
		result_bit := assignments[bit].eval(assignments)
		if result_bit {
			total += int_pow(2, i)
		}
	}

	fmt.Print("part 1 - ", total, "\n")

	// Part 2
	pairs := [][]string{}

	for i, bit := range result_bits {
		if i == len(result_bits)-1 {
			// Let's hope that the last bit is calculated correctly
			// Since `symbolic_eval` does not work for last bit
			continue
		}

		symbolic_eval := assignments[bit].symbolic_eval(assignments)
		correct_symbolic_eval := get_correct_symbolic_eval(i)

		if !reflect.DeepEqual(symbolic_eval, correct_symbolic_eval) {
			// These are two cases i had in my input.
			// I'm not sure what are all the general cases that can be wrong.
			if len(symbolic_eval) == len(correct_symbolic_eval) {
				for i := range symbolic_eval {
					if symbolic_eval[i] != correct_symbolic_eval[i] {
						part1 := find_item_by_symbolic_eval(assignments, all_bits, symbolic_eval[i:])
						part2 := find_item_by_symbolic_eval(assignments, all_bits, correct_symbolic_eval[i:])

						assignments[part1], assignments[part2] = assignments[part2], assignments[part1]
						pairs = append(pairs, []string{part1, part2})
					}
				}
			} else {
				part1 := bit
				var part2 string
				for _, swap_bit := range all_bits {
					if reflect.DeepEqual(assignments[swap_bit].symbolic_eval(assignments), correct_symbolic_eval) {
						part2 = swap_bit
					}
				}

				assignments[part1], assignments[part2] = assignments[part2], assignments[part1]
				pairs = append(pairs, []string{part1, part2})
			}
		}

	}

	pairs_slice := make([]string, 0)
	for _, pair := range pairs {
		pairs_slice = append(pairs_slice, pair[0], pair[1])
	}
	slices.Sort(pairs_slice)

	fmt.Print("part 2 - ", strings.Join(pairs_slice, ","), "\n")
}
