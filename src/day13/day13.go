package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type button struct {
	x decimal.Decimal
	y decimal.Decimal
}
type price struct {
	x decimal.Decimal
	y decimal.Decimal
}
type machine struct {
	a button
	b button
	p price
}
type game []machine

type Day13 struct{}

func parseExp(line string, sep string) int64 {
	s := strings.Split(line, sep)
	if len(s) != 2 {
		log.Fatal("This should never happend")
	}

	val, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func parseButton(line string) button {
	s := strings.Split(line, ": ")
	if len(s) != 2 {
		log.Fatal("This should never happend")
	}

	s = strings.Split(line, ", ")
	if len(s) != 2 {
		log.Fatal("This should never happend")
	}

	x := parseExp(s[0], "+")
	y := parseExp(s[1], "+")
	return button{decimal.NewFromInt(x), decimal.NewFromInt(y)}
}

func parsePrice(line string) price {
	s := strings.Split(line, ": ")
	if len(s) != 2 {
		log.Fatal("This should never happend")
	}

	s = strings.Split(line, ", ")
	if len(s) != 2 {
		log.Fatal("This should never happend")
	}

	x := parseExp(s[0], "=")
	y := parseExp(s[1], "=")
	return price{decimal.NewFromInt(x), decimal.NewFromInt(y)}

}

func (d Day13) Parse(lines []string) game {
	mList := make(game, 0, 4*len(lines))

	var m machine
	for i, line := range lines {
		switch i % 4 {
		case 0:
			m = machine{}
			m.a = parseButton(line)
		case 1:
			m.b = parseButton(line)
		case 2:
			m.p = parsePrice(line)
			mList = append(mList, m)
		case 3:
		}
	}

	return mList
}

func (m machine) getNb() decimal.Decimal {
	return m.p.y.Sub(m.p.x.Mul(m.a.y).Div(m.a.x)).Div(m.b.y.Sub(m.b.x.Mul(m.a.y).Div(m.a.x)))
}

func (m machine) getNa() decimal.Decimal {
	nB := m.getNb()
	return m.p.x.Sub(nB.Mul(m.b.x)).Div(m.a.x)
}

func epsilonEqual(d1, d2, e decimal.Decimal) bool {
	return d1.Sub(d2).Abs().LessThanOrEqual(e)
}

var curEpsilon = decimal.NewFromFloat(1e-4)

func (m machine) getMin(debug bool) (int64, int64, bool) {
	nA := m.getNa()
	nB := m.getNb()
	roundedNA := nA.Round(0)
	roundedNB := nB.Round(0)

	if !epsilonEqual(nA, roundedNA, curEpsilon) {
		if debug {
			fmt.Println("nA is not an integer :", nA)
		}
		return -1, -1, false
	} else if !epsilonEqual(nB, roundedNB, curEpsilon) {
		if debug {
			fmt.Println("nB is not an integer :", nB)
		}
		return -1, -1, false
	}

	return roundedNA.IntPart(), roundedNB.IntPart(), true
}

func (d Day13) Part1(debug bool, input game) (sum int64) {
	sum = 0
	for i, m := range input {
		if debug {
			fmt.Println("Process machine", i+1, "/", len(input))
		}
		nA, nB, ok := m.getMin(debug)
		if !ok {
			continue
		} else if nA > 100 || nB > 100 {
			continue
		} else {
			sum += 3*nA + nB
		}
	}
	return
}

var toAdd = decimal.NewFromInt(10000000000000)

func (m machine) copy() machine {
	newM := machine{}
	newM.a = m.a
	newM.b = m.b
	newM.p = price{m.p.x.Add(toAdd), m.p.y.Add(toAdd)}
	return newM
}

func (input game) transformInput() game {
	tInput := make(game, 0, len(input))

	for _, m := range input {
		tInput = append(tInput, m.copy())
	}

	return tInput
}

func (d Day13) Part2(debug bool, input game) (sum int64) {
	tInput := input.transformInput()
	sum = 0
	for i, m := range tInput {
		if debug {
			fmt.Println("Process machine", i+1, "/", len(tInput))
		}
		nA, nB, ok := m.getMin(debug)
		if !ok {
			continue
		} else {
			sum += 3*nA + nB
		}
	}
	return
}
