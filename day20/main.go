package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

type Pair struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find_neighbours(coordinate Pair) []Pair {
	c1 := add(coordinate, Pair{X: 1, Y: 0})
	c2 := add(coordinate, Pair{X: -1, Y: 0})
	c3 := add(coordinate, Pair{X: 0, Y: 1})
	c4 := add(coordinate, Pair{X: 0, Y: -1})
	cs := []Pair{c1, c2, c3, c4}
	return cs
}

func find_cheat_neighbours(coordinate Pair, max_len int) []Pair {
	m := make(map[Pair]bool)

	// Different lengths of cheats
	for l := 0; l <= max_len; l++ {
		for mov_x := -l; mov_x <= l; mov_x++ {
			p1 := add(coordinate, Pair{X: mov_x, Y: l - abs(mov_x)})
			p2 := add(coordinate, Pair{X: mov_x, Y: -(l - abs(mov_x))})
			m[p1] = true
			m[p2] = true
		}
	}

	return slices.Collect(maps.Keys(m))
}

func find_eval(blocked map[Pair]bool, starting_vector Pair) map[Pair]int {
	eval := map[Pair]int{}

	// This is 1 just so that it's not null value. The 1 needs to be subtracted later.
	eval[starting_vector] = 1
	to_search := []Pair{starting_vector}

	for len(to_search) > 0 {
		new_to_search := make([]Pair, 0)
		for _, v := range to_search {
			curr_eval := eval[v]
			neighbours := find_neighbours(v)
			for _, neighbour := range neighbours {
				if !blocked[neighbour] {
					new_eval := curr_eval + 1
					neighbour_eval := eval[neighbour]

					if neighbour_eval == 0 || new_eval < neighbour_eval {
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

func count_best_cheats(eval map[Pair]int, end Pair) int {
	count := 0
	curr_pos := end
	for eval[curr_pos] > 0 {
		// Find best cheats
		cheat_neighbours := find_cheat_neighbours(curr_pos, 2)
		for _, neighbour := range cheat_neighbours {
			time_saved := eval[curr_pos] - eval[neighbour] - 2
			if time_saved >= 100 && eval[neighbour] != 0 {
				count++
			}
		}

		// Backtrack previous position
		neighbours := find_neighbours(curr_pos)
		for _, neighbour := range neighbours {
			if eval[neighbour] == eval[curr_pos]-1 {
				curr_pos = neighbour
			}
		}
	}
	return count
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

	eval := find_eval(walls, start)

	// Subtract 1 (eval map is +1)
	fmt.Print("part 1 - ", count_best_cheats(eval, end), "\n")
}
