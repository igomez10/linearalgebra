package linearalgebra

import (
	"math"
	"sort"
)

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

func copyMatrix(matrix [][]float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))

	for i := range matrix {
		newMatrix[i] = make([]float64, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}

	return newMatrix
}

func GetEliminationMatrix(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]float64{}
	}
	// Swap the rows so that all rows with all zero entries are on the bottom
	matrix = SwapRows0sToBottom(matrix)
	changeMatrices := [][][]float64{}

	// Add/subtract multiples of the top row to the other rows so that all other
	// entries in the column containing the top row's leading entry are all zero.
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// find non 0
			if matrix[i][j] != 0 {
				// make this row pivot row
				tmp := matrix[i][j]
				matrix = MultiplyRowByScalar(matrix, i, float64(1/tmp))
				changeMatrix := MultiplyRowByScalar(GenerateIdentityMatrix(len(matrix)), i, float64(1/tmp))
				changeMatrices = append(changeMatrices, changeMatrix)

				// turn every column in this pivot to 0
				for z := 0; z < len(matrix); z++ {
					if z == i {
						continue
					}

					if matrix[z][j] != 0 {
						tmp := matrix[z][j]

						matrix = MultiplyRowByScalar(matrix, i, -tmp)
						changeMatrix = MultiplyRowByScalar(GenerateIdentityMatrix(len(matrix)), i, -tmp)

						matrix = AddRowToRow(matrix, matrix[i], z)
						changeMatrix = AddRowToRow(changeMatrix, changeMatrix[i], z)

						matrix = MultiplyRowByScalar(matrix, i, float64(1/-tmp))
						changeMatrix = MultiplyRowByScalar(changeMatrix, i, float64(1/-tmp))

						changeMatrices = append(changeMatrices, changeMatrix)
					}
				}
				break
			}
		}
	}

	// turn current row into pivot row by multiplying by the inverse of the leading entry
	// make every entry in the column of the leading entry 0
	// find next pivot and do the same
	_ = SwapLargetsLeftmostNonzeroEntry(matrix)
	// eliminationMatrix = SwapLargetsLeftmostNonzeroEntry(eliminationMatrix)

	eliminationMatrix := changeMatrices[0]
	for i := 1; i < len(changeMatrices); i++ {
		eliminationMatrix = MultiplyMatrices(changeMatrices[i], eliminationMatrix)
	}

	return eliminationMatrix
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
	if len(matrixA) == 0 || len(matrixB) == 0 {
		if len(matrixA) == 0 && len(matrixB) == 0 {
			return true
		}

		return false
	}

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
				newMatrix[i][j] += float64(matrixA[i][z]) * float64(matrixB[z][j])
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

// used to compare floats
func nearlyEqual(a, b float64, decimals int) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	allowedError := 1 / math.Pow(10, float64(decimals))
	if diff > allowedError {
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
			if !nearlyEqual(matrix[i][j], float64(0), 3) {
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

func GetVectorLength(vector []float64) float64 {
	var powed float64 = 0
	for i := range vector {
		powed += math.Pow(vector[i], 2)
	}

	return math.Sqrt(powed)
}

func GetUnitVector(vector []float64) []float64 {
	res := make([]float64, len(vector))
	copy(res, vector)
	vectorLength := GetVectorLength(vector)

	for i := range vector {
		res[i] = float64(vector[i]) / float64(vectorLength)
	}

	return res
}

// Get number of solutions, 0, 1
// We will assume that result column is not in the matrix
// therefore it is not possible to identify a matrix with
// infinite solutions
func GetNumberOfSolutions(matrix [][]float64) float64 {
	matrixCopied := copyMatrix(matrix)
	matrixCopied = ToRowEchelonForm(matrixCopied)

	// if in rref the last row is 0s then we have 0 solutions
	for i := range matrixCopied[len(matrixCopied)-1] {
		if matrixCopied[len(matrixCopied)-1][i] != 0 {
			// found a non 0 component, we have 1 solution
			return 1
		}
	}

	return 0
}
