package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// FileHandler opens a file by its name and returns string
func FileHandler(fileName string) [9][9]int {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	S := [9][9]int{}
	for i, eachline := range fileTextLines {
		for j, character := range eachline {
			c := string(character)
			if c == "." {
				S[i][j] = 0
			} else {
				S[i][j], err = strconv.Atoi(c)
			}
		}

	}
	return S
}
