package day9

import (
	"aoc_go_2024/src/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

type day9 struct {
	input []int
}

func parse(lines []string) day9 {
	if len(lines) > 1 && len(lines) == 0 {
		log.Fatal("Bad input")
	}

	input := make([]int, 0, len(lines[0]))
	for _, car := range lines[0] {
		input = append(input, int(car-'0'))
	}

	return day9{input}
}

func (pg *day9) day9_1(debug bool) {
	count := 0
	for _, elem := range pg.input {
		count += elem
	}

	disk := make([]int16, 0, count)
	isFree := false
	ID := 0
	for _, elem := range pg.input {
		var valToAppend int16
		if isFree {
			valToAppend = -1
		} else {
			valToAppend = int16(ID)
		}
		for i := 0; i < elem; i++ {
			disk = append(disk, valToAppend)
		}
		isFree = !isFree
		if !isFree {
			ID++
		}
	}
	if debug {
		fmt.Println("Disk:", disk)
	}

	inc, dec := 0, len(disk)-1
	for inc < dec {
		if disk[inc] != -1 {
			inc++
			continue
		} else if disk[dec] == -1 {
			dec--
			continue
		}

		disk[inc] = disk[dec]
		disk[dec] = -1
		inc++
		dec--
	}
	if debug {
		fmt.Println("Disk:", disk)
	}

	sum := 0
	for i, id := range disk {
		if id == -1 {
			break
		}
		if debug {
			fmt.Println("i * id:", i, "*", id, " = ", i*int(id))
		}
		sum += i * int(id)
	}
	fmt.Println("Part1:", sum)
}

type values struct {
	size int
	id   int
}

func (pg *day9) day9_2(debug bool) {
	fillList := make([]int, 0, len(pg.input))
	emptyList := make([]int, 0, len(pg.input))
	for i, val := range pg.input {
		if i%2 == 0 {
			fillList = append(fillList, val)
		} else {
			emptyList = append(emptyList, val)
		}
	}

	if len(pg.input)%2 == 1 {
		emptyList = append(emptyList, 0)
	}

	fillToPosList := make([]int, 0, len(fillList))
	for i := 0; i < len(fillList); i++ {
		fillToPosList = append(fillToPosList, i)
	}

	posToFillList := make([]int, 0, len(fillList))
	for i := 0; i < len(fillList); i++ {
		posToFillList = append(posToFillList, i)
	}

	for fi := len(fillList) - 1; fi >= 0; fi-- {
		fSize := fillList[fi]
		ei := 0
		eSize := emptyList[ei]
		for ei <= fi && eSize < fSize {
			ei++
			if ei < len(emptyList) {
				eSize = emptyList[ei]
			}
		}

		if ei > fi || eSize < fSize {
			continue
		}

		// countValid = 1
		// fSize = 3
		// fi = -2
		// ei = 2
		// 0, 1, 3, 3, 1, 1, 1, 1, 1, 2
		// 0, 1, 0, 0, 3, 1, 1, 1, 5, 2
		// 0, 1, 	 3  , 3, 1, 1, 1, 1,  1, 2
		// 0, 1, +0, 3-3, 3, 1, 1, 1, 1+3+1, 2
		// Or
		// countValid = 2
		// fSize = 2
		// fi = -5
		// ei = 4
		// 0, 1, 0, 0, 3, 1, 1, 1, 5, 2
		// 0, 1, 0, 0, 0, 1, 4, 1, 5, 2
		// 0, 1, 0, 0,     3  , 1,  1, 1, 5, 2
		// 0, 1, 0, 0, +0, 3-2, 1+2+1, 1, 5, 2
		// Or
		// countValid = 3
		// fSize = 1
		// fi = -7
		// ei = 1
		// 0, 1, 0, 0, 0, 1, 4, 1, 5, 2
		// 0, 0, 0, 0, 1, 1, 4, 1, 5, 2
		// 0,     1  , 0, 0,  0, 1, 4, 1, 5, 2
		// 0, +0, 1-1, 0, 0+1+0, 1, 4, 1, 5, 2
		// Or fSize = 1
		// fi: 0, 1, 2
		// fi: 0, 1, 2
		// ei: 1, 0, 0
		// ei: 0, 1, 0
		// ei: 1, 0, 0
		// ei: 0, 1, 0
		// ei: 1-1, 0+1+0, 0
		// ei: 	 0, 1,   0

		fiSrc := fillToPosList[fi]

		// Delete fill element so must merge close empty elements
		eiMerge1 := fiSrc
		eiMerge2 := fiSrc - 1
		eSizeMerge := emptyList[eiMerge1] + emptyList[eiMerge2] + fSize
		//emptyList = append(emptyList[:eiMerge1], emptyList[eiMerge1+1:]...)
		for i := eiMerge2; i < len(emptyList)-1; i++ {
			emptyList[i] = emptyList[i+1]
		}
		emptyList[eiMerge2] = eSizeMerge

		// reduce empty element of fill element size
		emptyList[ei] = eSize - fSize

		// add 0 element before new fill
		for i := len(emptyList) - 1; i > ei; i-- {
			emptyList[i] = emptyList[i-1]
		}
		emptyList[ei] = 0

		// countValid = 1
		// fi = 7
		// ei = 2
		// {s:2,i:0}, {s:2,i:9}, {s:3,i:1}, {s:1,i:2}, {s:3,i:3}, {s:2,i:4}, {s:4,i:5}, {s:4,i:6}, {s:3,i:7}, {s:4,i:8}
		// {s:2,i:0}, {s:2,i:9}, {s:3,i:1}, {s:3,i:7}, {s:1,i:2}, {s:3,i:3}, {s:2,i:4}, {s:4,i:5}, {s:4,i:6}, {s:4,i:8}
		// fillToPosList : [0, 2, 3, 4, 5, 6, 7, 8, 9, 1], fiSrc = 8
		// fillToPosList : [0, 2, 4, 5, 6, 7, 8, 3, 9, 1], fiDst = 3
		// posToFillList : [0, 9, 1, 2, 3, 4, 5, 6, 7, 8]
		// posToFillList : [0, 9, 1, 7, 2, 3, 4, 5, 6, 8]
		// Or
		// countValid = 2
		// fi = -5
		// ei = 4
		// {s:2,i:0}, {s:2,i:9}, {s:3,i:1}, {s:3,i:7}, {s:1,i:2}, {s:3,i:3}, {s:2,i:4}, {s:4,i:5}, {s:4,i:6}, {s:4,i:8}
		// {s:2,i:0}, {s:2,i:9}, {s:3,i:1}, {s:3,i:7}, {s:1,i:2}, {s:2,i:4}, {s:3,i:3}, {s:4,i:5}, {s:4,i:6}, {s:4,i:8}
		fiDst := ei + 1

		for i := fiSrc; i > fiDst; i-- {
			curFillToIncrPos := posToFillList[i-1]
			fillToPosList[curFillToIncrPos] += 1
			posToFillList[i] = curFillToIncrPos
		}
		posToFillList[fiDst] = fi
		fillToPosList[fi] = fiDst
	}

	fmt.Println(emptyList)
	fmt.Println(posToFillList)
	fmt.Println(fillList)
	sum := 0
	currentIndex := 0
	printRes := ""
	for i, fill := range posToFillList {
		size := fillList[fill]
		finalIndex := currentIndex + size
		for currentIndex < finalIndex {
			sum += currentIndex * fill
			currentIndex++
			printRes += strconv.Itoa(fill)
		}

		currentIndex += emptyList[i]
		for j := 0; j < emptyList[i]; j++ {
			printRes += "."
		}
	}

	fmt.Println("Part2:", sum)
	fmt.Println(printRes)
}

func Day9(isTest, debug bool) {
	fmt.Println("Welcome to day 9!!!")

	lines := utils.GetLines(isTest, 9)
	day9 := parse(lines)

	start := time.Now()
	day9.day9_1(debug)
	ellapsed := time.Since(start)

	fmt.Println("Time ellapsed : ", ellapsed)

	start = time.Now()
	day9.day9_2(debug)
	ellapsed = time.Since(start)
	fmt.Println("Time ellapsed : ", ellapsed)
}