package day10

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"time"
)

type position = utils.Position
type Visited map[position]bool

type graph struct {
	mat   utils.Grid[int]
	zeros []position
}

func parse(lines []string) graph {
	matrix := make([][]int, 0, len(lines))
	zeros := make([]position, 0)
	for i, line := range lines {
		row := make([]int, 0, len(line))
		for j, car := range line {
			row = append(row, int(car-'0'))
			if car == '0' {
				zeros = append(zeros, position{X: i, Y: j})
			}
		}

		matrix = append(matrix, row)
	}

	return graph{matrix, zeros}
}

func (v Visited) makeCopy() Visited {
	newV := make(Visited, len(v))
	for pos, visited := range v {
		newV[pos] = visited
	}
	return newV
}

func (g graph) foundNines(visited Visited, pos position, nines Visited) {
	curVal, ok := g.mat.Get(pos)
	if !ok {
		return
	} else if curVal == 9 {
		nines[pos] = true
		return
	}

	for _, dir := range utils.CardinalDirs {
		nextPos := pos.Add(dir)
		nextVal, ok := g.mat.Get(nextPos)
		if !ok {
			continue
		} else if nextVal != curVal+1 {
			continue
		} else if visited[nextPos] {
			continue
		} else {
			newVisited := visited.makeCopy()
			newVisited[nextPos] = true
			g.foundNines(newVisited, nextPos, nines)
		}
	}
}

func (g graph) foundAllPath(visited Visited, pos position) int {
	curVal, ok := g.mat.Get(pos)
	if !ok {
		return 0
	} else if curVal == 9 {
		return 1
	}

	sum := 0
	for _, dir := range utils.CardinalDirs {
		nextPos := pos.Add(dir)
		nextVal, ok := g.mat.Get(nextPos)
		if !ok {
			continue
		} else if nextVal != curVal+1 {
			continue
		} else if visited[nextPos] {
			continue
		} else {
			newVisited := visited.makeCopy()
			newVisited[nextPos] = true
			sum += g.foundAllPath(newVisited, nextPos)
		}
	}

	return sum
}

func (g graph) day10_1() {
	sum := 0
	for _, pos := range g.zeros {
		nines := make(Visited)
		visited := make(Visited, 9)
		visited[pos] = true
		g.foundNines(visited, pos, nines)
		sum += len(nines)
	}
	fmt.Println("Part 1:", sum)
}

func (g graph) day10_2() {
	sum := 0
	for _, pos := range g.zeros {
		visited := make(Visited, 9)
		visited[pos] = true
		sum += g.foundAllPath(visited, pos)
	}
	fmt.Println("Part 2:", sum)
}

func Day10(isTest bool) {
	fmt.Println("Welcome to day 10!!!")

	lines := utils.GetLines(isTest, 10)
	day10 := parse(lines)

	start := time.Now()
	day10.day10_1()
	ellapsed := time.Since(start)

	fmt.Println("Time ellapsed : ", ellapsed)

	start = time.Now()
	day10.day10_2()
	ellapsed = time.Since(start)
	fmt.Println("Time ellapsed : ", ellapsed)
}
