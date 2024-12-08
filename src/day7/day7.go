package day7

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type operation int

const (
	Add operation = 0
	Mul operation = 1
)

type equation struct {
	testValue int64
	values    []int64
}

func parse(lines []string) []equation {
	equations := make([]equation, 0)
	for _, line := range lines {
		splitTwoPoints := strings.Split(line, ": ")
		testValue, err := strconv.Atoi(splitTwoPoints[0])
		if err != nil {
			log.Fatal(err)
		}

		splitValues := strings.Split(splitTwoPoints[1], " ")
		values := make([]int64, 0, len(splitValues))
		for _, valStr := range splitValues {
			val, err := strconv.Atoi(valStr)
			if err != nil {
				log.Fatal(err)
			}

			values = append(values, int64(val))
		}
		equation := equation{int64(testValue), values}
		equations = append(equations, equation)
	}

	return equations
}

func (eq equation) isValidWithOps(ops []operation) bool {
	sum := eq.values[0]
	for i, val := range eq.values[1:] {
		if i >= len(ops) {
			log.Fatal("Why i:", i, " is sup to ops:", ops)
		}
		if ops[i] == Add {
			sum += val
		} else if ops[i] == Mul {
			sum *= val
		} else {
			log.Fatal("This operation does not exist yet:", ops[i])
		}
	}
	return sum == eq.testValue
}

func (eq equation) isResolvable(ops *[]operation) bool {
	// Termination
	if len(*ops) == len(eq.values)-1 {
		return eq.isValidWithOps(*ops)
	}

	*ops = append(*ops, Add)
	if eq.isResolvable(ops) {
		return true
	}

	(*ops)[len(*ops)-1] = Mul
	if eq.isResolvable(ops) {
		return true
	}

	if len(*ops) == 1 {
		return false
	}

	*ops = (*ops)[:len(*ops)-1]

	return false
}

func day7_1(equations []equation) {
	var sum int64 = 0
	for i, eq := range equations {
		curOps := make([]operation, 0, len(eq.values))
		fmt.Println("[", i, "/", len(equations), "]", "New eq:", eq)
		if eq.isResolvable(&curOps) {
			fmt.Println("It is resolvable!!")
			sum += eq.testValue
		}
	}
	fmt.Println("Part 1:", sum)
}

func day7_2(equations []equation) {
	fmt.Println("Part 2:")
}

func Day7(justATest bool) {
	fmt.Println("Welcome to day 7!!!")

	lines := utils.GetLines(justATest, 7)
	equations := parse(lines)

	day7_1(equations)
	day7_2(equations)
}
