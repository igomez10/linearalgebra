package linearalgebra

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/cmplx"
	"sort"
	"strconv"
)

// TODO
// Implement a compiler for matrix manipulation commands

// ToRowReducedEchelonForm returns a new matrix in row echelon form using Gaussian elimination
// Swap the rows so that all rows with all zero entries are on the bottom
// Swap the rows so that the row with the largest, leftmost nonzero entry is on top.
// Multiply the top row by a scalar so that top row's leading entry becomes 1.
// Add/subtract multiples of the top row to the other rows so that all other entries in the column containing the top row's leading entry are all zero.
// Repeat steps 2-4 for the next leftmost nonzero entry until all the leading entries are 1.
// Swap the rows so that the leading entry of each nonzero row is to the right of the leading entry of the row above it.
func ToRowReducedEchelonForm(pMatrix [][]float64) [][]float64 {
	// Swap the rows so that all rows with all zero entries are on the bottom
	matrix := CopyMatrix(pMatrix)
	matrix = SwapRows0sToBottom(matrix)

	// Add/subtract multiples of the top row to the other rows so that all other
	// entries in the column containing the top row's leading entry are all zero.
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// find non 0
			if !NearlyEqual(matrix[i][j], 0, 3) {
				// make this row pivot row
				matrix = MultiplyRowByScalar(matrix, i, float64(1/matrix[i][j]))

				// turn every column in this pivot to 0
				for z := 0; z < len(matrix); z++ {
					if z == i {
						continue
					}

					if !NearlyEqual(matrix[z][j], 0, 3) {
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

func CopyMatrix(matrix [][]float64) [][]float64 {
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
		eliminationMatrix = DotProduct(changeMatrices[i], eliminationMatrix)
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
	newMatrix := [][]float64{}
	rowsOf0s := [][]float64{}
	rowsOfNon0s := [][]float64{}
	for i := range matrix {
		isAll0s := true
		for j := range matrix[i] {
			if matrix[i][j] != 0 {
				rowsOfNon0s = append(rowsOfNon0s, matrix[i])
				isAll0s = false
				break
			}
		}
		if isAll0s {
			rowsOf0s = append(rowsOf0s, matrix[i])
		}
	}

	newMatrix = append(newMatrix, rowsOfNon0s...)
	newMatrix = append(newMatrix, rowsOf0s...)

	return newMatrix
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

// CanMultiplyMatrices returns true if the
// number of columns in the first matrix is equal to the number of rows in the second matrix
// 2x3 3x2 returns true
// 3x2 3x2 returns false
func CanMultiplyMatrices(matrixA, matrixB [][]float64) bool {
	if len(matrixA) == 0 || len(matrixB) == 0 {
		if len(matrixA) == 0 && len(matrixB) == 0 {
			return true
		}

		return false
	}

	// check number of columns in A is equal to number of rows in B
	return len(matrixA[0]) == len(matrixB)
}

// DotProductVectors will apply dot product to two vectors
// Since dot product only applies if the dimension of vectors is
// axb and bxc, we will assume that the vectors are 1xN and Nx1
// this function will take care of converting the row vector to a column vector
func DotProductVectors(vectorA, vectorB []float64) float64 {
	rowVectorA := [][]float64{vectorA}

	columnVectorB := [][]float64{}
	for i := range vectorB {
		columnVectorB = append(columnVectorB, []float64{vectorB[i]})
	}
	return DotProduct(rowVectorA, columnVectorB)[0][0]
}

// DotProduct multiply matrices will use dot product to multiply two matrices
// For DotProduct with vectors use DotProductVectors instead
func DotProduct(matrixA, matrixB [][]float64) [][]float64 {
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
	copiedMatrix := CopyMatrix(matrix)
	if rowIndex < 0 || rowIndex >= len(copiedMatrix) || len(copiedMatrix[0]) != len(rowToAdd) {
		panic("invalid change")
	}

	for i := range copiedMatrix[rowIndex] {
		copiedMatrix[rowIndex][i] += rowToAdd[i]
	}

	return copiedMatrix
}

// GetPivotEntries return the list of indexes where the pivot entries are located
func GetPivotEntries(matrix [][]float64) [][]int {
	// traverse row by row, when finds a non 0, move down until finds a 0,
	//  then move right until finds a non 0, exit when columns or rows are over
	answer := [][]int{}
	i := 0
	j := 0
	for i < len(matrix) && j < len(matrix[i]) {
		if !NearlyEqual(matrix[i][j], 0, 3) {
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
func NearlyEqual(a, b float64, decimals int) bool {
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
			if !NearlyEqual(matrix[i][j], float64(0), 3) {
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

// GetUnitVector returns the unit vector for a given vector
// A unit vector has a length of 1
// When called on itself it returns the same vector
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
	matrixCopied := CopyMatrix(matrix)
	matrixCopied = ToRowReducedEchelonForm(matrixCopied)

	// if in rref the last row is 0s then we have 0 solutions
	for i := range matrixCopied[len(matrixCopied)-1] {
		if matrixCopied[len(matrixCopied)-1][i] != 0 {
			// found a non 0 component, we have 1 solution
			return 1
		}
	}

	return 0
}

// find number of pivots, count them and return them
// for R2 we return 2
// for R3 we return 3
// for Rn we return n
// technically the span is all the possible linear combinations of the vectors
// but returning this in code is not possible
func GetMatrixSpan(matrix [][]float64) int {
	copiedMatrix := CopyMatrix(matrix)
	copiedMatrix = ToRowReducedEchelonForm(copiedMatrix)

	// count number of pivots

	counter := 0
	for i := range copiedMatrix {
		for j := range copiedMatrix[i] {
			if copiedMatrix[i][j] != 0 {
				counter++
				break
			}
		}
	}

	return counter
}

// verify if vectors are linearly independant by checking if we get the
// identity matrix when doing gaussian elimination
func areVectorsLinearlyIndependentByGaussianElimination(vectors [][]float64) bool {
	if len(vectors) == 0 {
		return true
	}

	if len(vectors) != len(vectors[0]) {
		return false
	}

	cols := len(vectors[0])
	for i := range vectors {
		if len(vectors[i]) != cols {
			return false
		}
	}

	rref := ToRowReducedEchelonForm(vectors)
	isIdentity := areMatricesEqual(rref, GenerateIdentityMatrix(len(rref)))
	return isIdentity
}

// verify if vectors are linearly independant by checking if
// when doing dot product of two vectors, we get that the absolute
// value of the dot product is **equal** to the multiplication of both vectors length
func areVectorsLinearlyIndependentByCauchySchwarz(vectorA, vectorB []float64) bool {
	if len(vectorA) == 0 && len(vectorB) == 0 {
		return true
	}

	resDotProduct := DotProductVectors(vectorA, vectorB)

	if NearlyEqual(math.Abs(resDotProduct), GetVectorLength(vectorA)*GetVectorLength(vectorB), 3) {
		return false
	}

	return true
}

// areMatricesEqual will check if two matrices are equal by
// checking if they have the same dimensions and if all their
// components are equal
func areMatricesEqual(matrixA, matrixB [][]float64) bool {
	if len(matrixA) != len(matrixB) {
		return false
	}

	if len(matrixA) == 0 {
		return true
	}

	cols := len(matrixA[0])
	if cols != len(matrixB[0]) {
		return false
	}

	// ensure matrix A has same number of cols in all rows
	for i := range matrixA {
		if len(matrixA[i]) != cols {
			return false
		}
	}

	for i := range matrixB {
		if len(matrixB[i]) != cols {
			return false
		}
	}

	for i := range matrixA {
		for j := range matrixA[i] {
			if !NearlyEqual(matrixA[i][j], matrixB[i][j], 3) {
				return false
			}
		}
	}

	return true
}

// verify if vectors are linearly independant by vector triangular inequality
// if  || u + v || == ||u|| + ||v|| then  u v linearly dependent
func areVectorsLinearlyIndependentByTriangularInequality(vectorA, vectorB []float64) bool {
	if len(vectorA) == 0 && len(vectorB) == 0 {
		return true
	}

	summedVector := AddRowToRow([][]float64{vectorA}, vectorB, 0)[0]
	summedVectorLength := GetVectorLength(summedVector)
	addedLength := GetVectorLength(vectorA) + GetVectorLength(vectorB)
	if NearlyEqual(summedVectorLength, addedLength, 3) {
		return false
	}

	return true
}

func SaveMatrix(matrix [][]float64, out io.Writer) error {
	for _, row := range matrix {
		for j, value := range row {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%v", value)
		}
		// if i < len(matrix)-1 {
		fmt.Fprintf(out, " ")
		fmt.Fprintln(out)
		// }
	}
	return nil
}

func LoadMatrix(input io.Reader) ([][]float64, error) {
	reader := bufio.NewReader(input)
	matrix := [][]float64{}
	reachedEnd := false
	for !reachedEnd {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			switch err {
			case io.EOF:
				reachedEnd = true
			default:
				return [][]float64{}, err
			}
		}

		row := []float64{}
		low := 0
		for high := 0; high < len(line); high++ {
			if line[high] == ' ' || high == len(line) {
				component := line[low:high]
				number, err := strconv.ParseFloat(string(component), 64)
				if err != nil {
					return [][]float64{}, err
				}
				row = append(row, number)
				low = high + 1
			}
		}

		if len(row) > 0 {
			matrix = append(matrix, row)
		}
	}

	return matrix, nil
}

// GetAngleBetweenVectors returns the angle between two vectors by
// using the following formula
// dotProduct(vectorA  vectorB) = length(vectorA) * length(vectorB)  * Cos(angle)
func GetAngleBetweenVectors(vectorA, vectorB []float64) float64 {
	resDotProduct := DotProductVectors(vectorA, vectorB)
	resLength := (GetVectorLength(vectorA) * GetVectorLength(vectorB))

	angleInRadians := math.Acos(resDotProduct / resLength)
	return RadiansToDegrees(angleInRadians)
}

func RadiansToDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

// IsUnitVector returns true if the vector length is 1
// otherwise it will return false
func IsUnitVector(vector []float64) bool {
	if GetVectorLength(vector) == 1 {
		return true
	}

	return false
}

// AreVectorsOrthogonal returns true if all vectors are orthogonal
// Orthogonal vectors are vectors that are perpendicular to each other
// The dot product of two orthogonal vectors is 0
// basically DotProduct(vectorA, vectorB) == 0
func AreVectorsOrthogonal(vectors ...[]float64) bool {
	for i := range vectors {
		for j := i; j < len(vectors); j++ {
			if i == j {
				continue
			}

			if DotProductVectors(vectors[i], vectors[j]) != 0 {
				return false
			}
		}
	}

	return true
}

// GetMinor returns the minor matrix, used in getting the determinant
// i j is used as the row and column to exlude, basically the row and column
// that will be excluded
func GetMinor(matrix [][]float64, i, j int) [][]float64 {
	if i > len(matrix) || j > len(matrix[i]) {
		panic("invalid i j ")
	}

	// traverse matrix, create rows and columns with fields that are
	minor := [][]float64{}
	for ci := range matrix {
		if ci == i {
			continue
		}

		newRow := []float64{}
		for cj := range matrix[ci] {
			if cj == j {
				continue
			}

			newRow = append(newRow, matrix[ci][cj])
		}

		minor = append(minor, newRow)
	}

	return minor
}

// GetDeterminant of a given matrix
func GetDeterminant(matrix [][]float64) float64 {
	if !IsMatrixSquare(matrix) {
		panic("cannot calculate determinant of non square matrix")
	}

	// case matrix is empty
	if len(matrix) == 0 {
		return 0
	}

	// case matrix has len 1
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	// case matrix has len 2
	if len(matrix) == 2 {
		var det float64
		det += matrix[0][0] * matrix[1][1]
		det -= matrix[0][1] * matrix[1][0]
		return det
	}

	// case matrix has len 3
	if len(matrix) == 3 {
		var det float64
		det += matrix[0][0] * (matrix[1][1]*matrix[2][2] - matrix[1][2]*matrix[2][1])
		det -= matrix[0][1] * (matrix[1][0]*matrix[2][2] - matrix[1][2]*matrix[2][0])
		det += matrix[0][2] * (matrix[1][0]*matrix[2][1] - matrix[1][1]*matrix[2][0])

		return det
	}

	// generic case n>3 with laplace expansion
	// we expand one row only
	var det float64
	for col := range matrix[0] {
		minor := GetMinor(matrix, 0, col)
		cofactor := math.Pow(-1, float64(col)) * matrix[0][col] * GetDeterminant(minor)
		det += cofactor
	}

	return det
}

func IsMatrixSquare(matrix [][]float64) bool {
	if len(matrix) == 0 {
		return true
	}

	expectedLen := len(matrix)
	for i := range matrix {
		if len(matrix[i]) != expectedLen {
			return false
		}
	}

	return true
}

// IsMatrixInvertible checks if determinant is non 0
// If the matrix has an inverse then it is invertible
func IsMatrixInvertible(matrix [][]float64) bool {
	if !IsMatrixSquare(matrix) {
		return false
	}

	determinant := GetDeterminant(matrix)
	if determinant == 0 {
		return false
	}

	return true
}

// GetMatrixRank returns the number of pivots
// the rank of a matrix is the dimension of the span (all possible linear combinations of the vectors)
// It is equal to the number of linearly independent rows or columns in the matrix.
// The rank can be found by converting the matrix to row echelon form and counting the number of non-zero rows.
func GetMatrixRank(matrix [][]float64) int {
	reduced := ToRowReducedEchelonForm(matrix)
	pivots := GetPivotEntries(reduced)
	return len(pivots)
}

// The dimension of the column space is equal to the number of linearly independent columns
func GetColumnSpaceDimension(matrix [][]float64) int {
	return GetMatrixRank(matrix)
}

// GetColumnSpace returns the column space of a matrix
// this is basically all the pivot columns (from the rref) but in the original matrix
func GetColumnSpace(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return [][]float64{}
	}

	if len(matrix[0]) == 0 {
		return [][]float64{}
	}

	// get the row reduced echelon form of the matrix
	rref := ToRowReducedEchelonForm(matrix)

	// get the pivots entries
	pivots := GetPivotEntries(rref)

	// add all the pivot columns to the column space
	columnSpace := [][]float64{}
	for i := range pivots {
		pivotColumn := GetColumn(matrix, pivots[i][1])
		columnSpace = AppendMatrix(columnSpace, pivotColumn)
	}

	return columnSpace
}

// The dimension of the row space is equal to the number of linearly independent rows
func GetRowSpaceDimension(matrix [][]float64) int {
	return GetMatrixRank(matrix)
}

// GetCofactorMatrix returns the cofactor matrix for a given matrix.
// Each element in the cofactor matrix is the determinant of the minor matrix
// multiplied by (-1) raised to the power of the sum of its row and column indices.
// The cofactor matrix is used to compute the adjugate (adjoint) matrix and the inverse of a matrix.
// Individual cofactors (not necessarily the entire matrix) are also used in Laplace expansion to compute the determinant.
func GetCofactorMatrix(matrix [][]float64) [][]float64 {
	if !IsMatrixSquare(matrix) {
		panic("cannot calculate cofactor of non square matrix")
	}

	if len(matrix) == 0 {
		return [][]float64{}
	}
	if len(matrix) == 1 {
		return [][]float64{{1}}
	}

	cofactorMatrix := [][]float64{}
	for i := range matrix {
		newRow := []float64{}
		for j := range matrix[i] {
			minor := GetMinor(matrix, i, j)
			cof := math.Pow(-1, float64(i+1+j+1)) * GetDeterminant(minor)
			newRow = append(newRow, cof)
		}

		cofactorMatrix = append(cofactorMatrix, newRow)
	}
	return cofactorMatrix
}

// TransposeMatrix returns the transpose of a given matrix
// The transpose of a matrix is obtained by swapping the rows and columns in a way that
// the first row becomes the first column, the second row becomes the second column, and so on
func TransposeMatrix[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 {
		return matrix
	}
	newmatrix := [][]T{}
	for col := 0; col < len(matrix[0]); col++ {
		newRow := []T{}
		for row := 0; row < len(matrix); row++ {
			newRow = append(newRow, matrix[row][col])
		}
		newmatrix = append(newmatrix, newRow)
	}
	return newmatrix
}

func RowToColumnVector[T any](vector []T) [][]T {
	if len(vector) == 0 {
		return [][]T{}
	}

	transposed := [][]T{}
	for i := range vector {
		transposed = append(transposed, []T{vector[i]})
	}

	return transposed
}

// GetAdjugateMatrix returns the adjugate matrix of a given matrix
// The adjugate matrix is the transpose of the cofactor matrix.
// It is used in the calculation of the inverse of a matrix.
func GetAdjugateMatrix(matrix [][]float64) [][]float64 {
	cofactorMatrix := GetCofactorMatrix(matrix)
	adjugateMatrix := TransposeMatrix(cofactorMatrix)
	return adjugateMatrix
}

// GetMatrixNillity returns the number of rows that are not pivot entries
// in a matrix nxn
// n = pivot + nullity
func GetMatrixNullity(matrix [][]float64) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix[0]) == 0 {
		return 0
	}

	return len(matrix[0]) - GetMatrixRank(matrix)
}

// matrix is invertible if there exists an inverse A^-1
// A^-1 = 1/det(A) * adj(A)
func GetInverseMatrixByDeterminant(matrix [][]float64) [][]float64 {
	if !IsMatrixSquare(matrix) {
		panic("cannot calculate inverse of non square matrix")
	}

	det := GetDeterminant(matrix)
	if det == 0 {
		panic("cannot calculate inverse of non invertible matrix")
	}

	adjMatrix := GetAdjugateMatrix(matrix)

	res := MultiplyMatrixByScalar(adjMatrix, 1/det)
	return res
}

// CrossProduct returns a vector that is orthogonal to both A and B
// we calculate it witha determinant, something like a matrix
/*
       	| i   j  k |
   axb=	| a1 a2 a3 |
       	| b1 b2 b3 |


a = [1,0,2]
b = [-2,1,0]

	 i j k
	1  0 2
	-2 1 0


*/
func CrossProduct(vectorA, vectorB []float64) []float64 {
	if len(vectorA) != 3 || len(vectorB) != 3 {
		panic("cross product is only defined in 3 dimensions")
	}

	res := []float64{
		vectorA[1]*vectorB[2] - vectorA[2]*vectorB[1],
		-(vectorA[0]*vectorB[2] - vectorA[2]*vectorB[0]),
		vectorA[0]*vectorB[1] - vectorA[1]*vectorB[0],
	}

	return res
}

// IsVectorInTheNullSpaceOfMatrix checks that
// dotProduct(A, v) = 0
// where 0 is a vector of 0s
func IsVectorInTheNullSpaceOfMatrix(vector []float64, matrix [][]float64) bool {
	// lets assume we do matrix vector multiplication
	// A*v
	// assuming A is mxn
	// v should be mx1
	columnVector := TransposeMatrix([][]float64{vector})
	if !CanMultiplyMatrices(matrix, columnVector) {
		panic("cannot mulitply this matrix with this vector")
	}

	// check if we should transpose vector to make it column or row vector
	resultVector := DotProduct(matrix, columnVector)
	if len(resultVector[0]) != 1 {
		panic("result vector should be a row vector")
	}

	for i := range resultVector {
		if !NearlyEqual(resultVector[i][0], 0, 3) {
			return false
		}
	}

	return true
}

// GetNullSpaceOfMatrix returns the null space of a matrix
// The null space of a matrix is the set of all vectors v such that A*v = 0
// It is the solution set of the homogeneous system of linear equations represented by the matrix A.
// The null space can be found by solving the system of equations represented by the matrix
// this is also the kernel of the matrix
func GetNullSpaceOfMatrix(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return [][]float64{}
	}

	if len(matrix[0]) == 0 {
		return [][]float64{}
	}

	// get the row reduced echelon form of the matrix
	rref := ToRowReducedEchelonForm(matrix)

	// get the pivots entries
	pivots := GetPivotEntries(rref)

	// create a list of vectors that are in the null space
	nullSpace := [][]float64{}

	for currentColumnIndex := 0; currentColumnIndex < len(rref[0]); currentColumnIndex++ {
		// if this column is a pivot column, we skip it
		isPivotColumn := false
		for _, currentPivot := range pivots {
			if currentPivot[1] == currentColumnIndex {
				isPivotColumn = true
				break
			}
		}

		if isPivotColumn {
			continue
		}

		// create a nullSpaceVector that has 0s in all pivot columns and 1 in this column
		nullSpaceVector := make([]float64, len(rref[0]))
		nullSpaceVector[currentColumnIndex] = 1

		for _, currentPivot := range pivots {
			nullSpaceVector[currentPivot[1]] = -rref[currentPivot[0]][currentColumnIndex]
		}

		nullSpace = append(nullSpace, nullSpaceVector)
	}

	return nullSpace
}

func GetColumn(matrix [][]float64, columnIndex int) [][]float64 {
	if len(matrix) == 0 || columnIndex < 0 || columnIndex >= len(matrix[0]) {
		panic("invalid column index")
	}

	transposed := TransposeMatrix(matrix)[columnIndex]
	res := TransposeMatrix([][]float64{transposed})
	return res
}

func AppendMatrix(matrixA [][]float64, matrixB [][]float64) [][]float64 {
	if len(matrixA) != len(matrixB) && len(matrixA) > 0 && len(matrixB) > 0 {
		panic("cannot append columns of matrices with different number of rows")
	}

	if len(matrixA) == 0 {
		return matrixB
	}

	if len(matrixB) == 0 {
		return matrixA
	}

	newmatrix := [][]float64{}
	for i := range matrixA {
		newmatrix = append(newmatrix, []float64{})
		newmatrix[i] = append(newmatrix[i], matrixA[i]...)
		newmatrix[i] = append(newmatrix[i], matrixB[i]...)
	}

	return newmatrix
}

func HadamardProduct(matrixA, matrixB [][]float64) [][]float64 {
	if !CanMultiplyMatrices(matrixA, matrixB) {
		panic("cannot calculate Hadamard product of matrices with different dimensions")
	}

	if len(matrixA) == 0 || len(matrixB) == 0 {
		return [][]float64{}
	}

	if len(matrixA) != len(matrixB) || len(matrixA[0]) != len(matrixB[0]) {
		panic("cannot calculate Hadamard product of matrices with different dimensions")
	}

	result := make([][]float64, len(matrixA))
	for i := range result {
		result[i] = make([]float64, len(matrixA[i]))
	}

	for i := range matrixA {
		for j := range matrixA[i] {
			result[i][j] = matrixA[i][j] * matrixB[i][j]
		}
	}

	return result
}

// GetEigenvalues returns the eigenvalues of a square matrix
// For 2x2 matrices we use the quadratic formula
// For larger matrices we use the QR algorithm to approximate the eigenvalues
// The eigenvalues are returned as a slice of complex128 to account for complex eigenvalues
func GetEigenvalues(matrix [][]float64) []complex128 {
	if !IsMatrixSquare(matrix) {
		panic("cannot calculate eigenvalues of non square matrix")
	}

	if len(matrix) == 0 {
		return []complex128{}
	}

	if len(matrix) == 1 {
		return []complex128{complex(matrix[0][0], 0)}
	}

	// For 2x2 matrix we can use the quadratic formula
	if len(matrix) == 2 {
		a := 1.0
		b := -(matrix[0][0] + matrix[1][1])
		c := matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]

		discriminant := b*b - 4*a*c
		if discriminant >= 0 {
			root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
			root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
			return []complex128{complex(root1, 0), complex(root2, 0)}
		} else {
			realPart := -b / (2 * a)
			imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
			return []complex128{complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)}
		}
	}

	// For larger matrices, we can use the QR algorithm.
	// This simple implementation drives the matrix toward real Schur form
	// (quasi-upper-triangular with 1x1 and 2x2 blocks). We'll then extract
	// eigenvalues from the resulting 1x1/2x2 blocks so complex pairs are handled.
	A := CopyMatrix(matrix)
	n := len(A)
	for iter := 0; iter < 1000; iter++ {
		Q, R := qrDecomposition(A)
		A = DotProduct(R, Q)
	}

	// Zero-out tiny subdiagonal elements for robust block detection
	eps := 1e-9
	for i := 0; i < n-1; i++ {
		if math.Abs(A[i+1][i]) < eps {
			A[i+1][i] = 0
		}
	}

	// Extract eigenvalues from 1x1 and 2x2 blocks
	eigenvalues := make([]complex128, 0, n)
	i := 0
	for i < n {
		if i < n-1 && math.Abs(A[i+1][i]) > 0 { // 2x2 block
			a, b := A[i][i], A[i][i+1]
			c, d := A[i+1][i], A[i+1][i+1]
			// Characteristic: λ^2 - (a+d)λ + (ad - bc) = 0
			tr := a + d
			det := a*d - b*c
			disc := tr*tr - 4*det
			if disc >= 0 {
				sqrtDisc := math.Sqrt(disc)
				lambda1 := 0.5 * (tr + sqrtDisc)
				lambda2 := 0.5 * (tr - sqrtDisc)
				eigenvalues = append(eigenvalues, complex(lambda1, 0), complex(lambda2, 0))
			} else {
				realPart := 0.5 * tr
				imagPart := 0.5 * math.Sqrt(-disc)
				eigenvalues = append(eigenvalues, complex(realPart, imagPart), complex(realPart, -imagPart))
			}
			i += 2
		} else { // 1x1 block
			eigenvalues = append(eigenvalues, complex(A[i][i], 0))
			i++
		}
	}

	return eigenvalues
}

// qrDecomposition performs QR decomposition of matrix A
// using the Gram-Schmidt process
// A = QR where Q is orthogonal and R is upper triangular
func qrDecomposition(A [][]float64) ([][]float64, [][]float64) {
	m := len(A)
	n := len(A[0])

	Q := make([][]float64, m)
	for i := range Q {
		Q[i] = make([]float64, n)
	}
	R := make([][]float64, n)
	for i := range R {
		R[i] = make([]float64, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			Q[i][j] = A[i][j]
		}
		for k := 0; k < j; k++ {
			var dot float64
			for i := 0; i < m; i++ {
				dot += Q[i][k] * A[i][j]
			}
			for i := 0; i < m; i++ {
				Q[i][j] -= dot * Q[i][k]
			}
		}
		var norm float64
		for i := 0; i < m; i++ {
			norm += Q[i][j] * Q[i][j]
		}
		norm = math.Sqrt(norm)
		for i := 0; i < m; i++ {
			Q[i][j] /= norm
		}
		for k := 0; k <= j; k++ {
			var dot float64
			for i := 0; i < m; i++ {
				dot += Q[i][k] * A[i][j]
			}
			R[k][j] = dot
		}
	}

	return Q, R
}

func GetEigenvectors(matrix [][]float64) [][]complex128 {
	if !IsMatrixSquare(matrix) {
		panic("cannot calculate eigenvectors of non square matrix")
	}

	if len(matrix) == 0 {
		return [][]complex128{}
	}

	n := len(matrix)
	eigenvalues := GetEigenvalues(matrix)

	// Cache eigenspaces to handle repeated eigenvalues
	// Key: eigenvalue (as complex), Value: list of eigenvectors for that eigenvalue
	eigenspaceCache := make(map[complex128][][]complex128)
	vectorIndexCache := make(map[complex128]int)

	eigenvectors := make([][]complex128, len(eigenvalues))

	for idx, lambda := range eigenvalues {
		// Check if we've already computed the eigenspace for this eigenvalue
		eigenspace, cached := eigenspaceCache[lambda]
		if !cached {
			eigenspace = computeEigenspace(matrix, lambda, n)
			eigenspaceCache[lambda] = eigenspace
			vectorIndexCache[lambda] = 0
		}

		// Get the next eigenvector from this eigenspace (or zero vector if exhausted)
		vectorIdx := vectorIndexCache[lambda]
		if vectorIdx < len(eigenspace) {
			eigenvectors[idx] = eigenspace[vectorIdx]
			vectorIndexCache[lambda]++
		} else {
			// Use zero vector if we've exhausted the eigenspace (defective matrix)
			eigenvectors[idx] = make([]complex128, n)
		}
	}

	return eigenvectors
}

// computeEigenspace finds all linearly independent eigenvectors for a given eigenvalue.
// For real eigenvalues, it solves (A - λI)v = 0 and returns the null space basis.
// For complex eigenvalues, it finds a single eigenvector.
func computeEigenspace(matrix [][]float64, lambda complex128, n int) [][]complex128 {
	if imag(lambda) == 0 {
		return computeRealEigenspace(matrix, real(lambda), n)
	}
	return computeComplexEigenspace(matrix, lambda, n)
}

// computeRealEigenspace solves (A - λI)v = 0 for real eigenvalue λ
func computeRealEigenspace(matrix [][]float64, lambda float64, n int) [][]complex128 {
	// Form (A - λI)
	AminusLambdaI := CopyMatrix(matrix)
	for i := 0; i < n; i++ {
		AminusLambdaI[i][i] -= lambda
	}

	// Find null space vectors
	nullSpace := GetNullSpaceOfMatrix(AminusLambdaI)

	// Convert null space vectors to complex and normalize
	eigenspace := make([][]complex128, len(nullSpace))
	for i, vec := range nullSpace {
		eigenvector := make([]complex128, n)
		for j := 0; j < n; j++ {
			eigenvector[j] = complex(vec[j], 0)
		}
		normalizeEigenvector(eigenvector)
		eigenspace[i] = eigenvector
	}

	return eigenspace
}

// computeComplexEigenspace solves (A - λI)v = 0 for complex eigenvalue λ
func computeComplexEigenspace(matrix [][]float64, lambda complex128, n int) [][]complex128 {
	// Form (A - λI) in complex arithmetic
	AminusLambdaI := make([][]complex128, n)
	for i := 0; i < n; i++ {
		AminusLambdaI[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			AminusLambdaI[i][j] = complex(matrix[i][j], 0)
		}
		AminusLambdaI[i][i] -= lambda
	}

	// Find an eigenvector
	eigenvector := solveComplexHomogeneousSystem(AminusLambdaI)
	normalizeEigenvector(eigenvector)

	return [][]complex128{eigenvector}
}

// normalizeEigenvector scales a complex vector to unit length and applies a consistent sign convention.
// The sign is chosen so that the first non-zero element is positive (real part positive, or
// if real part is zero, imaginary part positive).
func normalizeEigenvector(v []complex128) {
	// Calculate the norm (magnitude)
	normSquared := complex(0, 0)
	for i := range v {
		normSquared += v[i] * cmplx.Conj(v[i])
	}
	norm := math.Sqrt(real(normSquared))

	// Skip normalization for zero vectors
	if norm <= 1e-10 {
		return
	}

	// Normalize to unit length
	normComplex := complex(norm, 0)
	for i := range v {
		v[i] /= normComplex
	}

	// Apply sign convention: flip all elements if first non-zero element is "negative"
	for i := range v {
		if cmplx.Abs(v[i]) > 1e-10 {
			// For complex numbers, we use the convention:
			// - If real part is negative, flip the sign
			// - If real part is zero and imaginary part is negative, flip the sign
			if real(v[i]) < -1e-10 || (math.Abs(real(v[i])) < 1e-10 && imag(v[i]) < -1e-10) {
				for j := range v {
					v[j] = -v[j]
				}
			}
			break
		}
	}
}

// solveComplexHomogeneousSystem finds a non-trivial solution to Ax = 0 for complex matrix A
func solveComplexHomogeneousSystem(A [][]complex128) []complex128 {
	n := len(A)
	if n == 0 {
		return []complex128{}
	}

	// Create a copy for Gaussian elimination
	mat := make([][]complex128, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]complex128, n)
		copy(mat[i], A[i])
	}

	// Gaussian elimination to find rank
	rank := 0
	pivotCols := make([]int, 0, n)
	for col := 0; col < n && rank < n; col++ {
		// Find pivot
		pivot := rank
		for i := rank + 1; i < n; i++ {
			if cmplx.Abs(mat[i][col]) > cmplx.Abs(mat[pivot][col]) {
				pivot = i
			}
		}

		if cmplx.Abs(mat[pivot][col]) < 1e-10 {
			continue
		}

		// Swap rows
		mat[rank], mat[pivot] = mat[pivot], mat[rank]

		// Eliminate column
		for i := rank + 1; i < n; i++ {
			if cmplx.Abs(mat[i][col]) > 1e-10 {
				factor := mat[i][col] / mat[rank][col]
				for j := col; j < n; j++ {
					mat[i][j] -= factor * mat[rank][j]
				}
			}
		}
		pivotCols = append(pivotCols, col)
		rank++
	}

	// For a full-rank matrix, the only solution to Ax=0 is the zero vector
	// This shouldn't happen for (A - λI) when λ is a true eigenvalue,
	// but we handle it for completeness
	if rank == n {
		return make([]complex128, n) // Return zero vector
	}

	// Find the first free column (not a pivot column) and set it to 1
	eigenvector := make([]complex128, n)
	pivotSet := make(map[int]bool, rank)
	for _, pc := range pivotCols {
		pivotSet[pc] = true
	}
	freeCol := -1
	for c := 0; c < n; c++ {
		if !pivotSet[c] {
			freeCol = c
			break
		}
	}
	if freeCol >= 0 {
		eigenvector[freeCol] = complex(1, 0)
	}

	// Back-substitute using tracked pivot columns
	for i := rank - 1; i >= 0; i-- {
		pc := pivotCols[i]
		sum := complex(0, 0)
		for j := pc + 1; j < n; j++ {
			sum += mat[i][j] * eigenvector[j]
		}
		if cmplx.Abs(mat[i][pc]) > 1e-10 {
			eigenvector[pc] = -sum / mat[i][pc]
		}
	}

	return eigenvector
}

