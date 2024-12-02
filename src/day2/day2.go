package day2

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parse(lines []string) (matrix [][]int) {
	matrix = make([][]int, 0, len(lines))
	separator := " "
	for _, line := range lines {
		row_str := strings.Split(line, separator)
		row := make([]int, 0, len(row_str))

		for _, elem := range row_str {
			val, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
				continue
			}

			row = append(row, val)
		}

		matrix = append(matrix, row)
	}

	return
}

func isSafe(row []int) bool {
	if len(row) < 2 {
		return true
	}

	isIncreasing := true
	prevVal := row[0]

	for i, val := range row[1:] {
		if i == 0 {
			isIncreasing = val > prevVal
		} else {
			if isIncreasing && val < prevVal {
				return false
			} else if !isIncreasing && val > prevVal {
				return false
			}
		}

		if abs := utils.Abs(val - prevVal); abs < 1 || 3 < abs {
			return false
		}
		prevVal = val
	}

	return true
}

func day2_1(input [][]int) {
	fmt.Println(input)

	count := 0
	for _, row := range input {
		if isSafe(row) {
			count += 1
		}
	}

	fmt.Println(count)
}

func day2_2(input [][]int) {
	count := 0
	for _, row := range input {
		if isAlmostSafe(row) {
			count += 1
		}
	}

	fmt.Println(count)
}

func Day2(justATest bool) {
	fmt.Println("Welcome to day 2!!!")

	lines := utils.GetLines(justATest, 2)
	input := parse(lines)

	day2_1(input)
	day2_2(input)
}
