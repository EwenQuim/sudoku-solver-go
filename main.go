package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

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
