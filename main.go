package main

import "pkg.amethysts.studio/sudoku-solver-go/solver"

func main() {

	sudokus := PlugCLI()

	// multiple sudokus can be given
	for _, sudokuFile := range sudokus {
		PrettyTitle(sudokuFile)

		S := FileHandler(sudokuFile)

		T := solver.Solve(S)

		PrettyPrint(T)
	}

}
