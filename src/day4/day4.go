package day4

import (
	"aoc_go_2024/src/utils"
	"errors"
	"fmt"
)

func parse(lines []string) [][]rune {
	matrix := make([][]rune, 0, len(lines))
	for _, line := range lines {
		row := make([]rune, 0, len(line))
		for _, elem := range line {
			row = append(row, elem)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

type xmas struct {
	step  int
	chain string
	count uint
}

func (ctx *xmas) reset() {
	ctx.step = 0
}

func (ctx *xmas) cur() rune {
	if ctx.step >= len(ctx.chain) {
		ctx.reset()
	}
	return rune(ctx.chain[ctx.step])
}

func (ctx *xmas) advance(car rune) bool {
	if ctx.cur() != car {
		ctx.reset()
		return false
	}

	ctx.step += 1

	if ctx.step == len(ctx.chain) {
		ctx.count += 1
		ctx.reset()
		return false
	} else {
		return true
	}
}

type pt struct {
	x int
	y int
}

func (cur pt) add(other pt) pt {
	return pt{cur.x + other.x, cur.y + other.y}
}

func (cur pt) getRune(input [][]rune) (rune, error) {
	if cur.x < 0 || len(input) <= cur.x || cur.y < 0 || len(input[0]) <= cur.y {
		return rune(0), errors.New("not in matrix")
	}
	return input[cur.x][cur.y], nil
}

func day4_1(input [][]rune) {
	ctx := xmas{0, "XMAS", 0}
	directions := []pt{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}

	for i, row := range input {
		for j, car := range row {
			ctx.reset()
			found_x := ctx.advance(car)

			if !found_x {
				continue
			}

			xPt := pt{i, j}
			for _, curDir := range directions {
				ctx.reset()
				ctx.advance(car)

				curPt := xPt.add(curDir)
				curRune, err := curPt.getRune(input)
				for err == nil && ctx.advance(curRune) {
					curPt = curPt.add(curDir)
					curRune, err = curPt.getRune(input)
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("Part 1 :", ctx.count)
}

func day4_2(input [][]rune) {
	fmt.Println("Part 2 :")
}

func Day4(justATest bool) {
	fmt.Println("Welcome to day 4!!!")

	lines := utils.GetLines(justATest, 4)
	input := parse(lines)

	day4_1(input)
	day4_2(input)
}
