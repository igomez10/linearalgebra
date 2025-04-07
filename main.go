package linearalgebra

// TODO
// GetPivot entries
// ToRowEchelonForm returns a new matrix in row echelon form using Gaussian elimination
// Implement a compiler for matrix manipulation commands

func IsReducedRowEchelonForm(matrix [][]float64) bool {
	// all pivots are equal to 1
	if !IsRowEchelonForm(matrix) {
		return false
	}

	if !allPivotsEqualTo1(matrix) {
		return false
	}

	// all entries in the base column (column with pivot) are equal to 0 except for the pivot itself
	if !allEntriesInBaseColumnAre0(matrix) {
		return false
	}

	return true
}

func allPivotsEqualTo1(matrix [][]float64) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				if matrix[i][j] != 1 {
					return false
				}
				break
			}
		}
	}

	return true
}

func allEntriesInBaseColumnAre0(matrix [][]float64) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				// check other entries in same column are not not non 0
				for z := 0; z < len(matrix); z++ {
					if z == i {
						continue
					}

					current := matrix[z][j]
					if current != 0 {
						return false
					}
				}
			}
		}
	}

	return true
}

// IsRowEchelonForm Checks if the matrix is in row echelon form
func IsRowEchelonForm(matrix [][]float64) bool {
	// All 0 rows are at the bottom
	if !allZeroRowsAreAtBottom(matrix) {
		return false
	}

	// All pivot entries are to the right of the pivot entry in the row above
	if !allPivotEntriesAreRightOfPivotbove(matrix) {
		return false
	}

	return true
}

func allZeroRowsAreAtBottom(matrix [][]float64) bool {
	foundRowsOfOnly0s := false
	for _, row := range matrix {
		rowIsOnly0s := true
		for _, entry := range row {
			if entry != 0 {

				if foundRowsOfOnly0s {
					return false
				}

				rowIsOnly0s = false
			}
		}

		if rowIsOnly0s {
			foundRowsOfOnly0s = true
		}
	}

	return true
}

func allPivotEntriesAreRightOfPivotbove(matrix [][]float64) bool {
	// traverse matrix keeping track of currentPivot, if finds a pivot where column is smaller
	// than current pivot return false, else return true
	currentPivot := []int{-1, -1}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				if j <= currentPivot[1] {
					return false
				}

				currentPivot = []int{i, j}
				break
			}
		}
	}

	return true
}
