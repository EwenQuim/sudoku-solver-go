package main

func main() {

	sudokus := PlugCLI()

	// multiple sudokus can be given
	for _, sudokuFile := range sudokus {
		PrettyTitle(sudokuFile)

		S := FileHandler(sudokuFile)

		T := Solve(S)

		PrettyPrint(T)
	}

}
