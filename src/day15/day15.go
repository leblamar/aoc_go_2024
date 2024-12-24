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

func (d Day15) Part1(debug bool, g game) int64 {
	if debug {
		fmt.Println(g)
	}
	for len(g.moves) != 0 {
		curMove := g.moves[0]
		newMoves := g.moves[1:]
		g.moves = newMoves

		g.apply(curMove)
		if debug {
			fmt.Println(g)
		}
	}

	sum := 0
	for box := range g.boxSet {
		sum += box.X + 100*box.Y
	}
	return int64(sum)
}

func (d Day15) Part2(debug bool, g game) int64 {
	return 0
}