// complexToReal converts a matrix of complex eigenvectors
// to a matrix of real values by taking the real part of each component
func complexToReal(matrix [][]complex128) [][]float64 {
	res := make([][]float64, len(matrix))
	for i := range res {
		res[i] = make([]float64, len(matrix[i]))
		for j := range res[i] {
			res[i][j] = real(matrix[i][j])
		}
	}

	return res
}

func complexToRealVector(vector []complex128) []float64 {
	res := make([]float64, len(vector))
	for i := range res {
		res[i] = real(vector[i])
	}

	return res
}

type Matrix struct {
	data [][]float64
}

func (m *Matrix) Transpose() {
	m.data = TransposeMatrix(m.data)
}

func (m *Matrix) SetIndex(i, j int, value float64) {
	if i < 0 || i >= len(m.data) || j < 0 || j >= len(m.data[0]) {
		panic("index out of bounds")
	}

	m.data[i][j] = value
}

func (m *Matrix) Copy() *Matrix {
	newmatrix := &Matrix{data: CopyMatrix(m.data)}
	return newmatrix
}

func (m *Matrix) DotProduct(matrixB Matrix) Matrix {
	res := DotProduct(m.data, matrixB.data)
	return Matrix{data: res}
}

func (m *Matrix) GetColumn(columnIndex int) []float64 {
	if len(m.data) == 0 {
		return []float64{}
	}
	column := make([]float64, len(m.data))
	for i := range m.data {
		column[i] = m.data[i][columnIndex]
	}
	return column
}

