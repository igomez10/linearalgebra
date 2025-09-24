# Linear Algebra (Go)

Utilities for working with matrices and vectors in Go. This module provides core linear algebra routines (row-reduced echelon form, dot products, identity matrices, scalar operations, basic validation helpers) and a small demo app under `cmd/graph` that renders simple 2D vectors to an image.

## Project structure

- `go.mod` — Go module `github.com/igomez10/linearalgebra`
- `main.go` — Library package `linearalgebra` with matrix/vector helpers
- `main_test.go` — Unit tests for the library
- `cmd/graph/` — Demo app that draws vectors and saves `3dplot.png`
  - `main.go` — Render a simple grid and a few vectors
  - `main_test.go` — Tests for rendering helpers
- `python/main_unitest.py` — A small Python unittest example (unrelated to the Go module; useful for teaching/tests)

## Requirements

- Go 1.23+

## Install

Add the module to your project:

```bash
go get github.com/igomez10/linearalgebra@latest
```

Then import it in your code:

```go
import "github.com/igomez10/linearalgebra"
```

## Quick start (library)

### 1) Row-Reduced Echelon Form (RREF)

```go
package main

import (
    "fmt"
    la "github.com/igomez10/linearalgebra"
)

func main() {
    A := [][]float64{{1, 2, 3}, {2, 5, 3}, {1, 0, 8}}
    rref := la.ToRowReducedEchelonForm(A)
    fmt.Println("RREF:", rref)
}
```

### 2) Matrix multiplication (dot product)

```go
A := [][]float64{{1, 2}, {3, 4}} // 2x2
B := [][]float64{{5, 6}, {7, 8}} // 2x2
C := linearalgebra.DotProduct(A, B)
// C = [[19, 22], [43, 50]]
```

### 3) Identity matrix and scalar ops

```go
I := linearalgebra.GenerateIdentityMatrix(3) // 3x3 identity
J := linearalgebra.MultiplyMatrixByScalar(I, 2) // every entry *2
```

### 4) Vector dot product

```go
u := []float64{1, 2, 3}
v := []float64{4, 5, 6}
dp := linearalgebra.DotProductVectors(u, v) // 1*4 + 2*5 + 3*6 = 32
```

Common helpers you’ll find in the package include:

- `ToRowReducedEchelonForm(matrix [][]float64) [][]float64`
- `DotProduct(A, B [][]float64) [][]float64`
- `DotProductVectors(a, b []float64) float64`
- `GenerateIdentityMatrix(n int) [][]float64`
- `AddMatrices(A, B [][]float64) [][]float64`
- `MultiplyMatrixByScalar(M [][]float64, s float64) [][]float64`
- `MultiplyRowByScalar(M [][]float64, row int, s float64) [][]float64`
- `AddRowToRow(M [][]float64, row []float64, idx int) [][]float64`
- `CanMultiplyMatrices(A, B [][]float64) bool`
- `IsZeroMatrix(M [][]float64) bool`

Note: Some functions may panic on illegal operations (e.g., dimension mismatch); validate inputs with helpers like `CanMultiplyMatrices` first.

## Demo app: draw vectors to an image

There’s a small program under `cmd/graph` that renders a grid and a few 2D vectors, saving the result as `3dplot.png`.

Run it:

```bash
go run ./cmd/graph
```

You should see a `3dplot.png` file created in `cmd/graph/` with axis lines and example vectors.

## Testing

Run all Go tests:

```bash
go test ./...
```

The demo also has its own tests:

```bash
go test ./cmd/graph
```

Optional Python example (requires NumPy):

```bash
python3 python/main_unitest.py
```

## Notes and roadmap

- The library focuses on clarity and teaching; some routines are not optimized for large matrices.
- Future ideas: determinants, inverses, LU/QR decompositions, eigenvalues/eigenvectors, and improved numerical stability.

