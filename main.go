package main

import (
	"fmt"
	"time"

	"pkg.amethysts.studio/sudoku-solver-go/solver"
)

func main() {
	sudokus := PlugCLI()

	// multiple sudokus can be given
	for _, sudokuFile := range sudokus {
		PrettyTitle(sudokuFile)

		S := FileHandler(sudokuFile)

		start := time.Now()
		T := solver.Solve(S)
		fmt.Println("Solved in", time.Since(start))

		PrettyPrint(T)
	}
}
