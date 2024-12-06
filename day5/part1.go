package main

import (
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calc_part_1(action_array []string, rules_map map[string][]string) int {
	for index, item := range action_array {
		disallowed_before := rules_map[item]
		for previous := 0; previous < index; previous++ {
			for _, disallowed := range disallowed_before {
				if disallowed == action_array[previous] {
					return 0
				}
			}
		}
	}
	num, err := strconv.Atoi(action_array[(len(action_array)-1)/2])
	check(err)

	return num
}
