package day3

import (
	"strings"
	"unicode"
)

type expression = string
type Day3 struct{}

func (d Day3) GetNumber() uint {
	return 3
}

func (d Day3) Parse(lines []string) expression {
	return strings.Join(lines, "")
}

type mulContext struct {
	mulLvl   uint
	firstNb  uint
	secondNb uint
	nbLvl    uint
}

func (ctx *mulContext) reset() {
	ctx.mulLvl = 0
	ctx.firstNb = 0
	ctx.secondNb = 0
	ctx.nbLvl = 0
}

func (d Day3) Part1(input expression) int64 {
	ctx := mulContext{0, 0, 0, 0}
	var multSum uint = 0
	for _, token := range input {
		if token == 'm' && ctx.mulLvl == 0 {
			ctx.mulLvl += 1
		} else if token == 'u' && ctx.mulLvl == 1 {
			ctx.mulLvl += 1
		} else if token == 'l' && ctx.mulLvl == 2 {
			ctx.mulLvl += 1
		} else if token == '(' && ctx.mulLvl == 3 {
			ctx.mulLvl += 1
		} else if unicode.IsDigit(token) && ctx.mulLvl == 4 && ctx.nbLvl < 4 {
			ctx.firstNb *= 10
			ctx.firstNb += uint(token - '0')
			ctx.nbLvl += 1
		} else if token == ',' && ctx.mulLvl == 4 && ctx.nbLvl > 0 {
			ctx.mulLvl += 1
			ctx.nbLvl = 0
		} else if unicode.IsDigit(token) && ctx.mulLvl == 5 && ctx.nbLvl < 4 {
			ctx.secondNb *= 10
			ctx.secondNb += uint(token - '0')
			ctx.nbLvl += 1
		} else if token == ')' && ctx.mulLvl == 5 && ctx.nbLvl > 0 {
			multSum += ctx.firstNb * ctx.secondNb
			ctx.reset()
		} else {
			ctx.reset()
		}
	}

	return int64(multSum)
}

type doContext struct {
	doLvl   uint
	dontLvl uint
	isDo    bool
}

func (ctx *doContext) reset() {
	ctx.doLvl = 0
	ctx.dontLvl = 0
	ctx.isDo = true
}

func (d Day3) Part2(input expression) int64 {
	ctx := mulContext{0, 0, 0, 0}
	doCtx := doContext{0, 0, true}
	var multSum uint = 0
	for _, token := range input {
		if token == 'd' {
			doCtx.doLvl = 1
			doCtx.dontLvl = 1
		} else if token == 'o' && doCtx.doLvl == 1 {
			doCtx.doLvl += 1
			doCtx.dontLvl += 1
		} else if token == '(' && doCtx.doLvl == 2 {
			doCtx.doLvl += 1
			doCtx.dontLvl = 0
		} else if token == ')' && doCtx.doLvl == 3 {
			doCtx.reset()
		} else if token == 'n' && doCtx.dontLvl == 2 {
			doCtx.dontLvl += 1
			doCtx.doLvl = 0
		} else if token == '\'' && doCtx.dontLvl == 3 {
			doCtx.dontLvl += 1
		} else if token == 't' && doCtx.dontLvl == 4 {
			doCtx.dontLvl += 1
		} else if token == '(' && doCtx.dontLvl == 5 {
			doCtx.dontLvl += 1
		} else if token == ')' && doCtx.dontLvl == 6 {
			doCtx.dontLvl = 0
			doCtx.isDo = false
		} else if token == 'm' {
			ctx.mulLvl = 1
		} else if token == 'u' && ctx.mulLvl == 1 {
			ctx.mulLvl += 1
		} else if token == 'l' && ctx.mulLvl == 2 {
			ctx.mulLvl += 1
		} else if token == '(' && ctx.mulLvl == 3 {
			ctx.mulLvl += 1
		} else if unicode.IsDigit(token) && ctx.mulLvl == 4 && ctx.nbLvl < 4 {
			ctx.firstNb *= 10
			ctx.firstNb += uint(token - '0')
			ctx.nbLvl += 1
		} else if token == ',' && ctx.mulLvl == 4 && ctx.nbLvl > 0 {
			ctx.mulLvl += 1
			ctx.nbLvl = 0
		} else if unicode.IsDigit(token) && ctx.mulLvl == 5 && ctx.nbLvl < 4 {
			ctx.secondNb *= 10
			ctx.secondNb += uint(token - '0')
			ctx.nbLvl += 1
		} else if token == ')' && ctx.mulLvl == 5 && ctx.nbLvl > 0 {
			if doCtx.isDo {
				multSum += ctx.firstNb * ctx.secondNb
			}
			ctx.reset()
		} else {
			ctx.reset()
		}
	}
	return int64(multSum)
}
