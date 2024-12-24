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

func subParse(pos position, val rune) (rune, error) {
	return val, nil
}

func (d Day12) Parse(lines []string) grid {
	return utils.ParseGrid(lines, subParse)
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

// Return true if found at least one friend in the direction
func walkInDir(pCells map[position]bool, pos, dir position, v map[position]bool) (position, bool) {
	curPos := pos
	for pCells[curPos] {
		v[curPos] = true
		nextPos := curPos.Add(dir)

		if !pCells[nextPos] {
			break
		}
		curPos = nextPos
	}
	return curPos, curPos != pos
}

func isSideInDirection(pCells map[position]bool, pos position, isHorizontal bool, v map[position]bool) (prev, next position, isInDir bool) {
	var prevDir position
	var nextDir position
	if isHorizontal {
		prevDir = position{X: 0, Y: -1}
		nextDir = position{X: 0, Y: 1}
	} else {
		prevDir = position{X: -1, Y: 0}
		nextDir = position{X: 1, Y: 0}
	}

	next, hasNext := walkInDir(pCells, pos, nextDir, v)
	prev, hasPrev := walkInDir(pCells, pos, prevDir, v)

	isInDir = hasNext || hasPrev
	return
}

func walkForm(pCells, outV map[position]bool, pos position) []position {
	formV := map[position]bool{}
	queue := []position{pos}
	pFormInOrder := []position{}

	for len(queue) != 0 {
		curPos := queue[0]
		queue = queue[1:]
		pFormInOrder = append(pFormInOrder, curPos)
		outV[curPos] = true
		formV[curPos] = true
		for _, dir := range utils.WithDiagDirs {
			nextPos := curPos.Add(dir)
			if !pCells[nextPos] {
				continue
			} else if formV[nextPos] {
				continue
			} else {
				queue = append(queue, nextPos)
			}
		}
	}

	return pFormInOrder
}

func sameSide(pos1, pos2 position) bool {
	return pos1.X == pos2.X || pos1.Y == pos2.Y
}

func nbIn(inPCells map[position]bool, pos position) (in int64) {
	in = 0
	for _, dir := range utils.CardinalDirs {
		nextPos := pos.Add(dir)
		if inPCells[nextPos] {
			in += 1
		}
	}
	return
}

func nbSides(inPCells map[position]bool, pFormInOrder []position) (sides int64) {
	sides = 0
	if len(pFormInOrder) == 0 {
		return
	} else if len(pFormInOrder) == 1 {
		sides = 4
		return
	}

	prevPos := pFormInOrder[0]
	for _, curPos := range pFormInOrder[1:] {
		if sameSide(prevPos, curPos) {
			prevPos = curPos
			continue
		} else {
			prevPos = curPos
			sides += nbIn(inPCells, curPos)
		}
	}

	if !sameSide(prevPos, pFormInOrder[0]) {
		sides += 1
	}
	return
}

func subComputeSides(outPCells, inPCells map[position]bool) (sides int64) {
	v := make(map[position]bool)
	sides = 0
	for pos := range outPCells {
		if v[pos] {
			continue
		}

		//pFormInOrder := walkForm(outPCells, v, pos)
		//sides += nbSides(inPCells, pFormInOrder)

		_, _, isHorizontal := isSideInDirection(outPCells, pos, true, v)
		if isHorizontal {
			sides += 1
			continue
		}

		_, _, isVertical := isSideInDirection(outPCells, pos, false, v)
		if isVertical {
			sides += 1
			continue
		}

		sides += nbIn(inPCells, pos)
	}
	return
}

func computeSides(pCells perimeterCells) int64 {
	outPCells := make(map[position]bool)
	inPCells := make(map[position]bool)
	for pCell := range pCells {
		outPCells[pCell.out] = true
		inPCells[pCell.in] = true
	}
	// Computing internal sides of a form is equivalent to compute external sides of the internal form
	return subComputeSides(outPCells, inPCells)
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
