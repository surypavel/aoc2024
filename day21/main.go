package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Key struct {
	C Pair
	T rune
}

type Move struct {
	Diff Pair
	CanY bool
	CanX bool
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

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func absmin(items []int) int {
// 	curr_min := 100000
// 	for _, item := range items {
// 		if abs(item) < abs(curr_min) {
// 			curr_min = item
// 		}
// 	}
// 	return curr_min
// }

func move_on_keypad(line string, keypad []Key, start Pair) []Move {
	current := start
	moves := []Move{}
	for _, item := range line {
		next := slices.IndexFunc(keypad, func(c Key) bool { return c.T == item })
		avoid := slices.IndexFunc(keypad, func(c Key) bool { return c.T == '!' })
		x := current.X - keypad[next].C.X
		y := current.Y - keypad[next].C.Y
		diff := Pair{X: -x, Y: -y}

		x_first := add(current, Pair{X: -x, Y: 0})
		y_first := add(current, Pair{X: 0, Y: -y})

		moves = append(moves, Move{Diff: diff, CanX: keypad[avoid].C != x_first, CanY: keypad[avoid].C != y_first})
		current = add(current, diff)
	}
	return moves
}

func to_directions(directions []Move) []string {
	results := []string{""}
	for _, direction := range directions {
		new_results := make([]string, 0)

		for _, d := range combine(direction) {
			for _, res := range results {
				new_results = append(new_results, res+d)
			}
		}
		results = new_results
	}
	return results
}

func combine(p Move) []string {
	x_sign := ""
	y_sign := ""

	if p.Diff.X > 0 {
		x_sign = ">"
	}
	if p.Diff.X < 0 {
		x_sign = "<"
	}
	if p.Diff.Y > 0 {
		y_sign = "v"
	}
	if p.Diff.Y < 0 {
		y_sign = "^"
	}

	s1 := strings.Repeat(x_sign, abs(p.Diff.X)) + strings.Repeat(y_sign, abs(p.Diff.Y)) + "A"
	s2 := strings.Repeat(y_sign, abs(p.Diff.Y)) + strings.Repeat(x_sign, abs(p.Diff.X)) + "A"

	if !p.CanX {
		return []string{s2}
	}

	if !p.CanY {
		return []string{s1}
	}

	if s1 != s2 {
		return []string{s1, s2}
	}
	return []string{s1}
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	keypad := []Key{
		{C: Pair{X: 1, Y: 1}, T: '7'}, {C: Pair{X: 2, Y: 1}, T: '8'}, {C: Pair{X: 3, Y: 1}, T: '9'},
		{C: Pair{X: 1, Y: 2}, T: '4'}, {C: Pair{X: 2, Y: 2}, T: '5'}, {C: Pair{X: 3, Y: 2}, T: '6'},
		{C: Pair{X: 1, Y: 3}, T: '1'}, {C: Pair{X: 2, Y: 3}, T: '2'}, {C: Pair{X: 3, Y: 3}, T: '3'},
		{C: Pair{X: 1, Y: 4}, T: '!'}, {C: Pair{X: 2, Y: 4}, T: '0'}, {C: Pair{X: 3, Y: 4}, T: 'A'},
	}

	keypad_2 := []Key{
		{C: Pair{X: 1, Y: 1}, T: '!'}, {C: Pair{X: 2, Y: 1}, T: '^'}, {C: Pair{X: 3, Y: 1}, T: 'A'},
		{C: Pair{X: 1, Y: 2}, T: '<'}, {C: Pair{X: 2, Y: 2}, T: 'v'}, {C: Pair{X: 3, Y: 2}, T: '>'},
	}

	eval := 0
	for _, line := range strings.Split(string(input), "\n") {
		fmt.Print(line, " -- ")

		moves_1 := to_directions(move_on_keypad(line, keypad, Pair{X: 3, Y: 4}))

		min_len := 10000000
		for _, move_1 := range moves_1 {
			moves_2 := to_directions(move_on_keypad(move_1, keypad_2, Pair{X: 3, Y: 1}))
			for _, move_2 := range moves_2 {
				moves_3 := to_directions(move_on_keypad(move_2, keypad_2, Pair{X: 3, Y: 1}))
				for _, move_3 := range moves_3 {
					if len(move_3) < min_len {
						min_len = len(move_3)
					}
				}
			}
		}

		fmt.Print(min_len)
		eval += min_len * to_int(line[0:len(line)-1])
		fmt.Print("\n")
	}

	fmt.Print("part 1 - ", eval, "\n")
}
