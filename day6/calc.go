package main

import (
	"fmt"
	"strings"
)

func determine_viable_positions(obstacles map[string]bool, size Pair, starting_pos Pair, starting_dir Pair) []Pair {
	visited_positions_map := make(map[string]bool)
	visited_positions := make([]Pair, 0)
	current_dir := starting_dir
	current_pos := starting_pos

	// This is ultra naive way to detect loops - just look if it did not end in a number of steps
	// TODO: Do a proper loop detection
	for step := 1; step < 22000; step++ {
		if !visited_positions_map[hash(current_pos)] {
			visited_positions = append(visited_positions, current_pos)
		}

		visited_positions_map[hash(current_pos)] = true

		next := add(current_pos, current_dir)
		if obstacles[hash(next)] {
			current_dir = turn(current_dir)
		} else {
			current_pos = next
		}

		if out_of_bounds(current_pos, size) {
			return visited_positions
		}
	}

	return make([]Pair, 0)
}

func calc(dat []byte) {
	items := strings.Split(string(dat), "\n")
	obstacles := make(map[string]bool)
	current_pos := Pair{X: 0, Y: 0}
	current_dir := Pair{X: 0, Y: 0}

	size := Pair{X: len(items[0]) - 1, Y: len(items) - 1}

	for i, row := range items {
		for j, item := range row {
			pos := Pair{X: j, Y: i}

			if item == '#' {
				obstacles[hash(pos)] = true
			}
			if item == '^' {
				current_pos = pos
				current_dir = Pair{X: 0, Y: -1}
			}
		}
	}

	visited_positions := determine_viable_positions(obstacles, size, current_pos, current_dir)

	fmt.Print("part 1 - ", len(visited_positions), "\n")

	total_loops := 0

	for _, new_obstacle := range visited_positions {
		if new_obstacle == current_pos || out_of_bounds(new_obstacle, size) {
			continue
		}

		obstacles_copy := make(map[string]bool)
		for k, v := range obstacles {
			obstacles_copy[k] = v
		}

		obstacles_copy[hash(new_obstacle)] = true

		pos := determine_viable_positions(obstacles_copy, size, current_pos, current_dir)
		if len(pos) == 0 {
			total_loops++
		}
	}

	fmt.Print("part 2 - ", total_loops)
}
