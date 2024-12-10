package main

import (
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	ID  int
	Len int
}

type Partition struct {
	Space int
	Items []Item
}

func remove(slice []Item, s int) []Item {
	return append(slice[:s], slice[s+1:]...)
}

func item_sum(items []Item) int {
	sum := 0
	for _, item := range items {
		sum += item.Len
	}
	return sum
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

// Attempt for single-loop solution
// Ended up pretty ugly and hard to re-use for part 2
func part1() int {
	input, err := os.ReadFile("input.txt")
	check(err)

	total := len(input)

	sum := 0
	for block, l, r := 0, 0, total-1; l <= total-1 && r > 0; {
		num_l := to_int(string(input[l]))
		num_r := to_int(string(input[r]))

		if num_l == 0 {
			l++
			continue
		}

		if l%2 == 0 {
			input[l] -= 1
			sum += (l / 2) * block
			block++
			continue
		}

		if l%2 == 1 {
			if num_r == 0 {
				r -= 2
			} else if r%2 == 1 {
				r--
			} else {
				input[r] -= 1
				input[l] -= 1
				sum += (r / 2) * block
				block++
			}
		}
	}
	return sum
}

func part2() int {
	input, err := os.ReadFile("input.txt")
	check(err)

	arr := make([]Partition, 0)
	for i, num := range input {
		items := make([]Item, 0)
		if i%2 == 0 {
			items = append(items, Item{ID: i / 2, Len: int(num) - int('0')})
		}
		arr = append(arr, Partition{Space: int(num) - int('0'), Items: items})
	}

	for i := len(arr) - 1; i > 0; i-- {
		if len(arr[i].Items) > 0 {
			for j := 1; j < i; j++ {
				space_needed := arr[i].Items[0].Len
				space_available := arr[j].Space - item_sum(arr[j].Items)
				if space_available >= space_needed {
					arr[j].Items = append(arr[j].Items, arr[i].Items[0])
					arr[i].Items = remove(arr[i].Items, 0)
					break
				}
			}
		}
	}

	index := 0
	sum := 0

	for _, partition := range arr {
		for _, item := range partition.Items {
			sum += item.ID * item.Len * (2*index + item.Len - 1) / 2
			index += item.Len
		}
		index += partition.Space - item_sum(partition.Items)
	}

	return sum
}

func main() {
	fmt.Print("part 1 - ", part1(), "\n")
	fmt.Print("part 2 - ", part2(), "\n")
}
