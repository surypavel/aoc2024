package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

type Connection struct {
	a string
	b string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func is_connected_to_all(nodes []string, new_node string, connections map[Connection]bool) bool {
	is_connected_to_all := true
	for _, node := range nodes {
		if !connections[Connection{a: node, b: new_node}] {
			is_connected_to_all = false
			break
		}
	}
	return is_connected_to_all
}

func get_connected_nodes(nodes []string, options []string, connections map[Connection]bool) []string {
	sibling_options := make([]string, 0)
	for _, new_sibling := range options {

		if is_connected_to_all(nodes, new_sibling, connections) {
			sibling_options = append(sibling_options, new_sibling)
		}
	}
	return sibling_options
}

func find_continuations(group []string, options []string, connections map[Connection]bool) [][]string {
	if len(options) == 0 {
		return [][]string{group}
	}

	new_groups := make([][]string, 0)

	for i, new_sibling := range options {
		if is_connected_to_all(group, new_sibling, connections) {
			concat := make([]string, 0)
			concat = append(concat, group...)
			concat = append(concat, new_sibling)
			new_groups = append(new_groups, find_continuations(concat, options[(i+1):], connections)...)
		}
	}

	if len(new_groups) == 0 {
		return [][]string{group}
	}

	return new_groups
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	// Prepare some useful structures
	m := make(map[string][]string)
	c := make(map[Connection]bool)
	for _, line := range lines {
		items := strings.Split(line, "-")
		if items[0] < items[1] {
			m[items[0]] = append(m[items[0]], items[1])
		} else {
			m[items[1]] = append(m[items[1]], items[0])
		}
		c[Connection{a: items[0], b: items[1]}] = true
		c[Connection{a: items[1], b: items[0]}] = true
	}

	nodes := slices.Collect(maps.Keys(m))
	slices.Sort(nodes)

	// Calculate part 1
	sum := 0
	for _, node := range nodes {
		connected_nodes := m[node]
		slices.Sort(connected_nodes)
		for index1, sibling1 := range connected_nodes {
			for _, next := range get_connected_nodes([]string{node, sibling1}, connected_nodes[(index1+1):], c) {
				if strings.HasPrefix(node, "t") || strings.HasPrefix(sibling1, "t") || strings.HasPrefix(next, "t") {
					sum++
				}
			}
		}
	}

	fmt.Print("part 1 - ", sum, "\n")

	// Part 2 - calculate biggest subgraphs containing items
	max_continuation := []string{}
	for _, node := range nodes {
		connected_nodes := m[node]
		slices.Sort(connected_nodes)
		for _, continuation := range find_continuations([]string{node}, connected_nodes, c) {
			if len(continuation) > len(max_continuation) {
				max_continuation = continuation
			}
		}
	}

	fmt.Print("part 2 - ", strings.Join(max_continuation, ","), "\n")
}
