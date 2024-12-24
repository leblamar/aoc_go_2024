package day15

import (
	"aoc_go_2024/src/utils"
	"bytes"
	"errors"
	"fmt"
	"log"
)

type Day15 struct{}

type block rune

const (
	Robot   block = '@'
	Wall    block = '#'
	Box     block = 'O'
	LBox    block = '['
	RBox    block = ']'
	Nothing block = '.'
)

type position = utils.Position
type direction = utils.Position
type game struct {
	m      utils.Grid[block]
	r      position
	boxSet map[position]bool
	moves  []direction
}

func (g game) copy() game {
	newG := game{}
	newG.m = g.m.Copy()
	newG.r = g.r

	newBoxSet := make(map[position]bool)
	for box := range g.boxSet {
		newBoxSet[box] = true
	}
	newG.boxSet = newBoxSet

	newMoves := make([]direction, 0, len(g.moves))
	newMoves = append(newMoves, g.moves...)
	newG.moves = newMoves

	return newG
}

func parseMoves(line string) []direction {
	moves := make([]direction, 0, len(line))
	for _, car := range line {
		val, err := utils.ParseDir(car)
		if err != nil {
			log.Fatal(err)
			continue
		}

		moves = append(moves, val)
	}

	return moves
}

func subParseAndFillMap(rockSet map[position]bool, robotPos *position) utils.ParseRune[block] {
	return func(pos position, r rune) (block, error) {
		if r == rune(Robot) {
			*robotPos = pos
			return Robot, nil
		} else if r == rune(Wall) {
			return Wall, nil
		} else if r == rune(Box) {
			rockSet[pos] = true
			return Box, nil
		} else if r == rune(Nothing) {
			return Nothing, nil
		}
		var b block
		return b, errors.New("this caracter cannot be parsed: " + string(r))
	}
}

func (d Day15) Parse(lines []string) game {
	rockSet := make(map[position]bool)
	var robotPos position
	m := utils.ParseGrid(lines[:len(lines)-2], subParseAndFillMap(rockSet, &robotPos))

	moves := parseMoves(lines[len(lines)-1])
	return game{m, robotPos, rockSet, moves}
}

func (g game) String() string {
	var b bytes.Buffer
	if len(g.moves) > 0 {
		dir, err := g.moves[0].ToRune()
		var dirStr string
		if err == nil {
			dirStr = string(dir)
		} else {
			dirStr = g.moves[0].String()
		}
		b.WriteString("Move " + dirStr + ":\n")
	} else {
		b.WriteString("No more moves:\n")
	}

	for j := 0; j < g.m.Height(); j++ {
		for i := 0; i < g.m.Width(); i++ {
			b.WriteRune(rune(g.m[j][i]))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (b block) isPushable() bool {
	return b != Wall
}

func (g game) tryPush(move position) bool {
	curPos := g.r
	count := 0
	for {
		nextPos := curPos.Add(move)

		b, ok := g.m.Get(nextPos)
		if !ok {
			return false
		} else if !b.isPushable() {
			return false
		} else if b == Nothing {
			break
		}
		count += 1
		curPos = nextPos
	}

	for i := count + 1; i > 0; i-- {
		curX := i*move.X + g.r.X
		curY := i*move.Y + g.r.Y
		curPos := position{X: curX, Y: curY}
		nextX := (i-1)*move.X + g.r.X
		nextY := (i-1)*move.Y + g.r.Y
		nextPos := position{X: nextX, Y: nextY}

		if !g.m.IsInside(curPos) || !g.m.IsInside(nextPos) {
			log.Fatal("this should never happened")
		}

		val, ok := g.m.Get(nextPos)
		if !ok {
			log.Fatal("this should never happened")
		}

		if val == Box {
			delete(g.boxSet, nextPos)
			g.boxSet[curPos] = true
		}
		g.m.Set(curPos, val)
	}

	g.m.Set(g.r, Nothing)

	return true
}

func (g *game) apply(move position) {
	nextPos := g.r.Add(move)

	val, ok := g.m.Get(nextPos)
	if !ok {
		return
	}

	switch val {
	case Nothing:
		g.m.Set(g.r, Nothing)
		g.r = nextPos
		g.m.Set(g.r, Robot)
	case Box:
		ok := g.tryPush(move)
		if ok {
			g.r = nextPos
		}
	default:
	}
}

func (g game) getRes() int64 {
	sum := 0
	for box := range g.boxSet {
		sum += box.X + 100*box.Y
	}

	return int64(sum)
}

func (d Day15) Part1(debug bool, g game) int64 {
	copiedG := g.copy()
	if debug {
		fmt.Println(copiedG)
	}
	for len(g.moves) != 0 {
		curMove := copiedG.moves[0]
		newMoves := copiedG.moves[1:]
		copiedG.moves = newMoves

		copiedG.apply(curMove)
		if debug {
			fmt.Println(copiedG)
		}
	}

	return copiedG.getRes()
}

func transformPos(pos position) position {
	return position{X: pos.X * 2, Y: pos.Y}
}

func (g game) transformPart2() game {
	newTG := game{}

	newM := make(utils.Grid[block], 0, len(g.m))
	for _, row := range g.m {
		newRow := make([]block, 0)
		for _, val := range row {
			switch val {
			case Nothing, Wall:
				newRow = append(newRow, val)
				newRow = append(newRow, val)
			case Box:
				newRow = append(newRow, LBox)
				newRow = append(newRow, RBox)
			case Robot:
				newRow = append(newRow, Robot)
				newRow = append(newRow, Nothing)
			default:
			}
		}
		newM = append(newM, newRow)
	}
	newTG.m = newM

	newTG.r = transformPos(g.r)

	newBoxSet := make(map[position]bool)
	for pos := range g.boxSet {
		newPos := transformPos(pos)
		newBoxSet[newPos] = true
	}
	newTG.boxSet = newBoxSet

	newMoves := make([]direction, 0, len(g.moves))
	newMoves = append(newMoves, g.moves...)
	newTG.moves = newMoves

	return newTG
}

func (d Day15) Part2(debug bool, g game) int64 {
	tG := g.transformPart2()
	if debug {
		fmt.Println(tG)
		return 0
	}
	for len(g.moves) != 0 {
		//curMove := tG.moves[0]
		newMoves := tG.moves[1:]
		tG.moves = newMoves

		//tG.applyV2(curMove)
		if debug {
			fmt.Println(tG)
		}
	}
	return tG.getRes()
}
