package main

import (
	"aoc_go_2024/src/day1"
	"aoc_go_2024/src/day2"
	"aoc_go_2024/src/day3"
	"aoc_go_2024/src/day4"
	"errors"
)

func GetDay(day uint) (IDay[any], error) {
	switch day {
	case 1:
		return any(day1.Day1{}).(IDay[any]), nil
	case 2:
		return any(day2.Day2{}).(IDay[any]), nil
	case 3:
		return any(day3.Day3{}).(IDay[any]), nil
	case 4:
		return any(day4.Day4{}).(IDay[any]), nil
	default:
		err := errors.New("there is no such day yet")
		return nil, err
	}
}
