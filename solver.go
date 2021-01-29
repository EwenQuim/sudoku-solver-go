package main

import "math"

//is ok
func listeInit(S [9][9]int) []int {
	l := []int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if S[i][j] != 0 {
				// fmt.Println(i, j, S[i][j])
				l = append(l, int(10*i+j))
			}
		}
	}

	return l
}

// is ok
func nbPossible(S [9][9]int, i int, j int) int {
	for _, v := range listeInit(S) {

		if v == (10*i + j) {
			return 0
		}
	}

	c := 0
	for n := 1; n <= 9; n++ {
		if isAvailable(S, i, j, n) {
			c++
		}
	}

	return c
}

//is ok
func tableauPossibilities(S [9][9]int) [9][9]int {
	tab := [9][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			tab[i][j] = nbPossible(S, i, j)
		}
	}
	return tab
}

func tableauOrdre(S [9][9]int) []int {
	tab := tableauPossibilities(S)
	liste := []int{}
	for k := 1; k < 10; k++ {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if tab[i][j] == k {
					liste = append(liste, 10*i+j)
				}
			}
		}
	}
	return liste
}

////////////////////////////
//////   Résolution   //////
////////////////////////////

// not really block : only checks the 4 squares that can't be reached with col/row checking
func isAvailableInBloc(S [9][9]int, i int, j int, n int) bool {
	for k := 3 * floor3(i); k < 3*floor3(i)+3; k++ {
		if k != i {
			for l := 3 * floor3(j); l < 3*floor3(j)+3; l++ {
				// fmt.Println("coords", k, l, "=", S[k][l])
				if l != j && S[k][l] == n {
					return false
				}
			}
		}
	}
	return true
}

//is ok
func isAvailable(S [9][9]int, i int, j int, n int) bool {
	for k := 0; k < 9; k++ {
		if (S[i][k] == n && k != j) || (S[k][j] == n && k != i) {
			return false
		}
	}
	return isAvailableInBloc(S, i, j, n)
}

//is ok
func assigneValeur(S [9][9]int, i int, j int, mini int) ([9][9]int, int) {
	for n := mini; n < 10; n++ {
		if isAvailable(S, i, j, n) {
			S[i][j] = n
			return S, n
		}
	}
	S[i][j] = 0
	return S, 0
}

func floor3(n int) int {
	return int(math.Floor(float64(n) / 3.0))
}

// Solve a sudoku
func Solve(S [9][9]int) [9][9]int {
	// Timer
	defer Track(Runningtime("solve"))
	tabOrdre := tableauOrdre(S)
	nToChange := len(tabOrdre)
	tabMini := S
	rank := 0

	T, replacedValue := S, 0

	for rank < nToChange {

		// Computing
		n := tabOrdre[rank]
		i := n / 10
		j := n % 10

		T, replacedValue = assigneValeur(T, i, j, tabMini[i][j]+1)

		if replacedValue != 0 {
			rank++
		} else {
			rank--
		}

		tabMini[i][j] = replacedValue
	}

	return T
}
