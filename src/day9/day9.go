package day9

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
)

type day9 struct {
	input []int
}

func parse(lines []string) day9 {
	if len(lines) > 1 && len(lines) == 0 {
		log.Fatal("Bad input")
	}

	input := make([]int, 0, len(lines[0]))
	for _, car := range lines[0] {
		input = append(input, int(car-'0'))
	}

	return day9{input}
}

func (pg *day9) day9_1(debug bool) {
	count := 0
	for _, elem := range pg.input {
		count += elem
	}

	disk := make([]int16, 0, count)
	isFree := false
	ID := 0
	for _, elem := range pg.input {
		var valToAppend int16
		if isFree {
			valToAppend = -1
		} else {
			valToAppend = int16(ID)
		}
		for i := 0; i < elem; i++ {
			disk = append(disk, valToAppend)
		}
		isFree = !isFree
		if !isFree {
			ID++
		}
	}
	if debug {
		fmt.Println("Disk:", disk)
	}

	inc, dec := 0, len(disk)-1
	for inc < dec {
		if disk[inc] != -1 {
			inc++
			continue
		} else if disk[dec] == -1 {
			dec--
			continue
		}

		disk[inc] = disk[dec]
		disk[dec] = -1
		inc++
		dec--
	}
	if debug {
		fmt.Println("Disk:", disk)
	}

	sum := 0
	for i, id := range disk {
		if id == -1 {
			break
		}
		if debug {
			fmt.Println("i * id:", i, "*", id, " = ", i*int(id))
		}
		sum += i * int(id)
	}
	fmt.Println("Part1:", sum)
}

func (pg *day9) day9_2() {
	fmt.Println("Part2:")
}

func Day9(isTest, debug bool) {
	fmt.Println("Welcome to day 9!!!")

	lines := utils.GetLines(isTest, 9)
	day9 := parse(lines)

	day9.day9_1(debug)
	day9.day9_2()
}
