package day4

import (
	"aoc_go_2024/src/utils"
)

type grid = utils.Grid[rune]
type Day4 struct{}

func (d Day4) GetNumber() uint {
	return 4
}

func subParse(val rune) (rune, error) {
	return val, nil
}

func (d Day4) Parse(lines []string) grid {
	return utils.Parse(lines, subParse)
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

type pt = utils.Position

func (d Day4) Part1(input grid) int64 {
	ctx := xmas{0, "XMAS", 0}
	directions := []pt{{X: -1, Y: 0}, {X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1}, {X: -1, Y: 1}}

	for i, row := range input {
		for j, car := range row {
			ctx.reset()
			found_x := ctx.advance(car)

			if !found_x {
				continue
			}

			xPt := pt{X: i, Y: j}
			for _, curDir := range directions {
				ctx.reset()
				ctx.advance(car)

				curPt := xPt.Add(curDir)
				curRune, ok := input.Get(curPt)
				for ok && ctx.advance(curRune) {
					curPt = curPt.Add(curDir)
					curRune, ok = input.Get(curPt)
				}
			}
		}
	}

	return int64(ctx.count)
}

func xmasOposite(a, b rune) bool {
	if a == rune('M') {
		return b == rune('S')
	} else if a == rune('S') {
		return b == rune('M')
	} else {
		return false
	}
}

func (d Day4) Part2(input grid) (count int64) {
	aRune := rune('A')
	count = 0
	leftUpDir, rightDownPt := pt{X: -1, Y: -1}, pt{X: 1, Y: 1}
	leftDownDir, rightUpPt := pt{X: 1, Y: -1}, pt{X: -1, Y: 1}
	for i, row := range input {
		for j, car := range row {
			if car != aRune {
				continue
			}

			xPt := pt{X: i, Y: j}

			leftUpPt := xPt.Add(leftUpDir)
			letter1, ok := input.Get(leftUpPt)
			if !ok {
				continue
			}

			rightDownPt := xPt.Add(rightDownPt)
			letter2, ok := input.Get(rightDownPt)
			if !ok || !xmasOposite(letter1, letter2) {
				continue
			}

			leftDownPt := xPt.Add(leftDownDir)
			letter1, ok = input.Get(leftDownPt)
			if !ok {
				continue
			}

			rightUpPt := xPt.Add(rightUpPt)
			letter2, ok = input.Get(rightUpPt)
			if !ok || !xmasOposite(letter1, letter2) {
				continue
			}

			count += 1
		}
	}
	return
}
