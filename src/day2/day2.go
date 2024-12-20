package day2

import (
	"aoc_go_2024/src/utils"
	"log"
	"strconv"
	"strings"
)

type listOfList [][]int

type Day2 struct{}

func (d Day2) Parse(lines []string) (matrix listOfList) {
	matrix = make(listOfList, 0, len(lines))
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

func checkJumpIssue(prevVal, val int) bool {
	abs := utils.Abs(val - prevVal)
	return abs < 1 || 3 < abs
}

func checkIncDecIssue(prevVal, val int, isIncreasing bool) bool {
	return (isIncreasing && val < prevVal) || (!isIncreasing && val > prevVal)
}

func isSafe(row []int) bool {
	if len(row) < 2 {
		return true
	}

	prevVal := row[0]
	curVal := row[1]

	if checkJumpIssue(prevVal, curVal) {
		return false
	}

	isIncreasing := prevVal < curVal
	prevVal = curVal
	for _, val := range row[2:] {
		if checkJumpIssue(prevVal, val) || checkIncDecIssue(prevVal, val, isIncreasing) {
			return false
		}

		prevVal = val
	}

	return true
}

func (d Day2) Part1(debug bool, input listOfList) (count int64) {
	count = 0
	for _, row := range input {
		if isSafe(row) {
			count += 1
		}
	}
	return
}

func createAndRemove(row []int, i int) []int {
	newRow := make([]int, len(row)-1)
	copy(newRow[:i], row[:i])
	copy(newRow[i:], row[(i+1):])

	return newRow
}

func isSafeWithoutPreviouses(row []int, i int) bool {
	return isSafe(createAndRemove(row, i-2)) || isSafe(createAndRemove(row, i-1)) || isSafe(createAndRemove(row, i))
}

func isAlmostSafe(row []int) bool {
	if len(row) < 2 {
		return true
	}

	prevVal := row[0]
	curVal := row[1]

	if checkJumpIssue(prevVal, curVal) {
		if isSafe(createAndRemove(row, 0)) {
			return true
		}

		return isSafe(createAndRemove(row, 1))
	}

	isIncreasing := prevVal < curVal
	prevVal = curVal
	for i, val := range row[2:] {
		if checkJumpIssue(prevVal, val) || checkIncDecIssue(prevVal, val, isIncreasing) {
			return isSafeWithoutPreviouses(row, i+2)
		}

		prevVal = val
	}

	return true
}

func (d Day2) Part2(debug bool, input listOfList) (count int64) {
	count = 0
	for _, row := range input {
		if isAlmostSafe(row) {
			count += 1
		}
	}
	return
}
