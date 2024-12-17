package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

type Register struct {
	A int
	B int
	C int
}

func to_int(s string) int {
	num, err := strconv.Atoi(s)
	check(err)
	return num
}

func to_ints(ts []string) []int {
	output := []int{}
	for _, t := range ts {
		num, err := strconv.Atoi(t)
		check(err)
		output = append(output, num)
	}
	return output
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func combo(operand int, register Register) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return register.A
	case 5:
		return register.B
	case 6:
		return register.C
	default:
		panic("Incorrect operand.")
	}
}

func do(instruction int, operand int, register *Register) (out int, jump int) {
	switch instruction {
	case 0:
		adv := register.A / (IntPow(2, (combo(operand, *register))))
		register.A = adv
		return -1, -1
	case 1:
		bxl := register.B ^ operand
		register.B = bxl
		return -1, -1
	case 2:
		bst := combo(operand, *register) % 8
		register.B = bst
		return -1, -1
	case 3:
		if register.A != 0 {
			return -1, operand
		}
		return -1, -1
	case 4:
		bxc := register.B ^ register.C
		register.B = bxc
		return -1, -1
	case 5:
		out := combo(operand, *register) % 8
		return out, -1
	case 6:
		adv := register.A / (IntPow(2, (combo(operand, *register))))
		register.B = adv
		return -1, -1
	case 7:
		adv := register.A / (IntPow(2, (combo(operand, *register))))
		register.C = adv
		return -1, -1
	default:
		panic("Invalid instruction.")
	}
}

func run(items []int, register Register) []int {
	output := make([]int, 0)
	for i := 0; i < len(items)-1; {
		instruction := items[i]
		operand := items[i+1]
		out, jump := do(instruction, operand, &register)

		if out != -1 {
			output = append(output, out)
		}

		if jump != -1 {
			i = jump
		} else {
			i += 2
		}
	}
	return output
}

func main() {
	input, err := os.ReadFile("input.txt")
	check(err)

	register_regex := regexp.MustCompile(`Register (A|B|C): (\d+)`)
	register_match := register_regex.FindAllStringSubmatch(string(input), 3)
	register := Register{
		A: to_int(register_match[0][2]),
		B: to_int(register_match[1][2]),
		C: to_int(register_match[2][2]),
	}

	program_regex := regexp.MustCompile(`Program: ([\d,]+)`)
	program_match := program_regex.FindStringSubmatch(string(input))

	items := strings.Split(program_match[1], ",")

	code := to_ints(items)
	output := run(code, register)

	fmt.Print("part 1 - ", arrayToString(output, ","), "\n")

	for i := IntPow(8, len(input)-1); ; {
		test_register := Register{A: i, B: 0, C: 0}
		output := run(code, test_register)

		prod := 1
		for i := len(output) - 1; i >= 0; i-- {
			if output[i] != code[i] {
				prod = IntPow(8, i)
				break
			}
		}

		if arrayToString(output, ",") == program_match[1] {
			fmt.Print("part 2 - ", i, "\n")
			break
		}

		i += prod
	}
}
