package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	items := strings.Split(string(dat), "\n\n")
	rules := strings.Split(items[0], "\n")

	rules_map := make(map[string][]string)
	for _, rule := range rules {
		rule_array := strings.Split(rule, "|")

		if (rules_map[rule_array[0]]) == nil {
			rules_map[rule_array[0]] = make([]string, 0)
		}

		rules_map[rule_array[0]] = append(rules_map[rule_array[0]], rule_array[1])
	}

	actions := strings.Split(items[1], "\n")
	sum_part_1 := 0
	sum_part_2 := 0

	for _, action := range actions {
		action_array := strings.Split(action, ",")
		action_part_1 := calc_part_1(action_array, rules_map)

		if action_part_1 > 0 {
			sum_part_1 += action_part_1
		} else {
			sum_part_2 += calc_part_2(action_array, rules)
		}
	}

	fmt.Print("part 1 - ", sum_part_1, "\n")
	fmt.Print("part 2 - ", sum_part_2, "\n")
}
