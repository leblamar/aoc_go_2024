package day12

import (
	"aoc_go_2024/src/utils"
	"fmt"
)

type grid = utils.Grid[rune]
type position = utils.Position

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

func solve(input grid, allV visited, i, j int) (int64, int64) {
	// Visited current node
	curPlant, ok := input.Get(position{X: i, Y: j})
	if !ok {
		// should never happen
		return 0, 0
	}

	var area, perimeter int64 = 0, 0
	for _, dir := range utils.CardinalDirs {
		nextPos := position{X: i, Y: j}.Add(dir)
		nextPlant, ok := input.Get(nextPos)
		if !ok { // is outside
			perimeter += 1
		} else if curPlant != nextPlant {
			perimeter += 1
		} else if allV[nextPos.X][nextPos.Y] {
			continue
		} else {
			allV[nextPos.X][nextPos.Y] = true
			newArea, newPerimeter := solve(input, allV, nextPos.X, nextPos.Y)
			area += newArea
			perimeter += newPerimeter
		}

	}
	return area + 1, perimeter
}

type visited = utils.Grid[bool]

func createVisited(input grid) visited {
	v := make(visited, 0, input.Height())
	for i := 0; i < input.Height(); i++ {
		v = append(v, make(utils.Row[bool], input.Width()))
	}
	return v
}

func (d Day12) Part1(debug bool, input grid) (sum int64) {
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

			allV[i][j] = true
			area, perimeter := solve(input, allV, i, j)
			sum += area * perimeter
			if debug {
				fmt.Println("Area:", area, ", Perimeter:", perimeter)
			}
		}
	}
	return
}

func (d Day12) Part2(debug bool, input grid) int64 {
	return 0
}
