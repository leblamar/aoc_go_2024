package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getTestFilePath(day uint) string {
	relative_path := fmt.Sprintf("src/day%d/day%d_example.txt", day, day)
	absPath, err := filepath.Abs(relative_path)
	if err != nil {
		log.Fatal(err)
	}
	return absPath
}

func getInputFilePath(day uint) string {
	relative_path := fmt.Sprintf("src/day%d/day%d_input.txt", day, day)
	absPath, err := filepath.Abs(relative_path)
	if err != nil {
		log.Fatal(err)
	}
	return absPath
}

func getUrl(day uint) string {
	return fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func getLinesFromDl(day uint) []string {
	path := getInputFilePath(day)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fileUrl := getUrl(day)
		resp, err := http.Get(fileUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		newFile, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, resp.Body)

		if err != nil {
			log.Fatal(err)
		}
	}
	return readLines(path)
}

func GetLines(justATest bool, day uint) (lines []string) {
	if justATest {
		path := getTestFilePath(day)
		lines = readLines(path)
	} else {
		lines = getLinesFromDl(day)
	}

	return
}
