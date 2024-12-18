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

	fmt.Print(eval)

	return eval[end] - 1
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")
	obstacles := make(map[Pair]bool)

	// size := Pair{X: 6, Y: 6}
	// start := Pair{X: 0, Y: 0}
	// end := Pair{X: 6, Y: 6}
	// bytes_fallen := 12

	size := Pair{X: 70, Y: 70}
	start := Pair{X: 0, Y: 0}
	end := Pair{X: 70, Y: 70}
	bytes_fallen := 1024

	for i, line := range lines {
		if i < bytes_fallen {
			items := strings.Split(line, ",")
			obstacles[Pair{X: to_int(items[0]), Y: to_int(items[1])}] = true
		}
	}

	fmt.Print("part 1 - ", dijkstra(obstacles, size, start, end), "\n")
	fmt.Print("part 2 - ", 0, "\n")

}
