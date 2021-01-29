package main

import (
	"fmt"
)

func main() {

	S := FileHandler("data/sudoku_special_against_computers")

	prettyPrint(S)

	T := Solve(S)

	// Display
	prettyPrint(T)

}

func prettyPrint(S [9][9]int) {
	for i := 0; i < 9; i++ {
		fmt.Println(S[i])
	}
}
