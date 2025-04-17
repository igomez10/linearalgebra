package linearalgebra

import "sort"

// TODO
// Implement a compiler for matrix manipulation commands

// ToRowEchelonForm returns a new matrix in row echelon form using Gaussian elimination
// Swap the rows so that all rows with all zero entries are on the bottom
// Swap the rows so that the row with the largest, leftmost nonzero entry is on top.
// Multiply the top row by a scalar so that top row's leading entry becomes 1.
// Add/subtract multiples of the top row to the other rows so that all other entries in the column containing the top row's leading entry are all zero.
// Repeat steps 2-4 for the next leftmost nonzero entry until all the leading entries are 1.
// Swap the rows so that the leading entry of each nonzero row is to the right of the leading entry of the row above it.
func ToRowEchelonForm(matrix [][]float64) [][]float64 {
	// Swap the rows so that all rows with all zero entries are on the bottom
	matrix = SwapRows0sToBottom(matrix)

	// Add/subtract multiples of the top row to the other rows so that all other
	// entries in the column containing the top row's leading entry are all zero.
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// find non 0
			if matrix[i][j] != 0 {
				// make this row pivot row
				matrix = MultiplyRowByScalar(matrix, i, float64(1/matrix[i][j]))

				// turn every column in this pivot to 0
				for z := 0; z < len(matrix); z++ {
					if z == i {
						continue
					}

					if matrix[z][j] != 0 {
						tmp := matrix[z][j]
						matrix = MultiplyRowByScalar(matrix, i, -matrix[z][j])
						matrix = AddRowToRow(matrix, matrix[i], z)
						matrix = MultiplyRowByScalar(matrix, i, float64(1/-tmp))
					}
				}
				break
			}
		}
	}

	// turn current row into pivot row by multiplying by the inverse of the leading entry
	// make every entry in the column of the leading entry 0
	// find next pivot and do the same
	matrix = SwapLargetsLeftmostNonzeroEntry(matrix)

	return matrix
}

func SwapLargetsLeftmostNonzeroEntry(matrix [][]float64) [][]float64 {
	sort.Slice(matrix, func(i, j int) bool {
		for z := 0; z < len(matrix[i]); z++ {
			itemI := matrix[i][z]
			itemJ := matrix[j][z]

			if itemI != itemJ {
				return itemI > itemJ
			}
		}
		return true
	})

	return matrix
}

func GenerateIdentityMatrix(n int) [][]float64 {
	if n < 0 {
		panic("illegal operation")
	}

	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}

	for i := range n {
		matrix[i][i] = 1
	}

	return matrix
}

func SwapRows0sToBottom(matrix [][]float64) [][]float64 {
	sort.Slice(matrix, func(i, j int) bool {
		for x := range matrix[i] {
			if matrix[i][x] == 0 && matrix[j][x] == 0 {
				continue
			}

			if matrix[i][x] != 0 && matrix[j][x] == 0 {
				return true
			}
		}

		return false
	})

	return matrix
}

func SwapRows(matrix [][]float64, i, j int) [][]float64 {
	if i > len(matrix) || j > len(matrix) || i < 0 || j < 0 {
		panic("invalid change")
	}

	tmp := matrix[i]
	matrix[i] = matrix[j]
	matrix[j] = tmp

	return matrix
}

func AddMatrices(matrixA, matrixB [][]float64) [][]float64 {
	if len(matrixA) != len(matrixB) {
		panic("illegal operation")
	}

	response := make([][]float64, len(matrixA))
	for r := range response {
		response[r] = make([]float64, len(matrixA[0]))
	}

	for i := range matrixA {
		for j := range matrixA[i] {
			response[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}

	return response
}

func MultiplyRowByScalar(matrix [][]float64, rowIndex int, scalar float64) [][]float64 {
	if rowIndex < 0 || rowIndex >= len(matrix) {
		panic("invalid change")
	}

	for i := range matrix[rowIndex] {
		matrix[rowIndex][i] *= scalar
	}

	return matrix
}

func MultiplyVectorByScalar(vector []float64, scalar float64) []float64 {
	for i := range vector {
		vector[i] *= scalar
	}

	return vector
}

func IsZeroMatrix(matrix [][]float64) bool {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func CanMultiplyMatrices(matrixA, matrixB [][]float64) bool {
	return len(matrixA[0]) == len(matrixB)
}

// miltiply matrices will use dot product to multiply two matrices
func MultiplyMatrices(matrixA, matrixB [][]float64) [][]float64 {
	if !CanMultiplyMatrices(matrixA, matrixB) {
		panic("invalid multiplication")
	}

	newMatrix := make([][]float64, len(matrixA))
	for i := range newMatrix {
		newMatrix[i] = make([]float64, len(matrixB[0]))
	}

	for i := range newMatrix {
		for j := range newMatrix[i] {
			for z := range matrixB {
				newMatrix[i][j] += matrixA[i][z] * matrixB[z][j]
			}
		}
	}

	return newMatrix
}

func MultiplyMatrixByScalar(matrix [][]float64, scalar float64) [][]float64 {
	for i := range matrix {
		matrix = MultiplyRowByScalar(matrix, i, scalar)
	}
	return matrix
}

func AddRowToRow(matrix [][]float64, rowToAdd []float64, rowIndex int) [][]float64 {
	if rowIndex < 0 || rowIndex >= len(matrix) || len(matrix[0]) != len(rowToAdd) {
		panic("invalid change")
	}

	for i := range matrix[rowIndex] {
		matrix[rowIndex][i] += rowToAdd[i]
	}

	return matrix
}

// GetPivotEntries return the list of indexes where the pivot entries are located
func GetPivotEntries(matrix [][]float64) [][]int {
	// traverse row by row, when finds a non 0, move down until finds a 0,
	//  then move right until finds a non 0, exit when columns or rows are over
	answer := [][]int{}
	i := 0
	j := 0
	for i < len(matrix) && j < len(matrix[i]) {
		if matrix[i][j] != 0 {
			answer = append(answer, []int{i, j})
			i++
		}
		j++
	}

	return answer
}

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