func (m *Matrix) Append(other Matrix) Matrix {
	if len(m.data) == 0 {
		return other
	}
	if len(other.data) == 0 {
		return *m
	}
	appendedData := make([][]float64, len(m.data))
	for i := range m.data {
		appendedData[i] = append(m.data[i], other.data[i]...)
	}
	return Matrix{data: appendedData}
}

func (m *Matrix) GetRow(rowIndex int) []float64 {
	if len(m.data) == 0 {
		return []float64{}
	}
	return m.data[rowIndex]
}

// validateEigenDecomposition will validate if
// A * eigenvector = eivenValue * eigenvector
// A * v = lambda * v
func validateEigenDecomposition(matrix Matrix, eigenvalue float64, eigenvector []complex128) bool {
	realEigen := complexToRealVector(eigenvector)

	realEigenMatrix := Matrix{data: [][]float64{realEigen}}
	realEigenMatrix.Transpose()
	// A * v
	Av := DotProduct(matrix.data, realEigenMatrix.data)
	// lambda * v
	lambdaV := MultiplyVectorByScalar(realEigenMatrix.GetColumn(0), eigenvalue)
	lambdaVRes := TransposeMatrix([][]float64{lambdaV})
	return areMatricesEqual(Av, lambdaVRes)
}

type SVDResult struct {
	U Matrix
	S Matrix
	V Matrix
}

