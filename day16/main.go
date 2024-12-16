package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Vector struct {
	Start     Pair
	Direction Pair
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_neighbouring_vectors(coordinate Pair) []Vector {
	c1 := Vector{Start: add(coordinate, Pair{X: 1, Y: 0}), Direction: Pair{X: 1, Y: 0}}
	c2 := Vector{Start: add(coordinate, Pair{X: -1, Y: 0}), Direction: Pair{X: -1, Y: 0}}
	c3 := Vector{Start: add(coordinate, Pair{X: 0, Y: 1}), Direction: Pair{X: 0, Y: 1}}
	c4 := Vector{Start: add(coordinate, Pair{X: 0, Y: -1}), Direction: Pair{X: 0, Y: -1}}
	cs := []Vector{c1, c2, c3, c4}
	return cs
}

func find_backtrack_vectors(coordinate Pair) []Vector {
	vs := make([]Vector, 0)
	dirs := []Pair{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}

	for _, d := range dirs {
		for _, a := range dirs {
			vs = append(vs, Vector{Start: add(coordinate, a), Direction: d})
		}
	}
	return vs
}

func find_eval(blocked map[Pair]bool, starting_vector Vector) map[Vector]int {
	eval := make(map[Vector]int)

	// This is 1 just so that it's not null value. The 1 needs to be subtracted later.
	eval[starting_vector] = 1

	to_search := make([]Vector, 0)
	to_search = append(to_search, starting_vector)

	for len(to_search) > 0 {
		new_to_search := make([]Vector, 0)
		for _, v := range to_search {
			curr_eval := eval[v]
			neighbours := find_neighbouring_vectors(v.Start)
			for _, neighbour := range neighbours {
				if !blocked[neighbour.Start] {
					new_eval := curr_eval + 1

					if neighbour.Direction != v.Direction {
						new_eval += 1000
					}

					if eval[neighbour] == 0 || new_eval < eval[neighbour] {
						eval[neighbour] = new_eval
						new_to_search = append(new_to_search, neighbour)
					}
				}
			}
		}
		to_search = new_to_search
	}

	return eval
}

func append_if_missing[T comparable](a []T, b T) []T {
	if !slices.Contains(a, b) {
		return append(a, b)
	}
	return a

}

func find_backtrack_length(eval map[Vector]int, finishing_vectors_argmin []Vector) int {
	to_search := finishing_vectors_argmin
	found_points := make(map[Pair]bool)

	for len(to_search) > 0 {
		new_to_search := make([]Vector, 0)
		for _, v := range to_search {
			neighbours := find_backtrack_vectors(v.Start)
			for _, neighbour := range neighbours {
				if eval[neighbour] != 0 {
					if eval[neighbour]+1 == eval[v] || eval[neighbour]+1001 == eval[v] {
						new_to_search = append_if_missing(new_to_search, neighbour)
						found_points[neighbour.Start] = true
					}
				}
			}
		}
		to_search = new_to_search
	}

	return len(found_points) + 1
}

func parse_map(input string) (map[Pair]bool, Pair, Pair) {
	walls := make(map[Pair]bool)
	start := Pair{0, 0}
	end := Pair{0, 0}

	a := strings.Split(input, "\n")
	for y, line := range a {
		for x, item := range line {
			p := Pair{X: x, Y: y}
			if item == '#' {
				walls[p] = true
			} else if item == 'S' {
				start = p
			} else if item == 'E' {
				end = p
			}
		}
	}

	return walls, start, end
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	walls, start, end := parse_map(string(input))

	starting_vector := Vector{Start: start, Direction: Pair{X: 1, Y: 0}}
	eval := find_eval(walls, starting_vector)

	finishing_vector_1 := Vector{Start: end, Direction: Pair{X: 1, Y: 0}}
	finishing_vector_2 := Vector{Start: end, Direction: Pair{X: 0, Y: -1}}

	// This works only if end field is accessible from both directions
	shortest_path := min(eval[finishing_vector_1], eval[finishing_vector_2])

	// Subtract 1 (eval map is +1)
	fmt.Print("part 1 - ", shortest_path-1, "\n")

	finishing_vectors := []Vector{finishing_vector_1, finishing_vector_2}
	finishing_vectors_argmin := make([]Vector, 0)
	for _, v := range finishing_vectors {
		if eval[v] == shortest_path {
			finishing_vectors_argmin = append(finishing_vectors_argmin, v)
		}
	}

	// This does not work if there are paths whose eval differ by 1 but do an incompatible turn.
	// However, this does not happen in the input data, because the found path is only 439 steps long, not 1000+ steps long.
	// This means that every time the eval differs by 1, it is actually a valid path, no matter the turn.
	backtrack_length := find_backtrack_length(eval, finishing_vectors_argmin)

	fmt.Print("part 2 - ", backtrack_length, "\n")
}
