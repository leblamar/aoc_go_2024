package dayUtils

import (
	"log"
	"os"
	"strconv"
)

type Args struct {
	Day    int
	IsTest bool
	Debug  bool
	OnlyP1 bool
	OnlyP2 bool
}

func GetDefaultArgs() Args {
	return Args{-1, false, false, false, false}
}

func GetArgs() (retArgs Args) {
	retArgs = GetDefaultArgs()

	args := os.Args
	if len(args) <= 1 {
		return
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	retArgs.Day = day

	if len(args) >= 3 {
		for _, arg := range args {
			if arg == "-t" {
				retArgs.IsTest = true
			} else if arg == "-d" {
				retArgs.Debug = true
			} else if arg == "-p1" {
				retArgs.OnlyP1 = true
			} else if arg == "-p2" {
				retArgs.OnlyP2 = true
			}
		}
	}

	return
}
