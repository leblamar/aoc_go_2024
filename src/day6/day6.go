package day6

import (
	"aoc_go_2024/src/utils"
	"bytes"
	"errors"
	"fmt"
	"time"
)

type caseState rune

const (
	Visited    caseState = 'X'
	Block      caseState = '#'
	NotVisited caseState = '.'
)

type direction int

const (
	Up    direction = 0
	Right direction = 1
	Down  direction = 2
	Left  direction = 3
)

type labMap = [][]caseState

type position struct {
	x int
	y int
}

type player struct {
	pos position
	dir direction
}

type program struct {
	m     labMap
	pl    player
	count int
}

type status int

const (
	Finished status = 0
	Running  status = 1
	Looping  status = 2
)

var toRune = [4]rune{'^', '>', 'v', '<'}

func (dir direction) rune() rune {
	return toRune[int(dir)]
}

func parseCaseState(car rune) caseState {
	if car == Up.rune() || car == Right.rune() || car == Down.rune() || car == Left.rune() {
		return Visited
	} else {
		return caseState(car)
	}
}

func parsePlayer(i, j int, car rune) (player, error) {
	if car == Up.rune() {
		return player{position{i, j}, Up}, nil
	} else if car == Right.rune() {
		return player{position{i, j}, Right}, nil
	} else if car == Down.rune() {
		return player{position{i, j}, Down}, nil
	} else if car == Left.rune() {
		return player{position{i, j}, Left}, nil
	} else {
		return player{}, errors.New("not a player")
	}
}

func parse(lines []string) program {
	m := make(labMap, 0, len(lines))
	p := player{}
	for i, line := range lines {
		mRow := make([]caseState, 0, len(line))
		for j, car := range line {
			mRow = append(mRow, parseCaseState(car))
			curP, err := parsePlayer(i, j, car)
			if err == nil {
				p = curP
			}
		}
		m = append(m, mRow)
	}

	return program{m, p, 1}
}

func (pg program) String() string {
	var b bytes.Buffer
	for i, row := range pg.m {
		for j, car := range row {
			if pg.pl.pos.x == i && pg.pl.pos.y == j {
				b.WriteRune(pg.pl.dir.rune())
			} else {
				b.WriteRune(rune(car))
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (dir direction) getDir() position {
	if dir == Up {
		return position{-1, 0}
	} else if dir == Right {
		return position{0, 1}
	} else if dir == Down {
		return position{1, 0}
	} else if dir == Left {
		return position{0, -1}
	} else {
		return position{0, 0}
	}
}

func (dir direction) getNext() direction {
	return (dir + 1) % 4
}

func (p player) getNextPos() position {
	curDir := p.dir.getDir()
	return position{p.pos.x + curDir.x, p.pos.y + curDir.y}
}

func (pg *program) insideBoundaries(pos position) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < len(pg.m) && pos.y < len(pg.m[0])
}

// Return true if is finished
func (pg *program) move() status {
	nextPos := pg.pl.getNextPos()
	if !pg.insideBoundaries(nextPos) {
		return Finished
	}

	if pg.m[nextPos.x][nextPos.y] == Block {
		pg.pl.dir = pg.pl.dir.getNext()
		return pg.move()
	} else if pg.m[nextPos.x][nextPos.y] == NotVisited {
		pg.count += 1
		pg.m[nextPos.x][nextPos.y] = Visited
	}

	pg.pl.pos = nextPos
	return Running
}

func (pg program) testRun(withPrint bool) (int, status) {
	if withPrint {
		fmt.Println(pg)
	}
	var st status = Running
	for st != Finished {
		if withPrint {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(pg)
		}
		st = pg.move()
	}

	return pg.count, st
}

func (pg program) day6_1(withPrint bool) {
	count, _ := pg.testRun(withPrint)
	fmt.Println("Part 1:", count)
}

func (pg program) day6_2() {
	fmt.Println("Part 2:")
}

func Day6(justATest, debug bool) {
	fmt.Println("Welcome to day 6!!!")

	lines := utils.GetLines(justATest, 6)
	pg := parse(lines)

	pg.day6_1(debug)
	pg.day6_2()
}
