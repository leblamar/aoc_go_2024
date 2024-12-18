package day1

import (
	"aoc_go_2024/src/utils"
	"sort"
	"strconv"
	"strings"
)

type lists struct {
	left  []int
	right []int
}

type Day1 struct{}

func (d Day1) GetNumber() uint {
	return 1
}

func (d Day1) Parse(lines []string) lists {
	left_list := make([]int, 0)
	right_list := make([]int, 0)
	separator := "   "
	for _, line := range lines {
		splitted_line := strings.Split(line, separator)

		if len(splitted_line) != 2 {
			continue
		}

		left_val, left_err := strconv.Atoi(splitted_line[0])
		right_val, right_err := strconv.Atoi(splitted_line[1])

		if left_err == nil && right_err == nil {
			left_list = append(left_list, left_val)
			right_list = append(right_list, right_val)
		}
	}

	return lists{left_list, right_list}
}

func (d Day1) Part1(debug bool, input lists) int64 {
	left_list := input.left
	right_list := input.right

	if len(left_list) != len(right_list) {
		return -1
	}

	sort.Ints(left_list)
	sort.Ints(right_list)

	sum := 0
	for i := 0; i < len(left_list); i++ {
		sum += utils.Abs(left_list[i] - right_list[i])
	}

	return int64(sum)
}

func (d Day1) Part2(debug bool, input lists) int64 {
	left_list := input.left
	right_list := input.right

	right_count_map := make(map[int]int)
	for _, elem := range right_list {
		old_count, ok := right_count_map[elem]

		if ok {
			right_count_map[elem] = old_count + 1
		} else {
			right_count_map[elem] = 1
		}
	}

	sum := 0
	for _, elem := range left_list {
		count, ok := right_count_map[elem]

		if ok {
			sum += elem * count
		}
	}

	return int64(sum)
}
