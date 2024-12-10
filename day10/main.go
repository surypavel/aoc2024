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

type Trailhead struct {
	start  Pair
	lookup []Pair
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

type TopoMap map[Pair]int
type checker func(i int, t Trailhead, m TopoMap) Trailhead

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

func neighbours(coordinate Pair) []Pair {
	c1 := add(coordinate, Pair{X: 1, Y: 0})
	c2 := add(coordinate, Pair{X: -1, Y: 0})
	c3 := add(coordinate, Pair{X: 0, Y: 1})
	c4 := add(coordinate, Pair{X: 0, Y: -1})
	cs := []Pair{c1, c2, c3, c4}
	return cs
}

func find_distinct_paths_by_height(i int, t Trailhead, m TopoMap) Trailhead {
	new_coodinates := make([]Pair, 0)
	seen := make(map[Pair]bool)

	for _, coordinate := range t.lookup {
		cs := neighbours(coordinate)
		for _, c := range cs {
			if m[c] == i && !seen[c] {
				seen[c] = true
				new_coodinates = append(new_coodinates, c)
			}
		}
	}

	return Trailhead{start: t.start, lookup: new_coodinates}
}

func find_paths_by_height(i int, t Trailhead, m TopoMap) Trailhead {
	new_coodinates := make([]Pair, 0)

	for _, coordinate := range t.lookup {
		cs := neighbours(coordinate)
		for _, c := range cs {
			if m[c] == i {
				new_coodinates = append(new_coodinates, c)
			}
		}
	}

	return Trailhead{start: t.start, lookup: new_coodinates}
}

func sum_by(ts []Trailhead, m TopoMap, c checker) int {
	sum := 0
	for _, trailhead := range ts {
		for i := 1; i <= 9; i++ {
			trailhead = c(i, trailhead, m)
		}
		sum += len(trailhead.lookup)
	}

	return sum
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	topo_map := make(map[Pair]int)
	trailheads := make([]Trailhead, 0)

	for y, row := range lines {
		for x, item := range row {
			height := to_int(string(item))
			coordinate := Pair{X: x, Y: y}
			topo_map[coordinate] = height
			if height == 0 {
				trailheads = append(trailheads, Trailhead{start: coordinate, lookup: []Pair{coordinate}})
			}
		}
	}

	fmt.Print("part 1 - ", sum_by(trailheads, topo_map, find_distinct_paths_by_height), "\n")
	fmt.Print("part 2 - ", sum_by(trailheads, topo_map, find_paths_by_height), "\n")
}
