package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type button struct {
	x int64
	y int64
}
type price struct {
	x int64
	y int64
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
	return button{x, y}
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
	return price{x, y}

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

type sol struct {
	nbA int64
	nbB int64
}

func (s sol) value() int64 {
	return s.nbA*3 + s.nbB
}

func (s1 sol) Compare(s2 sol) int {
	val1 := s1.value()
	val2 := s2.value()

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	} else {
		return 0
	}
}

func (m machine) getPos(nbA, nbB int64) (int64, int64) {
	nbX := nbA*m.a.x + nbB*m.b.x
	nbY := nbA*m.a.y + nbB*m.b.y

	return nbX, nbY
}

func (m machine) solve(debug bool, nbA, nbB int64) []sol {
	if debug && nbA == 80 {
		fmt.Println("Try nbA:", nbA, ", nbB:", nbB)
	}
	if nbA > 100 || nbB > 100 {
		return []sol{}
	}

	nbX, nbY := m.getPos(nbA, nbB)
	if nbX == m.p.x && nbY == m.p.y {
		return []sol{{nbA, nbB}}
	} else if nbX > m.p.x && nbY > m.p.y {
		return []sol{}
	}

	sols1 := m.solve(debug, nbA+1, nbB)
	sols2 := m.solve(debug, nbA, nbB+1)
	return append(sols1, sols2...)
}

func (m machine) getMin(debug bool) int64 {
	sols := m.solve(debug, 0, 0)
	if len(sols) == 0 {
		return 0
	}

	var min int64 = sols[0].value()
	for _, sol := range sols[1:] {
		if sol.value() < min {
			min = sol.value()
		}
	}

	return min
}

func (d Day13) Part1(debug bool, input game) (sum int64) {
	sum = 0
	for i, m := range input {
		if debug {
			fmt.Println("Process machine", i+1, "/", len(input))
		}
		sum += m.getMin(debug)
	}
	return
}

func (d Day13) Part2(debug bool, input game) int64 {
	return 0
}
