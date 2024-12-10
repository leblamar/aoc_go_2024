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
	Add    operation = 0
	Mul    operation = 1
	Concat operation = 2
)

type equation struct {
	testValue    int64
	values       []int64
	potentialOps []operation
}

func parse(lines []string) []equation {
	equations := make([]equation, 0)
	for _, line := range lines {
		splitTwoPoints := strings.Split(line, ": ")
		testValue, err := strconv.ParseInt(splitTwoPoints[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		splitValues := strings.Split(splitTwoPoints[1], " ")
		values := make([]int64, 0, len(splitValues))
		for _, valStr := range splitValues {
			val, err := strconv.ParseInt(valStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			values = append(values, val)
		}
		equation := equation{testValue, values, []operation{}}
		equations = append(equations, equation)
	}

	return equations
}

func (op operation) compute(a, b int64) int64 {
	if op == Add {
		return a + b
	} else if op == Mul {
		return a * b
	} else if op == Concat {
		aStr := strconv.FormatInt(a, 10)
		bStr := strconv.FormatInt(b, 10)
		resStr := aStr + bStr
		res, err := strconv.ParseInt(resStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		return res
	} else {
		panic(fmt.Sprintf("This operator is not implemented yet: %d", op))
	}
}

func (eq equation) isResolvable(nextIdx int, accValue int64) bool {
	// Termination
	if accValue > eq.testValue {
		return false
	} else if nextIdx == len(eq.values) {
		return eq.testValue == accValue
	}

	nextValue := eq.values[nextIdx]
	for _, op := range eq.potentialOps {
		nextAccValue := op.compute(accValue, nextValue)
		if eq.isResolvable(nextIdx+1, nextAccValue) {
			return true
		}
	}

	return false
}

func sumCorrectEquations(equations []equation, potentialOps []operation, isDebug bool) (sum int64) {
	sum = 0
	for i, eq := range equations {
		eq.potentialOps = potentialOps
		if isDebug {
			fmt.Println("[", i, "/", len(equations), "]", "New eq:", eq)
		}
		if eq.isResolvable(1, eq.values[0]) {
			if isDebug {
				fmt.Println("It is resolvable!!")
			}
			sum += eq.testValue
		}
	}

	return
}

func day7_1(equations []equation, isDebug bool) {
	potentialOps := []operation{Add, Mul}
	res := sumCorrectEquations(equations, potentialOps, isDebug)
	fmt.Println("Part 1:", res)
}

func day7_2(equations []equation, isDebug bool) {
	potentialOps := []operation{Add, Mul, Concat}
	res := sumCorrectEquations(equations, potentialOps, isDebug)
	fmt.Println("Part 2:", res)
}

func Day7(isTest, isDebug bool) {
	fmt.Println("Welcome to day 7!!!")

	lines := utils.GetLines(isTest, 7)
	equations := parse(lines)

	day7_1(equations, isDebug)
	day7_2(equations, isDebug)
}
