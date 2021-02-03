package main

import (
	"flag"
	"fmt"

	"github.com/pkg/profile"
)

func main() {

	var profiling = flag.Bool(
		"cpu",
		false,
		"Profiling CPU usage",
	)
	flag.Parse()

	flag.Args()

	if *profiling {
		defer profile.Start().Stop()
	}

	// multiple sudokus can be given
	for _, sudokuFile := range flag.Args() {
		fmt.Println("\nFile:", sudokuFile)

		S := FileHandler(sudokuFile)

		T := Solve(S)

		PrettyPrint(T)
	}

}

// PrettyPrint prints matrix
func PrettyPrint(S [9][9]uint8) {
	for i := 0; i < 9; i++ {
		fmt.Println(S[i])
	}
}
