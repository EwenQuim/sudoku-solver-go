package solver

// Board represents the sudoku board. uint8 allows a 2x faster processing (and we always have 0 <= 9 <= 255)
type Board = [9][9]uint8

func digitsPossible(S Board, i uint8, j uint8) []uint8 {
	if S[i][j] != 0 { // cell already set
		return []uint8{}
	}

	digits := []uint8{}
	for n := uint8(1); n <= 9; n++ {
		if isAvailable(S, i, j, n) {
			digits = append(digits, n)
		}
	}

	return digits
}

func matrixPossibilities(S Board) [9][9][]uint8 {
	tab := [9][9][]uint8{}

	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			tab[i][j] = digitsPossible(S, i, j)
		}
	}
	return tab
}

func tableauOrder(S Board) []uint8 {
	tab := matrixPossibilities(S)
	liste := []uint8{}
	for k := int(1); k < 10; k++ {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if len(tab[i][j]) == k {
					liste = append(liste, uint8(10*i+j))
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
func isAvailableInBloc(S Board, i uint8, j uint8, n uint8) bool {
	for k := floor3(i); k < floor3(i)+3; k++ {
		if k != i {
			for l := floor3(j); l < floor3(j)+3; l++ {
				if l != j && S[k][l] == n {
					return false
				}
			}
		}
	}
	return true
}

func isAvailableInLine(S Board, i uint8, j uint8, n uint8) bool {
	for k := uint8(0); k < 9; k++ {
		if (S[i][k] == n && k != j) || (S[k][j] == n && k != i) {
			return false
		}
	}
	return true
}

func isAvailable(S Board, i uint8, j uint8, n uint8) bool {
	return isAvailableInLine(S, i, j, n) && isAvailableInBloc(S, i, j, n)
}

func floor3(n uint8) uint8 {
	return 3 * (n / 3)
}

// Solve solves a sudoku and returns the answer
func Solve(S Board) Board {
	// defer Track(Runningtime("solving"))

	// Initialise possibilities, order and digit position
	possibilities := matrixPossibilities(S)
	sliceOrder := tableauOrder(S)
	maxDigitToFind := uint8(len(sliceOrder))
	var indexOfCurrentDigitFor [9][9]int

	// variables
	var rank uint8

	for rank < maxDigitToFind {

		// Which cell must be processed ?
		n := sliceOrder[rank]
		i := n / 10
		j := n % 10

		if indexOfCurrentDigitFor[i][j] < len(possibilities[i][j]) {
			// There is a digit in the list of possibilities to put here
			client := possibilities[i][j][indexOfCurrentDigitFor[i][j]]

			if isAvailable(S, i, j, client) {
				// digit available -> go forward to next cell
				S[i][j] = client
				rank++
			} else {
				// digit already in line, col or 3x3 cell -> try higher digit for same cell
				indexOfCurrentDigitFor[i][j]++
			}
		} else {
			// there is no digit possible for this cell

			// first 'reset' the cell state
			S[i][j] = 0
			indexOfCurrentDigitFor[i][j] = 0

			// then go back to previous cell
			rank--
			n = sliceOrder[rank]
			i = n / 10
			j = n % 10
			// previous digit was available but we found it wasn't okay, so increase it
			indexOfCurrentDigitFor[i][j]++
		}

	}

	return S
}
