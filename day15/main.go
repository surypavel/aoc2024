package main

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	X int
	Y int
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func mul(a Pair, n int) Pair {
	return Pair{X: a.X * n, Y: a.Y * n}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func to_direction(r rune) Pair {
	if r == '>' {
		return Pair{X: 1, Y: 0}
	}
	if r == '<' {
		return Pair{X: -1, Y: 0}
	}
	if r == '^' {
		return Pair{X: 0, Y: -1}
	}
	if r == 'v' {
		return Pair{X: 0, Y: 1}
	}
	panic("Oh wee :(")
}

func parse_map(input string) (map[Pair]rune, Pair) {
	items := make(map[Pair]rune)
	current := Pair{0, 0}

	robots_array := strings.Split(input, "\n")
	for y, robot_line := range robots_array {
		for x, robot_item := range robot_line {
			p := Pair{X: x, Y: y}
			if robot_item == '#' || robot_item == 'O' {
				items[p] = robot_item
			} else if robot_item == '@' {
				current = p
			}
		}
	}

	return items, current
}

func parse_map_2(input string) (map[Pair]rune, Pair) {
	items := make(map[Pair]rune)
	current := Pair{0, 0}

	robots_array := strings.Split(input, "\n")
	for y, robot_line := range robots_array {
		for x, robot_item := range robot_line {
			fst := Pair{X: 2 * x, Y: y}
			snd := Pair{X: 2*x + 1, Y: y}
			if robot_item == '#' {
				items[fst] = robot_item
				items[snd] = robot_item
			} else if robot_item == 'O' {
				items[fst] = '['
				items[snd] = ']'
			} else if robot_item == '@' {
				current = fst
			}
		}
	}

	return items, current
}

func find_path_length(m *map[Pair]rune, c Pair, d Pair) int {
	for i := 1; ; i++ {
		obstacle := (*m)[add(c, mul(d, i))]

		if obstacle == '#' {
			return 0
		}

		if obstacle == rune(0) {
			return i
		}
	}
}

func mov_x(m *map[Pair]rune, c Pair, d Pair) (Pair, []Pair) {
	blockers := make([]Pair, 0)
	for i := 1; ; i++ {
		coordinate := add(c, mul(d, i))
		obstacle := (*m)[add(c, mul(d, i))]

		if obstacle == '#' {
			return c, []Pair{}
		}

		if obstacle == '[' || obstacle == ']' {
			blockers = append(blockers, coordinate)
		}

		if obstacle == rune(0) {
			return add(c, d), blockers
		}
	}
}

func move(m *map[Pair]rune, c Pair, d Pair) Pair {
	lookahead := find_path_length(m, c, d)

	// Wall blocking, do not move
	if lookahead == 0 {
		return c
	}

	if lookahead >= 2 {
		(*m)[add(c, d)] = rune(0)
		(*m)[add(c, mul(d, lookahead))] = 'O'
	}
	return add(c, d)
}

func eval_map(m *map[Pair]rune, c rune) int {
	sum := 0
	for p, value := range *m {
		if value == c {
			sum += p.X + 100*p.Y
		}
	}
	return sum
}

func part1() int {
	input, err := os.ReadFile("input.txt")
	check(err)

	input_parts := strings.Split(string(input), "\n\n")

	m, c := parse_map(input_parts[0])

	for _, b := range input_parts[1] {
		// Some unprintable character there
		if b == 10 {
			continue
		}

		direction := to_direction(b)
		c = move(&m, c, direction)
	}

	return eval_map(&m, 'O')
}

func part2() int {
	input, err := os.ReadFile("input.txt")
	check(err)

	input_parts := strings.Split(string(input), "\n\n")

	m, c := parse_map_2(input_parts[0])

	for _, b := range input_parts[1] {
		// Some unprintable character there
		if b == 10 {
			continue
		}

		direction := to_direction(b)
		c = move(&m, c, direction)
	}

	return eval_map(&m, '[')
}

func main() {
	fmt.Print("part 1 - ", part1(), "\n")
	// fmt.Print("part 2 - ", part2(), "\n")
}
