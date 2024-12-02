package day1

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type lists struct {
	left  []int
	right []int
}

func parse(lines []string) lists {
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

	fmt.Println(left_list)
	fmt.Println(right_list)

	return lists{left_list, right_list}
}

func day1_1(input lists) {
	left_list := input.left
	right_list := input.right

	if len(left_list) != len(right_list) {
		return
	}

	sort.Ints(left_list)
	sort.Ints(right_list)

	sum := 0
	for i := 0; i < len(left_list); i++ {
		sum += utils.Abs(left_list[i] - right_list[i])
	}

	fmt.Println(sum)
}

func day1_2(input lists) {
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

	fmt.Println(sum)
}

func Day1(justATest bool) {
	fmt.Println("Welcome to day 1!!!")

	lines := utils.GetLines(justATest, 1)
	input := parse(lines)

	day1_1(input)
	day1_2(input)
}
