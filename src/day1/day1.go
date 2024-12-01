package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func get_test_file_path(day uint) string {
	relative_path := fmt.Sprintf("src/day%d/day%d_example.txt", day, day)
	absPath, err := filepath.Abs(relative_path)
	if err != nil {
		log.Fatal(err)
	}
	return absPath
}

func get_input_file_path(day uint) string {
	relative_path := fmt.Sprintf("src/day%d/day%d_input.txt", day, day)
	absPath, err := filepath.Abs(relative_path)
	if err != nil {
		log.Fatal(err)
	}
	return absPath
}

func get_url(day uint) string {
	return fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func get_lines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func get_lines_from_dl() []string {
	path := get_input_file_path(1)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fileUrl := get_url(1)
		resp, err := http.Get(fileUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		newFile, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, resp.Body)

		if err != nil {
			log.Fatal(err)
		}
	}
	return get_lines(path)
}

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
		sum += Abs(left_list[i] - right_list[i])
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

func main() {
	fmt.Println("Welcome to day 1!!!")

	args := os.Args

	just_a_test := false
	if len(args) == 2 && args[1] == "test" {
		just_a_test = true
	}

	var lines []string
	if just_a_test {
		path := get_test_file_path(1)
		lines = get_lines(path)
	} else {
		lines = get_lines_from_dl()
	}

	input := parse(lines)

	day1_1(input)
	day1_2(input)
}
