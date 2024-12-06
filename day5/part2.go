package main

import (
	"slices"
	"strconv"
	"strings"
)

func find_first_index(action_array []string, rules []string) int {
	tails := make(map[string]bool)
	for _, rule := range rules {
		rule_array := strings.Split(rule, "|")
		if slices.Contains(action_array, rule_array[0]) && slices.Contains(action_array, rule_array[1]) {
			tails[rule_array[1]] = true
		}
	}

	first_index := -1

	for index, action := range action_array {
		if !tails[action] {
			first_index = index
			break
		}
	}

	return first_index
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func calc_part_2(action_array []string, rules []string) int {
	action_array_length := len(action_array)
	result := ""

	for i := 0; i < (action_array_length+1)/2; i++ {
		first_index := find_first_index(action_array, rules)
		result = action_array[first_index]
		action_array = remove(action_array, first_index)
	}

	num, err := strconv.Atoi(result)
	check(err)

	return num
}
