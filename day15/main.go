package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func append_if_missing[T comparable](a []T, b T) []T {
	if !slices.Contains(a, b) {
		return append(a, b)
	}
	return a

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

// This could actually also work as mov_y if there was normal vector used
func mov_x(m *map[Pair]rune, c Pair, d Pair) (Pair, []Pair) {
	boxes := make([]Pair, 0)
	for i := 1; ; i++ {
		coordinate := add(c, mul(d, i))
		obstacle := (*m)[coordinate]

		if obstacle == '#' {
			return c, []Pair{}
		}

		if obstacle == '[' || obstacle == ']' || obstacle == 'O' {
			boxes = append(boxes, coordinate)
		}

		if obstacle == rune(0) {
			return add(c, d), boxes
		}
	}
}

func mov_y(m *map[Pair]rune, c Pair, d Pair) (Pair, []Pair) {
	boxes := make([]Pair, 0)

	pushed_offsets := make(map[int]bool)
	pushed_offsets[0] = true

	for i := 1; ; i++ {
		if len(pushed_offsets) == 0 {
			return add(c, d), boxes
		}

		pushed_offsets_copy := make([]int, 0)

		for k := range pushed_offsets {
			pushed_offsets_copy = append(pushed_offsets_copy, k)
		}

		for _, j := range pushed_offsets_copy {
			coordinate := add(add(c, mul(d, i)), Pair{X: j, Y: 0})
			obstacle := (*m)[coordinate]

			if obstacle == '#' {
				return c, []Pair{}
			}

			if obstacle == 'O' {
				boxes = append_if_missing(boxes, coordinate)
			}

			if obstacle == '[' {
				boxes = append_if_missing(boxes, coordinate)
				boxes = append_if_missing(boxes, add(coordinate, Pair{X: 1, Y: 0}))
				pushed_offsets[j+1] = true
			}

			if obstacle == ']' {
				boxes = append_if_missing(boxes, coordinate)
				boxes = append_if_missing(boxes, add(coordinate, Pair{X: -1, Y: 0}))
				pushed_offsets[j-1] = true
			}

			if obstacle == rune(0) {
				delete(pushed_offsets, j)
			}
		}
	}
}

func move_and_update_map_2(m *map[Pair]rune, c Pair, d Pair) Pair {
	var (
		new_c Pair
		boxes []Pair
	)

	if d.X != 0 {
		new_c, boxes = mov_x(m, c, d)
	} else {
		new_c, boxes = mov_y(m, c, d)
	}

	for _, box := range slices.Backward(boxes) {
		(*m)[add(box, d)] = (*m)[box]
		(*m)[box] = rune(0)
	}

	return new_c
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
		c = move_and_update_map_2(&m, c, direction)
	}

	return eval_map(&m, 'O')
}

func part2() int {
	input, err := os.ReadFile("input.txt")
	check(err)

	input_parts := strings.Split(string(input), "\n\n")

	m, c := parse_map_2(input_parts[0])

	for _, b := range input_parts[1] {
		// print_map(&m, c)
		// print(string(b))

		// Some unprintable character there
		if b == 10 {
			continue
		}

		direction := to_direction(b)
		c = move_and_update_map_2(&m, c, direction)
	}

	return eval_map(&m, '[')
}

func print_map(m *map[Pair]rune, c Pair) {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	for y := 0; y < 12; y++ {
		for x := 0; x < 24; x++ {

			p := Pair{X: x, Y: y}

			if c == p {
				fmt.Print("@")
			} else if (*m)[p] != rune(0) {
				fmt.Print(string((*m)[p]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fmt.Print("part 1 - ", part1(), "\n")
	fmt.Print("part 2 - ", part2(), "\n")
}
