package day14

import (
	"aoc_go_2024/src/utils"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type position = utils.Position
type speed = utils.Position
type Day14 struct {
	IsTest bool
}

type robot struct {
	p position
	s speed
}

type matrix struct {
	height int
	width  int
	robots []robot
}

func parseTwoVal(toParse string) (int, int) {
	s := strings.Split(toParse, "=")
	if len(s) != 2 {
		log.Fatal("This should never happened")
	}

	s = strings.Split(s[1], ",")
	if len(s) != 2 {
		log.Fatal("This should never happened")
	}

	val1, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal("This should never happened")
	}
	val2, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal("This should never happened")
	}
	return val1, val2
}

func parseRobot(line string) robot {
	s := strings.Split(line, " ")
	if len(s) != 2 {
		log.Fatal("This should never happened")
	}

	p1, p2 := parseTwoVal(s[0])
	s1, s2 := parseTwoVal(s[1])

	return robot{position{X: p1, Y: p2}, speed{X: s1, Y: s2}}
}

func (d Day14) Parse(lines []string) matrix {
	var height int
	var width int
	if d.IsTest {
		width = 11
		height = 7
	} else {
		width = 101
		height = 103
	}

	robots := make([]robot, 0, len(lines))
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}

	return matrix{height, width, robots}
}

func (r robot) move(nbSecond, height, width int) robot {
	newPX := (((r.p.X + nbSecond*r.s.X) % width) + width) % width
	newPY := (((r.p.Y + nbSecond*r.s.Y) % height) + height) % height

	return robot{position{X: newPX, Y: newPY}, r.s}
}

func (r robot) isPresentIn(startH, endH, startW, endW int) bool {
	if r.p.X < startW {
		return false
	} else if r.p.X >= endW {
		return false
	} else if r.p.Y < startH {
		return false
	} else if r.p.Y >= endH {
		return false
	} else {
		return true
	}
}

func countPresent(robots []robot, startH, endH, startW, endW int) (count int64) {
	count = 0
	for _, r := range robots {
		if r.isPresentIn(startH, endH, startW, endW) {
			count += 1
		}
	}
	return
}

func (m matrix) String() string {
	mapPos := make(map[position]int)
	for _, r := range m.robots {
		mapPos[r.p] += 1
	}

	var b bytes.Buffer
	for j := 0; j < m.height; j++ {
		for i := 0; i < m.width; i++ {
			if count, ok := mapPos[position{X: i, Y: j}]; ok {
				b.WriteString(strconv.Itoa(count))
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (d Day14) Part1(debug bool, m matrix) int64 {
	fmt.Println(m)
	newRobots := make([]robot, 0, len(m.robots))
	for _, r := range m.robots {
		newR := r.move(100, m.height, m.width)
		newRobots = append(newRobots, newR)
	}

	middleH := m.height / 2
	middleW := m.width / 2

	newMatrix := matrix{m.height, m.width, newRobots}
	fmt.Println(newMatrix)

	countTopLeft := countPresent(newRobots, 0, middleH, 0, middleW)
	fmt.Println("Top left:", countTopLeft)
	countTopRight := countPresent(newRobots, 0, middleH, middleW+1, m.width)
	fmt.Println("Top right:", countTopRight)
	countBottomLeft := countPresent(newRobots, middleH+1, m.height, 0, middleW)
	fmt.Println("Bottom left:", countBottomLeft)
	countBottomRight := countPresent(newRobots, middleH+1, m.height, middleW+1, m.width)
	fmt.Println("Bottom right:", countBottomRight)

	return countTopLeft * countTopRight * countBottomLeft * countBottomRight
}

func (d Day14) Part2(debug bool, m matrix) int64 {
	fmt.Println(m)
	nbSeconds := 1000
	curM := m
	for i := 0; i < nbSeconds; i++ {
		time.Sleep(100 * time.Millisecond)
		newRobots := make([]robot, 0, len(curM.robots))
		for _, r := range curM.robots {
			newR := r.move(1, curM.height, curM.width)
			newRobots = append(newRobots, newR)
		}
		curM = matrix{curM.height, curM.width, newRobots}
		fmt.Println(curM)
	}
	return 0
}
