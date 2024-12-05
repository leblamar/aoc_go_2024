package day5

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type rule struct {
	X int
	Y int
}

type update []int

type program struct {
	updates []update
	rules   []rule
}

func (up update) getMid() int {
	return up[len(up)/2]
}

func parseRule(line string) rule {
	vals := strings.Split(line, "|")
	if len(vals) != 2 {
		log.Fatal("Not good rule size")
	}

	X, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal(err)
	}

	Y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal(err)
	}

	return rule{X, Y}
}

func parseUpdate(line string) update {
	vals := strings.Split(line, ",")
	up := make(update, 0, len(vals))

	for _, val := range vals {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		up = append(up, intVal)
	}

	return up
}

func parse(lines []string) program {
	rules := make([]rule, 0)
	updates := make([]update, 0)

	beforeEnter := true
	for _, line := range lines {
		if line == "" {
			beforeEnter = false
			continue
		}

		if beforeEnter {
			rules = append(rules, parseRule(line))
		} else {
			updates = append(updates, parseUpdate(line))
		}
	}

	return program{updates, rules}
}

func (r rule) check(a, b int) bool {
	return !(r.X == b && r.Y == a)
}

func (up update) checkRule(r rule) bool {
	for i := 0; i < len(up); i++ {
		for j := i + 1; j < len(up); j++ {
			if !r.check(up[i], up[j]) {
				return false
			}
		}
	}

	return true
}

func (up update) isCorrectOrder(rules []rule) bool {
	for _, r := range rules {
		if !up.checkRule(r) {
			return false
		}
	}

	return true
}

func (pg program) day5_1() {
	count := 0
	for _, up := range pg.updates {
		if up.isCorrectOrder(pg.rules) {
			count += up.getMid()
		}
	}

	fmt.Println("Part 1:", count)
}

func (up *update) sort(r rule) bool {
	hasSwap := false
	for i := 0; i < len(*up); i++ {
		for j := i + 1; j < len(*up); j++ {
			A := (*up)[i]
			B := (*up)[j]
			if !r.check(A, B) {
				hasSwap = true
				(*up)[i] = B
				(*up)[j] = A
			}
		}
	}
	return hasSwap
}

func (up update) correctOrder(rules []rule) update {
	newUp := make(update, len(up))
	copy(newUp, up)

	hasSwap := true
	for hasSwap {
		hasSwap = false
		for _, rule := range rules {
			hasSwap = hasSwap || newUp.sort(rule)
		}
	}

	return newUp
}

func (pg program) day5_2() {
	count := 0
	for _, up := range pg.updates {
		if !up.isCorrectOrder(pg.rules) {
			newUp := up.correctOrder(pg.rules)
			count += newUp.getMid()
		}
	}

	fmt.Println("Part 2:", count)
}

func Day5(justATest bool) {
	fmt.Println("Welcome to day 5!!!")

	lines := utils.GetLines(justATest, 5)
	pg := parse(lines)

	pg.day5_1()
	pg.day5_2()
}
