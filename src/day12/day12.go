package day12

import (
	"aoc_go_2024/src/utils"
	"fmt"
)

type grid = utils.Grid[rune]
type position = utils.Position
type visited = utils.Grid[bool]
type perimeterCell struct {
	in  position
	out position
}
type perimeterCells = map[perimeterCell]bool

type Day12 struct{}

func (d Day12) GetNumber() uint {
	return 12
}

func subParse(val rune) (rune, error) {
	return val, nil
}

func (d Day12) Parse(lines []string) grid {
	return utils.Parse(lines, subParse)
}

func solvePlant(input grid, allV visited, pCells perimeterCells, curPos position) int64 {
	// Visited current node
	curPlant, ok := input.Get(curPos)
	if !ok {
		// should never happen
		return 0
	}

	var area int64 = 0
	for _, dir := range utils.CardinalDirs {
		nextPos := curPos.Add(dir)
		nextPlant, ok := input.Get(nextPos)
		if !ok { // is outside
			pCells[perimeterCell{curPos, nextPos}] = true
		} else if curPlant != nextPlant {
			pCells[perimeterCell{curPos, nextPos}] = true
		} else if allV[nextPos.X][nextPos.Y] {
			continue
		} else {
			allV[nextPos.X][nextPos.Y] = true
			newArea := solvePlant(input, allV, pCells, nextPos)
			area += newArea
		}

	}
	return area + 1
}

func createVisited(input grid) visited {
	v := make(visited, 0, input.Height())
	for i := 0; i < input.Height(); i++ {
		v = append(v, make(utils.Row[bool], input.Width()))
	}
	return v
}

func computePerimeter(pCells perimeterCells) int64 {
	return int64(len(pCells))
}

func computeSides(pCells perimeterCells) int64 {
	return int64(len(pCells))
}

func solve(debug bool, input grid, isPart1 bool) (sum int64) {
	allV := createVisited(input)
	sum = 0
	for i, row := range input {
		for j, plant := range row {
			if allV[i][j] {
				continue
			}
			if debug {
				fmt.Println("Process ", string(plant))
			}

			pCells := make(perimeterCells)

			allV[i][j] = true
			area := solvePlant(input, allV, pCells, position{X: i, Y: j})

			if isPart1 {
				perimeter := computePerimeter(pCells)
				sum += area * perimeter
				if debug {
					fmt.Println("Area:", area, ", Perimeter:", perimeter)
				}
			} else {
				sides := computeSides(pCells)
				sum += area * sides
				if debug {
					fmt.Println("Area:", area, ", Sides:", sides)
				}
			}
		}
	}
	return
}

func (d Day12) Part1(debug bool, input grid) int64 {
	return solve(debug, input, true)
}

func (d Day12) Part2(debug bool, input grid) int64 {
	return solve(debug, input, false)
}
