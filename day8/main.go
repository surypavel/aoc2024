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

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func subtract(a Pair, b Pair) Pair {
	return Pair{X: a.X - b.X, Y: a.Y - b.Y}
}

func reduce(a Pair) Pair {
	return Pair{X: a.X / gcd(a.X, a.Y), Y: a.Y / gcd(a.X, a.Y)}
}

func add(a Pair, b Pair) Pair {
	return Pair{X: a.X + b.X, Y: a.Y + b.Y}
}

func out_of_bounds(a Pair, b Pair) bool {
	return a.X < 0 || a.Y < 0 || a.X >= b.X || a.Y >= b.Y
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_1(antennas map[rune][]Pair, bounding_box Pair) int {
	antinodes := make(map[Pair]bool, 0)
	for nodes := range maps.Values(antennas) {
		for _, node1 := range nodes {
			for _, node2 := range nodes {
				if node1 != node2 {
					difference := subtract(node1, node2)
					node_add := add(node1, difference)
					node_subtract := subtract(node2, difference)

					if !out_of_bounds(node_add, bounding_box) {
						antinodes[node_add] = true
					}
					if !out_of_bounds(node_subtract, bounding_box) {
						antinodes[node_subtract] = true
					}
				}
			}
		}
	}

	return len(slices.Collect(maps.Values(antinodes)))
}

func part_2(antennas map[rune][]Pair, bounding_box Pair) int {
	antinodes := make(map[Pair]bool, 0)
	for nodes := range maps.Values(antennas) {
		for _, node1 := range nodes {
			for _, node2 := range nodes {
				if node1 != node2 {
					difference := reduce(subtract(node1, node2))

					for point := node1; !out_of_bounds(point, bounding_box); point = add(point, difference) {
						antinodes[point] = true
					}

					for point := node1; !out_of_bounds(point, bounding_box); point = subtract(point, difference) {
						antinodes[point] = true
					}
				}
			}
		}
	}

	return len(slices.Collect(maps.Values(antinodes)))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	antennas := make(map[rune][]Pair, 0)
	items := strings.Split(string(dat), "\n")
	bounding_box := Pair{X: len(items[0]), Y: len(items)}
	for i, row := range items {
		for j, item := range row {
			if item != '.' {
				antennas[item] = append(antennas[item], Pair{X: j, Y: i})
			}
		}
	}

	fmt.Print("part 1 - ", part_1(antennas, bounding_box), "\n")
	fmt.Print("part 2 - ", part_2(antennas, bounding_box), "\n")
}
