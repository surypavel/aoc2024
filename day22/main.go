package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type History struct {
	a int
	b int
	c int
	d int
}

type IndexSequence struct {
	index   int
	history History
}

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

func mix(secret int, val int) int {
	return secret ^ val
}

func prune(secret int) int {
	return secret % 16777216
}

func evolve(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))
	return secret
}

func iterate_evolve(start int, n int) int {
	secret := start
	for i := 0; i < n; i++ {
		secret = evolve(secret)
	}
	return secret
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		sum += iterate_evolve(to_int(line), 2000)
	}

	fmt.Print("part 1 - ", sum, "\n")

	sequence_map := make(map[IndexSequence]int)
	for index, line := range lines {
		last := to_int(line)
		// Idk if the math is good here but it should be
		for i := 0; i < 1996; i++ {
			p1a := last
			p2a := evolve(p1a)
			p3a := evolve(p2a)
			p4a := evolve(p3a)
			p5a := evolve(p4a)

			p1 := p1a % 10
			p2 := p2a % 10
			p3 := p3a % 10
			p4 := p4a % 10
			p5 := p5a % 10

			sequence := IndexSequence{index: index, history: History{a: p2 - p1, b: p3 - p2, c: p4 - p3, d: p5 - p4}}
			if sequence_map[sequence] == 0 {
				sequence_map[sequence] = 100 + p5
			}

			last = p2a
		}
	}

	sum_by_sequence := make(map[History]int)
	for s, bananas := range sequence_map {
		sum_by_sequence[s.history] += bananas - 100
	}

	max_sequence := History{a: 0, b: 0, c: 0, d: 0}
	for s, bananas := range sum_by_sequence {
		if bananas > sum_by_sequence[max_sequence] {
			max_sequence = s
		}
	}

	fmt.Print("part 2 - ", sum_by_sequence[max_sequence], "\n")
}