// SVD performs Singular Value Decomposition on a matrix A
// It returns matrices U, S, and V such that A = U * S * V^T
func SVD(m Matrix) SVDResult {
	// compute AtA to find eigenvalues, we need a square matrix
	AtA := m.DotProduct(m)
	eigenValues := GetEigenvalues(AtA.data)
	singularValues := make([]float64, len(eigenValues))
	for i, val := range eigenValues {
		if real(val) < 0 {
			singularValues[i] = 0
		} else {
			singularValues[i] = math.Sqrt(real(val))
		}
	}

	// create diagonal matrix S with singular values
	diagonalScaling := make([][]float64, len(singularValues))
	// fill with 0s
	for i := range diagonalScaling {
		diagonalScaling[i] = make([]float64, len(singularValues))
	}
	for i := range singularValues {
		diagonalScaling[i][i] = singularValues[i]
	}

	// compute V by finding the eigenvectors of AtA
	vectors := GetEigenvectors(AtA.data)
	V := make([][]float64, len(vectors))
	for i := range vectors {
		V[i] = make([]float64, len(vectors))
		for j := range vectors {
			V[i][j] = real(vectors[i][j])
		}
	}

	// compute U as A * V * S^-1
	SInv := make([][]float64, len(diagonalScaling))
	for i := range diagonalScaling {
		SInv[i] = make([]float64, len(diagonalScaling))
		for j := range diagonalScaling {
			if diagonalScaling[i][j] != 0 {
				SInv[i][j] = 1 / diagonalScaling[i][j]
			} else {
				SInv[i][j] = 0
			}
		}
	}

	AV := DotProduct(m.data, V)
	U := DotProduct(AV, SInv)

	svd := SVDResult{
		U: Matrix{data: U},
		S: Matrix{data: diagonalScaling},
		V: Matrix{data: V},
	}

	return svd
}
