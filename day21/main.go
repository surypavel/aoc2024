package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
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

func move_on_keypad(line string, keypad map[rune]Pair, start Pair) []Move {
	current := start
	moves := []Move{}
	for _, item := range line {
		x := current.X - keypad[item].X
		y := current.Y - keypad[item].Y
		diff := Pair{X: -x, Y: -y}

		x_first := add(current, Pair{X: -x, Y: 0})
		y_first := add(current, Pair{X: 0, Y: -y})

		moves = append(moves, Move{Diff: diff, CanX: keypad['!'] != x_first, CanY: keypad['!'] != y_first})
		current = add(current, diff)
	}
	return moves
}

func to_directions(directions []Move) string {
	results := []byte{}
	for _, direction := range directions {
		results = append(results, combine(direction)...)
	}
	return string(results)
}

func combine(p Move) string {
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
		return s2
	}

	if !p.CanY {
		return s1
	}

	if s1 != s2 {
		if p.Diff.X < 0 {
			return s1
		}

		if p.Diff.Y < 0 {
			return s2
		}

		return s2
	}
	return s1
}

func recursive_move_on_keypad(line string, keypad map[rune]Pair, start Pair, n int) string {
	if n == 0 {
		return line
	}

	move_next := to_directions(move_on_keypad(line, keypad, Pair{X: 3, Y: 1}))
	return recursive_move_on_keypad(move_next, keypad, start, n-1)
}

func calc_complexity(input string, robot_count int) int {
	keypad := map[rune]Pair{
		'7': {X: 1, Y: 1},
		'8': {X: 2, Y: 1},
		'9': {X: 3, Y: 1},
		'4': {X: 1, Y: 2},
		'5': {X: 2, Y: 2},
		'6': {X: 3, Y: 2},
		'1': {X: 1, Y: 3},
		'2': {X: 2, Y: 3},
		'3': {X: 3, Y: 3},
		'!': {X: 1, Y: 4},
		'0': {X: 2, Y: 4},
		'A': {X: 3, Y: 4},
	}

	keypad_2 := map[rune]Pair{
		'!': {X: 1, Y: 1},
		'^': {X: 2, Y: 1},
		'A': {X: 3, Y: 1},
		'<': {X: 1, Y: 2},
		'v': {X: 2, Y: 2},
		'>': {X: 3, Y: 2},
	}

	eval := 0
	for _, line := range strings.Split(input, "\n") {
		moves_1 := to_directions(move_on_keypad(line, keypad, Pair{X: 3, Y: 4}))

		robot_move := recursive_move_on_keypad(moves_1, keypad_2, Pair{X: 3, Y: 1}, robot_count)

		eval += len(robot_move) * to_int(line[0:len(line)-1])
	}
	return eval
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	fmt.Print("part 1 - ", calc_complexity(string(input), 2), "\n")
	fmt.Print("part 2 - ", calc_complexity(string(input), 19), "\n")
}
