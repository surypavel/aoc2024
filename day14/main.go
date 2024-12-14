package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Robot struct {
	Pos Pair
	Vel Pair
}

func mod(x int, n int) int {
	return ((x % n) + n) % n
}

func move(r Robot, size Pair, steps int) Robot {
	return Robot{Vel: r.Vel, Pos: Pair{
		X: mod((r.Pos.X + steps*r.Vel.X), size.X),
		Y: mod((r.Pos.Y + steps*r.Vel.Y), size.Y),
	}}
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

func parse_input(input []byte) []Robot {
	robots := make([]Robot, 0)
	robots_array := strings.Split(string(input), "\n")

	for _, robot_string := range robots_array {
		parse := regexp.MustCompile(`(-?\d+)`)

		match := parse.FindAllStringSubmatch(robot_string, -1)
		robots = append(robots, Robot{
			Pos: Pair{X: to_int(match[0][0]), Y: to_int(match[1][0])},
			Vel: Pair{X: to_int(match[2][0]), Y: to_int(match[3][0])},
		})
	}

	return robots
}

func part1(robots []Robot, size Pair) int {
	partitions := make(map[string]int)

	for _, robot := range robots {
		moved_robot := move(robot, size, 100)
		if 2*moved_robot.Pos.X != size.X-1 && 2*moved_robot.Pos.Y != size.Y-1 {
			partition_x := 2 * moved_robot.Pos.X / size.X
			partition_y := 2 * moved_robot.Pos.Y / size.Y
			partitions[strconv.Itoa(partition_x)+"-"+strconv.Itoa(partition_y)] += 1
		}
	}

	product := 1

	for _, n := range partitions {
		product *= n
	}

	return product
}

func render(robots []Robot, size Pair) {
	m := make(map[Pair]bool)

	for _, robot := range robots {
		m[robot.Pos] = true
	}

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			has_robot := m[Pair{X: x, Y: y}]
			if has_robot {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	robots := parse_input(input)
	size := Pair{X: 101, Y: 103}

	fmt.Print("part 1 - ", part1(robots, size), "\n")

	// Very manual lookup, result 7520
	frequency := 103
	for i := 1; i < 100000; i += frequency {
		moved_robots := make([]Robot, 5)
		for _, robot := range robots {
			moved_robots = append(moved_robots, move(robot, size, i))
		}

		render(moved_robots, size)
		fmt.Print(i)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
