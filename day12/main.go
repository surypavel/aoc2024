package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	X int
	Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func find_neighbours(coordinate Pair) []Pair {
	c1 := add(coordinate, Pair{X: 1, Y: 0})
	c2 := add(coordinate, Pair{X: -1, Y: 0})
	c3 := add(coordinate, Pair{X: 0, Y: 1})
	c4 := add(coordinate, Pair{X: 0, Y: -1})
	cs := []Pair{c1, c2, c3, c4}
	return cs
}

func is_out_of_bounds(a Pair, b Pair) bool {
	return a.X < 0 || a.Y < 0 || a.X > b.X || a.Y > b.Y
}

func calc_perimeter(plot map[Pair]bool) int {
	perimeter := 0
	for item := range plot {
		neighbours := find_neighbours(item)
		for _, neighbour := range neighbours {
			if !plot[neighbour] {
				perimeter++
			}
		}
	}
	return perimeter
}

func calc_perimeter_with_discount(plot map[Pair]bool) int {
	p := make(map[string][]int)

	for item := range plot {
		neighbours := find_neighbours(item)
		for direction, neighbour := range neighbours {
			if !plot[neighbour] {
				if direction == 0 {
					k := "R" + string(item.X)
					p[k] = append(p[k], item.Y)
				}

				if direction == 1 {
					k := "L" + string(item.X)
					p[k] = append(p[k], item.Y)
				}

				if direction == 2 {
					k := "T" + string(item.Y)
					p[k] = append(p[k], item.X)
				}

				if direction == 3 {
					k := "B" + string(item.Y)
					p[k] = append(p[k], item.X)
				}
			}
		}
	}

	perimeter := 0

	for _, value := range p {
		sort.Ints(value)
		last_edge := -2
		for _, edge := range value {
			if last_edge != edge-1 {
				perimeter += 1
			}
			last_edge = edge
		}
	}

	return perimeter
}

func find_plot(m map[Pair]rune, size Pair, starting_point Pair) map[Pair]bool {
	visited := make(map[Pair]bool)
	plot_kind := m[starting_point]

	for to_visit := []Pair{starting_point}; len(to_visit) > 0; {
		new_to_visit := make([]Pair, 0)
		new_to_visit_added := make(map[Pair]bool)

		for _, coordinate := range to_visit {
			visited[coordinate] = true
			neighbours := find_neighbours(coordinate)
			for _, neighbour := range neighbours {
				if m[neighbour] == plot_kind && !visited[neighbour] && !is_out_of_bounds(neighbour, size) && !new_to_visit_added[neighbour] {
					new_to_visit = append(new_to_visit, neighbour)
					new_to_visit_added[neighbour] = true
				}
			}
		}

		to_visit = new_to_visit
	}
	return visited
}

type perimeter_fn func(plot map[Pair]bool) int

func sum_map(m map[Pair]rune, size Pair, calc perimeter_fn) int {
	visited := make(map[Pair]bool)
	price := 0

	for y := 0; y <= size.Y; y++ {
		for x := 0; x <= size.X; x++ {
			starting_point := Pair{X: x, Y: y}
			if !visited[starting_point] {
				plot := find_plot(m, size, starting_point)

				// Increase price
				// fmt.Print(len(plot), calc(plot), starting_point, string(m[starting_point]), "\n")
				price += len(plot) * calc(plot)

				// Mark the whole plot as visited
				for v := range plot {
					visited[v] = true
				}
			}
		}
	}

	return price
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	// Make a map for faster access
	m := make(map[Pair]rune)
	size := Pair{X: len(lines[0]) - 1, Y: len(lines) - 1}

	for y, row := range lines {
		for x, item := range row {
			coordinate := Pair{X: x, Y: y}
			m[coordinate] = item
		}
	}

	fmt.Print("part 1 - ", sum_map(m, size, calc_perimeter), "\n")
	fmt.Print("part 2 - ", sum_map(m, size, calc_perimeter_with_discount), "\n")
}
