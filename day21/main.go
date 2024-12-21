package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CacheEntry struct {
	n    int
	line string
}

type Pair struct {
	X int
	Y int
}

type Move struct {
	Diff Pair
	CanY bool
	CanX bool
}

var NUMERIC_KEYPAD = map[byte]Pair{
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

var DIRECTIONAL_KEYPAD = map[byte]Pair{
	'!': {X: 1, Y: 1},
	'^': {X: 2, Y: 1},
	'A': {X: 3, Y: 1},
	'<': {X: 1, Y: 2},
	'v': {X: 2, Y: 2},
	'>': {X: 3, Y: 2},
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

func sgn_x(p Pair) string {
	if p.X > 0 {
		return ">"
	} else {
		return "<"
	}
}

func sgn_y(p Pair) string {
	if p.Y > 0 {
		return "v"
	} else {
		return "^"
	}
}

// This function is pretty much magic but it seems to determine the best route on the keyboard
func combine(p Move) []byte {
	p1 := strings.Repeat(sgn_x(p.Diff), abs(p.Diff.X))
	p2 := strings.Repeat(sgn_y(p.Diff), abs(p.Diff.Y))
	s1 := []byte(p1 + p2 + "A")
	s2 := []byte(p2 + p1 + "A")

	if !p.CanX {
		return s2
	} else if !p.CanY {
		return s1
	} else if p.Diff.X < 0 {
		return s1
	} else {
		return s2
	}
}

func move_on_keypad(line []byte, keypad map[byte]Pair) [][]byte {
	current := keypad['A']
	moves := [][]byte{}
	for _, item := range line {
		x := current.X - keypad[item].X
		y := current.Y - keypad[item].Y
		diff := Pair{X: -x, Y: -y}

		x_first := add(current, Pair{X: -x, Y: 0})
		y_first := add(current, Pair{X: 0, Y: -y})

		moves = append(moves, combine(Move{
			Diff: diff,
			CanX: keypad['!'] != x_first,
			CanY: keypad['!'] != y_first,
		}))

		current = add(current, diff)
	}
	return moves
}

func cached_recursive_robot_move(line []byte, n int, m *map[CacheEntry]int) int {
	if n == 0 {
		return len(line)
	}

	cache_key := CacheEntry{n: n, line: string(line)}
	cached_entry := (*m)[cache_key]

	if cached_entry != 0 {
		return cached_entry
	}

	move_next := move_on_keypad(line, DIRECTIONAL_KEYPAD)

	sum := 0
	for _, next := range move_next {
		sum += cached_recursive_robot_move(next, n-1, m)
	}
	(*m)[cache_key] = sum
	return sum
}

func calc_complexity(input string, robot_count int) int {
	eval := 0
	m := make(map[CacheEntry]int)

	for _, line := range strings.Split(input, "\n") {
		moves := move_on_keypad([]byte(line), NUMERIC_KEYPAD)
		moves_sum := 0

		for _, move := range moves {
			moves_sum += cached_recursive_robot_move(move, robot_count, &m)
		}

		eval += moves_sum * to_int(line[0:len(line)-1])
	}

	return eval
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	fmt.Print("part 1 - ", calc_complexity(string(input), 2), "\n")
	fmt.Print("part 2 - ", calc_complexity(string(input), 25), "\n")
}
