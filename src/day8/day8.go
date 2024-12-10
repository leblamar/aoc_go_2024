package day8

import (
	"aoc_go_2024/src/utils"
	"bytes"
	"fmt"
)

type position = utils.Position
type frequency = rune

const (
	nothing frequency = '.'
	sym     frequency = '#'
)

type day8 struct {
	radioMap map[frequency][]position
	matrix   [][]frequency
	result   map[frequency]int
	result2  [][]bool
}

func parse(lines []string) day8 {
	radioMap := make(map[frequency][]position)
	matrix := make([][]frequency, 0, len(lines))
	result2 := make([][]bool, 0, len(lines))

	for i, line := range lines {
		row := make([]frequency, 0, len(line))
		for j, car := range line {
			row = append(row, car)
			if car == nothing {
				continue
			}

			val, ok := radioMap[car]
			if !ok {
				val = make([]position, 0)
			}

			val = append(val, position{X: i, Y: j})
			radioMap[car] = val
		}

		matrix = append(matrix, row)
		result2 = append(result2, make([]bool, len(line)))
	}

	return day8{radioMap, matrix, make(map[frequency]int), result2}
}

func (pg *day8) get(p position) (frequency, bool) {
	if len(pg.matrix) == 0 || len(pg.matrix[0]) == 0 {
		return nothing, false
	} else if p.X < 0 || p.Y < 0 {
		return nothing, false
	} else if p.X >= len(pg.matrix) || p.Y >= len(pg.matrix[0]) {
		return nothing, false
	} else {
		return pg.matrix[p.X][p.Y], true
	}
}

func (pg *day8) mark(freq frequency, p position) {
	pg.result[freq] += 1
	pg.result2[p.X][p.Y] = true
}

func (pg *day8) drawSym(freq frequency, positions []position) {
	for i, pos1 := range positions {
		for j, pos2 := range positions {
			if i == j {
				// No sym of himself
				continue
			}

			symPos := pos1.Sym(pos2)
			//fmt.Println("Sym:", symPos)
			_, ok := pg.get(symPos)
			if ok {
				pg.mark(freq, symPos)
			}
		}
	}
}

func (pg day8) String() string {
	var b bytes.Buffer
	for _, row := range pg.matrix {
		for _, freq := range row {
			b.WriteRune(freq)
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (pg day8) day8_1() {
	for freq, positions := range pg.radioMap {
		pg.drawSym(freq, positions)
	}

	count2 := 0
	for _, countFreq := range pg.result {
		count2 += countFreq
	}

	count3 := 0
	for _, row := range pg.result2 {
		for _, isOk := range row {
			if isOk {
				count3++
			}
		}
	}
	fmt.Println("Part 1:", count2, ", ", count3)
}
func (pg day8) day8_2() {
	fmt.Println("Part 2:")
}

func Day8(isTest bool) {
	fmt.Println("Welcome to day 8!!!")

	lines := utils.GetLines(isTest, 8)
	day8 := parse(lines)

	day8.day8_1()
	day8.day8_2()
}
