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
	result   [][]bool
}

func parse(lines []string) day8 {
	radioMap := make(map[frequency][]position)
	matrix := make([][]frequency, 0, len(lines))

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
	}

	return day8{radioMap, matrix, [][]bool{}}
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

func (pg *day8) mark(p position) {
	pg.result[p.X][p.Y] = true
}

func (pg *day8) drawSym(positions []position) {
	for i, pos1 := range positions {
		for j, pos2 := range positions {
			if i == j {
				// No sym of himself
				continue
			}

			symPos := pos1.Sym(pos2)
			_, ok := pg.get(symPos)
			if ok {
				pg.mark(symPos)
			}
		}
	}
}

func (pg day8) String() string {
	var b bytes.Buffer
	for i, row := range pg.matrix {
		for j, freq := range row {
			if pg.result[i][j] { // && freq == '.' {
				b.WriteRune('#')
			} else {
				b.WriteRune(freq)
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (pg *day8) getResult() int {
	count := 0
	for _, row := range pg.result {
		for _, isOk := range row {
			if isOk {
				count++
			}
		}
	}

	return count
}

func (pg *day8) reset() {
	pg.result = make([][]bool, 0, len(pg.matrix))
	for _, row := range pg.matrix {
		pg.result = append(pg.result, make([]bool, len(row)))
	}
}

func (pg *day8) day8_1() {
	pg.reset()
	for _, positions := range pg.radioMap {
		pg.drawSym(positions)
	}

	count := pg.getResult()
	fmt.Println("Part 1:", count)
}

func (pg *day8) drawRepSym(positions []position) {
	for i, pos1 := range positions {
		for j, pos2 := range positions {
			if i == j {
				// No sym of himself
				continue
			}

			pg.mark(pos1)
			pg.mark(pos2)
			tmpPos1 := pos1
			tmpPos2 := pos2
			symPos := tmpPos1.Sym(tmpPos2)
			_, ok := pg.get(symPos)
			for ok {
				pg.mark(symPos)
				tmpPos1 = tmpPos2
				tmpPos2 = symPos
				symPos = tmpPos1.Sym(tmpPos2)
				_, ok = pg.get(symPos)
			}
		}
	}
}

func (pg *day8) day8_2() {
	pg.reset()
	for _, positions := range pg.radioMap {
		pg.drawRepSym(positions)
	}

	count := pg.getResult()
	fmt.Println("Part 2:", count)
}

func Day8(isTest bool) {
	fmt.Println("Welcome to day 8!!!")

	lines := utils.GetLines(isTest, 8)
	day8 := parse(lines)

	day8.day8_1()
	day8.day8_2()
}
