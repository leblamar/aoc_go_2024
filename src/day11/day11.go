package day11

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type stone string

func parse(lines []string) []stone {
	if len(lines) > 1 || len(lines) == 0 {
		log.Fatal("This should never happened")
	}

	strStones := strings.Split(lines[0], " ")
	stones := make([]stone, len(strStones))
	for i, strStone := range strStones {
		stones[i] = stone(strStone)
	}

	return stones
}

func (s stone) next() []stone {
	if s == "0" {
		return []stone{"1"}
	} else if len(s)%2 == 0 {
		middle := len(s) / 2
		firstPart := string(s)[:middle]
		secondPart := string(s)[middle:]
		val, err := strconv.ParseInt(secondPart, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		newSecondPart := strconv.FormatInt(val, 10)
		return []stone{stone(firstPart), stone(newSecondPart)}
	} else {
		val, err := strconv.ParseInt(string(s), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		newS := strconv.FormatInt(val*2024, 10)
		return []stone{stone(newS)}
	}
}

func toString(stones []stone) {
	ret := ""
	for _, s := range stones {
		ret += string(s) + " "
	}
	fmt.Println(ret)
}

func day11_1(debug bool, stones []stone) {
	curStones := stones
	for i := 0; i < 25; i++ {
		if debug {
			toString(curStones)
		}
		newStones := make([]stone, 0, len(curStones)*2)
		for _, s := range curStones {
			newStones = append(newStones, s.next()...)
		}
		curStones = newStones
	}
	if debug {
		toString(curStones)
	}
	fmt.Println("Part 1:", len(curStones))
}

func makeCopy(mapStones map[stone]int64) map[stone]int64 {
	newMapStones := make(map[stone]int64, len(mapStones))
	for k, v := range mapStones {
		newMapStones[k] = v
	}
	return newMapStones
}

func computeVal(mapStones map[stone]int64) int64 {
	var sum int64 = 0
	for _, count := range mapStones {
		sum += count
	}
	return sum
}

func day11_2(debug bool, stones []stone) {
	mapStones := make(map[stone]int64, 0)
	for _, s := range stones {
		mapStones[s] += 1
	}

	for i := 0; i < 6; i++ {
		fmt.Println("Current step", i+1, "/75 :", computeVal(mapStones), " with a map of len", len(mapStones))
		if debug {
			fmt.Println(mapStones)
		}
		newMap := make(map[stone]int64, len(mapStones))
		for s, count := range mapStones {
			newStones := s.next()
			for _, newStone := range newStones {
				newMap[newStone] += count
			}
		}
		mapStones = newMap
	}

	if debug {
		fmt.Println(mapStones)
	}
	// 0:2 14168:1 2:4 2024:1 2097446912:1 3:1 4:1 40:2 48:2 6:2 7:1 8:1 80:1 96:1
	// 2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
	// 80:1 96:1
	// 4048 80 96
	fmt.Println("Part 2:", computeVal(mapStones))
}

func Day11(isTest, debug bool) {
	fmt.Println("Welcome to day 11!!!")

	lines := utils.GetLines(isTest, 11)
	stones := parse(lines)

	start := time.Now()
	day11_1(debug, stones)
	ellapsed := time.Since(start)

	fmt.Println("Time ellapsed : ", ellapsed)

	start = time.Now()
	day11_2(debug, stones)
	ellapsed = time.Since(start)
	fmt.Println("Time ellapsed : ", ellapsed)
}
