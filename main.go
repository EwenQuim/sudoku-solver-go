package main

import (
	"flag"
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	// Arguments
	var sudokuFile = flag.String(
		"input",
		"data/sudoku.txt",
		"Select a sudoku file from your path.",
	)
	var profiling = flag.Bool(
		"cpu",
		false,
		"Profiling CPU usage",
	)
	flag.Parse()

	if *profiling {
		defer profile.Start().Stop()
	}

	S := FileHandler(*sudokuFile)

	PrettyPrint(S)

	T := Solve(S)

	// Display
	PrettyPrint(T)

}

// PrettyPrint prints matrix
func PrettyPrint(S [9][9]uint8) {
	for i := 0; i < 9; i++ {
		fmt.Println(S[i])
	}
}
