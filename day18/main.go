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

func IntPow(n, m int) int {
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

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func out_of_bounds(a Pair, b Pair) bool {
	return a.X < 0 || a.Y < 0 || a.X > b.X || a.Y > b.Y
}

func neighbours(coordinate Pair) []Pair {
	c1 := add(coordinate, Pair{X: 1, Y: 0})
	c2 := add(coordinate, Pair{X: -1, Y: 0})
	c3 := add(coordinate, Pair{X: 0, Y: 1})
	c4 := add(coordinate, Pair{X: 0, Y: -1})
	cs := []Pair{c1, c2, c3, c4}
	return cs
}

func to_int(s string) int {
	num, err := strconv.Atoi(s)
	check(err)
	return num
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dijkstra(obstacles map[Pair]bool, size Pair, start Pair, end Pair) int {
	eval := make(map[Pair]int)
	to_check := []Pair{start}
	eval[start] = 1

	for len(to_check) > 0 {
		new_to_check := make([]Pair, 0)
		for _, coordinate := range to_check {
			for _, neighbour := range neighbours(coordinate) {
				if !obstacles[neighbour] && !out_of_bounds(neighbour, size) && eval[neighbour] == 0 {
					eval[neighbour] = eval[coordinate] + 1
					new_to_check = append(new_to_check, neighbour)
				}
			}
		}
		to_check = new_to_check
	}

	return eval[end] - 1
}

func to_map(lines []string, bytes_fallen int) map[Pair]bool {
	obstacles := make(map[Pair]bool)
	for _, line := range lines[0:bytes_fallen] {
		items := strings.Split(line, ",")
		obstacles[Pair{X: to_int(items[0]), Y: to_int(items[1])}] = true
	}
	return obstacles
}

func main() {
	// input, err := os.ReadFile("example.txt")
	// size_int := 6
	// buffer := 12

	input, err := os.ReadFile("input.txt")
	size_int := 70
	buffer := 1024

	size := Pair{X: size_int, Y: size_int}
	start := Pair{X: 0, Y: 0}
	end := Pair{X: size_int, Y: size_int}

	check(err)

	lines := strings.Split(string(input), "\n")

	fmt.Print("part 1 - ", dijkstra(to_map(lines, buffer), size, start, end), "\n")

	dijkstra(to_map(lines, buffer), size, start, end)

	unreachable_offsets := make([]int, 0)

	// 15 should be enough, would have to calculate approximate log2(len(lines)) to be exact
	for i, offset := 1, len(lines)/2; i < 15; i++ {
		route := dijkstra(to_map(lines, offset), size, start, end)
		diff := (len(lines) / IntPow(2, i+1)) + 1

		if route == -1 {
			unreachable_offsets = append(unreachable_offsets, offset)
			offset -= diff
		} else {
			offset += diff
		}
	}

	first_unreachable := slices.Min(unreachable_offsets) - 1

	fmt.Print("part 2 - ", lines[first_unreachable], "\n")
}
