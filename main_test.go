package linearalgebra

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"reflect"
	"testing"
)

type TestCases struct {
	// Test case name
	Name string
	// Input matrix
	Matrix [][]float64
	// Expected output matrix
	Expected [][]float64
}

var testcases = []TestCases{
	{
		Name: "Test 1",
		Matrix: [][]float64{
			{1, 0, 4, 11},
			{1, -1, 4, 6},
			{2, 0, 9, 25},
		},
		Expected: [][]float64{},
	},
}

func TestIsrowEchelonForm(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is_row_echelon_form",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 1},
					{0, 1, 0, 1},
					{0, 0, 1, 1},
				},
			},
			want: true,
		},
		{
			name: "is_not_row_echelon_form_because_of_row_3",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 1},
					{0, 1, 0, 1},
					{1, 0, 0, 1},
				},
			},
			want: false,
		},
		{
			name: "is_not_row_echelon_form_because_of_row_2",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 1},
					{1, 1, 0, 1},
					{0, 0, 1, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRowEchelonForm(tt.args.matrix); got != tt.want {
				t.Errorf("IsrowEchelonForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allZeroRowsAreAtBottom(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all_zeroes_are_at_bottom",
			args: args{
				matrix: [][]float64{
					{1, 1},
					{0, 0},
				},
			},
			want: true,
		},
		{
			name: "not_all_zeroes_are_at_bottom",
			args: args{
				matrix: [][]float64{
					{0, 0},
					{1, 1},
				},
			},
			want: false,
		},
		{
			name: "all_zeroes",
			args: args{
				matrix: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			want: true,
		},
		{
			name: "non_zeroes",
			args: args{
				matrix: [][]float64{
					{1, 1},
					{1, 1},
				},
			},
			want: true,
		},
		{
			name: "zero_row_not_in_bottom",
			args: args{
				matrix: [][]float64{
					{1, 1},
					{0, 0},
					{1, 1},
				},
			},
			want: false,
		},
		{
			name: "zero_row_not_in_bottom",
			args: args{
				matrix: [][]float64{
					{1, 1},
					{0, 0},
					{0, 0},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allZeroRowsAreAtBottom(tt.args.matrix); got != tt.want {
				t.Errorf("allZeroRowsAreAtBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allPivotEntriesAreRightOfPivotbove(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should_be_true",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "should_be_false",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 1, 0},
				},
			},
			want: false,
		},
		{
			name: "should_be_true_if_row_is_not_inmediatly_right",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "should_be_false_if_last_row_is_not_right",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 1, 0, 0},
				},
			},
			want: false,
		},
		{
			name: "should_be_true_if_only_0s",
			args: args{
				matrix: [][]float64{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should_be_true_if_only_one_1",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should_be_true_if_only_one_1_at_bottom",
			args: args{
				matrix: [][]float64{
					{1, 0, 0, 0},
					{0, 0, 0, 1},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPivotEntriesAreRightOfPivotbove(tt.args.matrix); got != tt.want {
				t.Errorf("allPivotEntriesAreRightOfPivotbove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allPivotsEqualTo1(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be true",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "should be false pivot is 2",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 2, 0},
					{0, 0, 1},
				},
			},
			want: false,
		},
		{
			name: "should be true entries after pivot are not 0",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{0, 1, 1},
					{0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "should be true with only two pivots",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{0, 0, 1},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should be true if there are no pivots",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should be true if pivot is last column",
			args: args{
				matrix: [][]float64{
					{0, 0, 1},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPivotsEqualTo1(tt.args.matrix); got != tt.want {
				t.Errorf("allPivotsEqualTo1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allEntriesInBaseColumnAre0(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be true",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "should be false multiple non 0 entries in column",
			args: args{
				matrix: [][]float64{
					{1, 1, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: false,
		},
		{
			name: "should be true all entries are 0",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should be false all entries are 1",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			want: false,
		},
		{
			name: "should be false non 0 entry in last row",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 0, 1},
					{1, 0, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allEntriesInBaseColumnAre0(tt.args.matrix); got != tt.want {
				t.Errorf("allEntriesInBaseColumnAre0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsReducedRowEchelonForm(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty matrix",
			args: args{
				matrix: [][]float64{},
			},
			want: true,
		},
		{
			name: "should_pass",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "should_not_pass",
			args: args{
				matrix: [][]float64{
					{1, 0, 1},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: false,
		},
		{
			name: "should_pass_0_rows_at_bottom",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should_pass_all_0s",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should_not_pass_one_1_bottom",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{1, 0, 0},
				},
			},
			want: false,
		},
		{
			name: "should_not_pass_not_1s",
			args: args{
				matrix: [][]float64{
					{2, 0, 0},
					{0, 2, 0},
					{0, 0, 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReducedRowEchelonForm(tt.args.matrix); got != tt.want {
				t.Errorf("IsReducedRowEchelonForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPivotEntries(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "should_return pivots",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: [][]int{
				{0, 0},
				{1, 1},
				{2, 2},
			},
		},
		{
			name: "should_return_only_first_pivot",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0},
			},
		},
		{
			name: "should_return_only_first_pivot_in_first_row",
			args: args{
				matrix: [][]float64{
					{0, 1, 1},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 1},
			},
		},
		{
			name: "should_return_only_first_pivot_if_rest_0s",
			args: args{
				matrix: [][]float64{
					{0, 1, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 1},
			},
		},
		{
			name: "should_return_only_first_pivot_if_rest_0s",
			args: args{
				matrix: [][]float64{
					{0, 1, 1},
					{0, 0, 1},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 1},
				{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPivotEntries(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPivotEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapRows(t *testing.T) {
	type args struct {
		matrix [][]float64
		i      int
		j      int
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "should_pass",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				i: 0,
				j: 2,
			},
			want: [][]float64{
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			},
		},
		{
			name: "should_pass_changing_nothing",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				i: 1,
				j: 1,
			},
			want: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "should_pass_changing_first_second",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				i: 0,
				j: 1,
			},
			want: [][]float64{
				{4, 5, 6},
				{1, 2, 3},
				{7, 8, 9},
			},
		},
		{
			name: "should_pass_changing_second_first",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				i: 1,
				j: 0,
			},
			want: [][]float64{
				{4, 5, 6},
				{1, 2, 3},
				{7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapRows(tt.args.matrix, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyScalarByRow(t *testing.T) {
	type args struct {
		matrix   [][]float64
		rowIndex int
		scalar   float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "should_pass_vector",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
				},
				rowIndex: 0,
				scalar:   2,
			},
			want: [][]float64{
				{2, 4, 6},
			},
		},
		{
			name: "should_pass_matrix",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{2, 4, 6},
					{7, 8, 9},
				},
				rowIndex: 1,
				scalar:   2,
			},
			want: [][]float64{
				{1, 2, 3},
				{4, 8, 12},
				{7, 8, 9},
			},
		},
		{
			name: "should_pass_multiply_by_1",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				rowIndex: 1,
				scalar:   1,
			},
			want: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "should_pass_division_vector",
			args: args{
				matrix: [][]float64{
					{4, 8, 12},
				},
				rowIndex: 0,
				scalar:   0.5,
			},
			want: [][]float64{
				{2, 4, 6},
			},
		},
		{
			name: "should_pass_division_matrix",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				rowIndex: 1,
				scalar:   0.5,
			},
			want: [][]float64{
				{1, 2, 3},
				{2, 2.5, 3},
				{7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyRowByScalar(tt.args.matrix, tt.args.rowIndex, tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				fmt.Println("got:")
				for i := range got {
					fmt.Printf("%f ", got[i])
				}
				fmt.Println()

				fmt.Println("want:")
				for i := range got {
					fmt.Printf("%f ", tt.want[i])
				}
				fmt.Println()
				t.Fatalf("unexpected got want")
			}
		})
	}
}

func TestAddRowToRow(t *testing.T) {
	type args struct {
		matrix   [][]float64
		rowToAdd []float64
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "should_pass_add_0",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
				},
				rowToAdd: []float64{0, 0, 0},
				rowIndex: 0,
			},
			want: [][]float64{{1, 2, 3}},
		},
		{
			name: "should_pass_add_row_1s",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
				},
				rowToAdd: []float64{1, 1, 1},
				rowIndex: 0,
			},
			want: [][]float64{{2, 3, 4}},
		},
		{
			name: "should_pass_add_negative",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
				},
				rowToAdd: []float64{-1, -1, -1},
				rowIndex: 0,
			},
			want: [][]float64{{0, 1, 2}},
		},
		{
			name: "should_pass_add_floats",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
				},
				rowToAdd: []float64{0.5, 0.5, 0.5},
				rowIndex: 0,
			},
			want: [][]float64{{1.5, 2.5, 3.5}},
		},
		{
			name: "should_pass_add_to_matrix",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				rowToAdd: []float64{1, 2, 3},
				rowIndex: 1,
			},
			want: [][]float64{
				{1, 2, 3},
				{5, 7, 9},
				{7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRowToRow(tt.args.matrix, tt.args.rowToAdd, tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddRowToRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapRows0sToBottom(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "swap_rows",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{1, 2, 3},
				}},
			want: [][]float64{
				{1, 2, 3},
				{0, 0, 0},
			},
		},
		{
			name: "do_nothing",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{0, 0, 0},
				}},
			want: [][]float64{
				{1, 2, 3},
				{0, 0, 0},
			},
		},
		{
			name: "do_nothing_all_0s",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
				}},
			want: [][]float64{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "swap_last_column_is_not_0",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 1},
				}},
			want: [][]float64{
				{0, 0, 1},
				{0, 0, 0},
			},
		},
		{
			name: "do_nothing_non_0s",
			args: args{
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				}},
			want: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "swap_first_row_is_0s",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{4, 5, 6},
					{7, 8, 9},
				}},
			want: [][]float64{
				{4, 5, 6},
				{7, 8, 9},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapRows0sToBottom(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapRows0sToBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToRowEchelonForm(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "only 1 row",
			args: args{
				matrix: [][]float64{
					{1},
				},
			},
			want: [][]float64{
				{1},
			},
		},
		{
			name: "only 1 row mutliplied by scalar",
			args: args{
				matrix: [][]float64{
					{5},
				},
			},
			want: [][]float64{
				{1},
			},
		},
		{
			name: "should pass one row with 0s and multiple non 0s",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{4, 8, 16},
				},
			},
			want: [][]float64{
				{1, 2, 4},
				{0, 0, 0},
			},
		},
		{
			name: "nothing to do",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "nothing to do",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "single 1 multiple of 3",
			args: args{
				matrix: [][]float64{
					{3, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "example 0",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{1, 0, 1},
					{0, 0, 1},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 0, 1},
				{0, 0, 0},
			},
		},
		{
			name: "example 1.1",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 0, 0},
					{0, 1, 1},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 1},
				{0, 0, 0},
			},
		},
		{
			name: "example 1.2",
			args: args{
				matrix: [][]float64{
					{1, 1, 0},
					{0, 1, 1},
					{0, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, -1},
				{0, 1, 1},
				{0, 0, 0},
			},
		},
		{
			name: "example 1.3",
			args: args{
				matrix: [][]float64{
					{1, 1, 0},
					{0, 1, 1},
					{0, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, -1},
				{0, 1, 1},
				{0, 0, 0},
			},
		},
		{
			name: "example 1.4",
			args: args{
				matrix: [][]float64{
					{1, -1, 0},
					{-4, 4, 1},
					{0, 0, -1},
				},
			},
			want: [][]float64{
				{1, -1, 0},
				{0, 0, 1},
				{0, 0, 0},
			},
		},
		{
			name: "example 1.5",
			args: args{
				matrix: [][]float64{
					{2, -2, 1},
					{-2, 5, 6},
					{10, 7, 1},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			name: "example 1.6",
			args: args{
				matrix: [][]float64{
					{-1, -2, 3},
					{-2, -3, -5},
					{1, 5, 5},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			name: "example 1.7",
			args: args{
				matrix: [][]float64{
					{3, 9, -3, 24},
					{1, -3, 11, -2},
					{-2, 5, -20, -5},
				},
			},
			want: [][]float64{
				{1, 0, 5, 0},
				{0, 1, -2, 0},
				{0, 0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRowEchelonForm(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToRowEchelonForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapLargetsLeftmostNonzeroEntry(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "nothing do do",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "swap 1 and 2",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{2, 0, 0},
				},
			},
			want: [][]float64{
				{2, 0, 0},
				{1, 0, 0},
			},
		},
		{
			name: "swap 1 and 2 other bigger number in same row",
			args: args{
				matrix: [][]float64{
					{1, 9999, 9999},
					{2, 0, 0},
				},
			},
			want: [][]float64{
				{2, 0, 0},
				{1, 9999, 9999},
			},
		},
		{
			name: "swapped identity matrix",
			args: args{
				matrix: [][]float64{
					{0, 0, 1},
					{0, 1, 0},
					{1, 0, 0},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapLargetsLeftmostNonzeroEntry(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapLargetsLeftmostNonzeroEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddMatrices(t *testing.T) {
	type args struct {
		matrixA [][]float64
		matrixB [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "add 0s",
			args: args{
				matrixA: [][]float64{{0, 0, 0}},
				matrixB: [][]float64{{0, 0, 0}},
			},
			want: [][]float64{{0, 0, 0}},
		},
		{
			name: "one matrix is 0",
			args: args{
				matrixA: [][]float64{{1, 0, 0}},
				matrixB: [][]float64{{0, 0, 0}},
			},
			want: [][]float64{{1, 0, 0}},
		},
		{
			name: "simple sum",
			args: args{
				matrixA: [][]float64{{1, 1, 1}},
				matrixB: [][]float64{{1, 1, 1}},
			},
			want: [][]float64{{2, 2, 2}},
		},
		{
			name: "should be zero matrix",
			args: args{
				matrixA: [][]float64{{-1, -1, -1}},
				matrixB: [][]float64{{1, 1, 1}},
			},
			want: [][]float64{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMatrices(tt.args.matrixA, tt.args.matrixB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyMatrixByScalar(t *testing.T) {
	type args struct {
		matrix [][]float64
		scalar float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "multtiply by 0",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				scalar: 0,
			},
			want: [][]float64{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "multtiply by 1",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				scalar: 1,
			},
			want: [][]float64{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			name: "multtiply by 2",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				scalar: 2,
			},
			want: [][]float64{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
		},
		{
			name: "multtiply identity matrix by 9",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 2, 0},
					{0, 0, 3},
				},
				scalar: 9,
			},
			want: [][]float64{
				{9, 0, 0},
				{0, 18, 0},
				{0, 0, 27},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyMatrixByScalar(tt.args.matrix, tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyMatrixByScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroMatrix(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should_be_true",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "should_be_false",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: false,
		},
		{
			name: "should_be_end_row",
			args: args{
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("isZeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canMultiplyMatrices(t *testing.T) {
	type args struct {
		matrixA [][]float64
		matrixB [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty matrices",
			args: args{
				matrixA: [][]float64{},
				matrixB: [][]float64{},
			},
			want: true,
		},
		{
			name: "can multiply 3x3 3x3",
			args: args{
				matrixA: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				matrixB: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "can multiply 2x3 3x2",
			args: args{
				matrixA: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
				},
				matrixB: [][]float64{
					{0, 0},
					{0, 0},
					{0, 0},
				},
			},
			want: true,
		},
		{
			name: "can multiply  3x2 2x3",
			args: args{
				matrixA: [][]float64{
					{0, 0},
					{0, 0},
					{0, 0},
				},
				matrixB: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "can multiply 1x1 1x1",
			args: args{
				matrixA: [][]float64{
					{0},
				},
				matrixB: [][]float64{
					{0},
				},
			},
			want: true,
		},
		{
			name: "cannot multiply 1x3 1x1",
			args: args{
				matrixA: [][]float64{
					{0, 0, 0},
				},
				matrixB: [][]float64{
					{0},
				},
			},
			want: false,
		},
		{
			name: "cannot multiply 1x3 1x1",
			args: args{
				matrixA: [][]float64{
					{0, 0, 0},
				},
				matrixB: [][]float64{
					{0},
				},
			},
			want: false,
		},
		{
			name: "cannot multiply 2x3 4x2",
			args: args{
				matrixA: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
				},
				matrixB: [][]float64{
					{0, 0},
					{0, 0},
					{0, 0},
					{0, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanMultiplyMatrices(tt.args.matrixA, tt.args.matrixB); got != tt.want {
				t.Errorf("canMultuplyMatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyMatrices(t *testing.T) {
	type args struct {
		matrixA [][]float64
		matrixB [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "simple case 1x1 1x1",
			args: args{
				matrixA: [][]float64{
					{3},
				},
				matrixB: [][]float64{
					{2},
				},
			},
			want: [][]float64{
				{6},
			},
		},
		{
			name: "simple case 1x2 2x1",
			args: args{
				matrixA: [][]float64{
					{1, 2},
				},
				matrixB: [][]float64{
					{2},
					{3},
				},
			},
			want: [][]float64{
				{1*2 + 2*3},
			},
		},
		{
			name: "example 1",
			args: args{
				matrixA: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				matrixB: [][]float64{
					{1, 2, 1},
					{2, 4, 6},
					{7, 2, 5},
				},
			},
			want: [][]float64{
				{26, 16, 28},
				{56, 40, 64},
				{86, 64, 100},
			},
		},
		{
			name: "example 2",
			args: args{
				matrixA: [][]float64{
					{1, 0, 0},
					{5, 1, 0},
					{0, 0, 1},
				},
				matrixB: [][]float64{
					{-1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: [][]float64{
				{-1, 0, 0},
				{-5, 1, 0},
				{0, 0, 1},
			},
		},
		{
			name: "example 3",
			args: args{
				matrixA: [][]float64{
					{1, 0},
					{0, float64(1) / float64(2)},
				},
				matrixB: [][]float64{
					{-1, 0},
					{0, 0},
				},
			},
			want: [][]float64{
				{-1, 0},
				{0, 0},
			},
		},
		{
			name: "example 4",
			args: args{
				matrixA: [][]float64{
					{1, -2},
					{0, 1},
				},
				matrixB: [][]float64{
					{-1, 0},
					{0, 0},
				},
			},
			want: [][]float64{
				{-1, 0},
				{0, 0},
			},
		},
		{
			name: "example 5",
			args: args{
				matrixA: [][]float64{
					{-1, -5, 1},
					{-5, -5, 5},
					{2, 5, -3},
				},
				matrixB: [][]float64{
					{float64(-1) / float64(2), float64(-1) / float64(2), -1},
					{float64(-1) / float64(4), float64(1) / float64(20), 0},
					{float64(-3) / float64(4), float64(-1) / float64(4), -1},
				},
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyMatrices(tt.args.matrixA, tt.args.matrixB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiplyMatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateIdentityMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: [][]float64{},
		},
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: [][]float64{
				{1},
			},
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: [][]float64{
				{1, 0},
				{0, 1},
			},
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateIdentityMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateIdentityMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEliminationMatrix(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{
			name: "size 0",
			args: args{
				matrix: [][]float64{},
			},
			want: [][]float64{},
		},
		{
			name: "size 1",
			args: args{
				matrix: [][]float64{
					{1},
				},
			},
			want: [][]float64{
				{1},
			},
		},
		{
			name: "multiple of 5",
			args: args{
				matrix: [][]float64{
					{5},
				},
			},
			want: [][]float64{
				{float64(1) / float64(5)},
			},
		},
		{
			name: "size 2",
			args: args{
				matrix: [][]float64{
					{5, 0},
					{0, 0},
				},
			},
			want: [][]float64{
				{float64(1) / float64(5), 0},
				{0, 1},
			},
		},
		{
			name: "example 0",
			args: args{
				matrix: [][]float64{
					{2, 4},
					{0, -3},
				},
			},
			want: [][]float64{
				{float64(1) / float64(2), float64(2) / float64(3)},
				{0, float64(-1) / float64(3)},
			},
		},
		{
			name: "example",
			args: args{
				matrix: [][]float64{
					{-1, -5, 1},
					{-5, -5, 5},
					{2, 5, -3},
				},
			},
			want: [][]float64{
				{float64(-1) / float64(2), float64(-1) / float64(2), -1},
				{float64(-1) / float64(4), float64(1) / float64(20), 0},
				{float64(-3) / float64(4), float64(-1) / float64(4), -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalMatrix := copyMatrix(tt.args.matrix)
			got := GetEliminationMatrix(tt.args.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEliminationMatrix() = %v, want %v", got, tt.want)
			}

			// TODO confirm inverse works

			multiplied := MultiplyMatrices(got, originalMatrix)
			if !IsRowEchelonForm(multiplied) {
				t.Errorf("expected to be reduced row echelon form")
			}
		})
	}
}

func TestMultiplyVectorByScalar(t *testing.T) {
	type args struct {
		vector []float64
		scalar float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "no changes",
			args: args{
				vector: []float64{},
				scalar: 1,
			},
			want: []float64{},
		},
		{
			name: "multiply by 1",
			args: args{
				vector: []float64{1},
				scalar: 1,
			},
			want: []float64{1},
		},
		{
			name: "multiply by 0",
			args: args{
				vector: []float64{1},
				scalar: 0,
			},
			want: []float64{0},
		},
		{
			name: "multiply nums by 5",
			args: args{
				vector: []float64{1, 2, 3, 4},
				scalar: 5,
			},
			want: []float64{5, 10, 15, 20},
		},
		{
			name: "multiply nums by -5",
			args: args{
				vector: []float64{1, 2, 3, 4},
				scalar: -5,
			},
			want: []float64{-5, -10, -15, -20},
		},
		{
			name: "multiply nums by 0",
			args: args{
				vector: []float64{1, 2, 3, 4},
				scalar: 0,
			},
			want: []float64{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyVectorByScalar(tt.args.vector, tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyVectorByScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nearlyEqual(t *testing.T) {
	type args struct {
		a        float64
		b        float64
		decimals int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exactly equal",
			args: args{
				a:        float64(0),
				b:        float64(0),
				decimals: 1,
			},
			want: true,
		},
		{
			name: "exactly equal",
			args: args{
				a:        float64(5),
				b:        float64(5),
				decimals: 10,
			},
			want: true,
		},
		{
			name: "not equal",
			args: args{
				a:        float64(1),
				b:        float64(2),
				decimals: 1,
			},
			want: false,
		},
		{
			name: "exactly equal with one decimal",
			args: args{
				a:        float64(0.01),
				b:        float64(0.01),
				decimals: 3,
			},
			want: true,
		},
		{
			name: "exactly equal with two decimals",
			args: args{
				a:        float64(0.001),
				b:        float64(0.001),
				decimals: 3,
			},
			want: true,
		},
		{
			name: "exactly equal with three decimals",
			args: args{
				a:        float64(0.0001),
				b:        float64(0.0001),
				decimals: 3,
			},
			want: true,
		},
		{
			name: "equal with three decimals",
			args: args{
				a:        float64(0.0001),
				b:        float64(0.0009),
				decimals: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NearlyEqual(tt.args.a, tt.args.b, tt.args.decimals); got != tt.want {
				t.Errorf("nearlyEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVectorLength(t *testing.T) {
	type args struct {
		vector []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "empty vector",
			args: args{
				vector: []float64{},
			},
			want: 0,
		},
		{
			name: "single number vector",
			args: args{
				vector: []float64{1},
			},
			want: 1,
		},
		{
			name: "two dimensions",
			args: args{
				vector: []float64{1, 1},
			},
			want: math.Sqrt(2),
		},
		{
			name: "two dimensions but longer",
			args: args{
				vector: []float64{2, 2},
			},
			want: math.Sqrt(8),
		},
		{
			name: "two dimensions where one is negative",
			args: args{
				vector: []float64{2, -2},
			},
			want: math.Sqrt(8),
		},
		{
			name: "two dimensions where both are negative",
			args: args{
				vector: []float64{-2, -2},
			},
			want: math.Sqrt(8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVectorLength(tt.args.vector); got != tt.want {
				t.Errorf("GetVectorLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUnitVector(t *testing.T) {
	type args struct {
		vector []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "empty vector",
			args: args{
				vector: []float64{},
			},
			want: []float64{},
		},
		{
			name: "single component",
			args: args{
				vector: []float64{1},
			},
			want: []float64{1},
		},
		{
			name: "two components 1 1",
			args: args{
				vector: []float64{1, 1},
			},
			want: []float64{1 / float64(math.Sqrt(2)), 1 / float64(math.Sqrt(2))},
		},
		{
			name: "two components 2 2",
			args: args{
				vector: []float64{2, 2},
			},
			want: []float64{2 / float64(math.Sqrt(8)), 2 / float64(math.Sqrt(8))},
		},
		{
			name: "example 1",
			args: args{
				vector: []float64{4, -3},
			},
			want: []float64{4 / float64(5), -3 / float64(5)},
		},
		{
			name: "example 2",
			args: args{
				vector: []float64{1, 2, 3},
			},
			want: []float64{1 / float64(math.Sqrt(14)), 2 / float64(math.Sqrt(14)), 3 / float64(math.Sqrt(14))},
		},
		{
			name: "calling unit vector on itself should return the same vector",
			args: args{
				vector: []float64{1 / float64(math.Sqrt(14)), 2 / float64(math.Sqrt(14)), 3 / float64(math.Sqrt(14))},
			},
			want: []float64{1 / float64(math.Sqrt(14)), 2 / float64(math.Sqrt(14)), 3 / float64(math.Sqrt(14))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUnitVector(tt.args.vector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumberOfSolutions(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "identity matrix",
			args: args{
				matrix: [][]float64{
					{1, 0},
					{0, 1},
				},
			},
			want: 1,
		},
		{
			name: "identity matrix before rref",
			args: args{
				matrix: [][]float64{
					{1, 1},
					{0, 1},
				},
			},
			want: 1,
		},
		{
			name: "no solutions",
			args: args{
				matrix: [][]float64{
					{1, 0},
					{0, 0},
				},
			},
			want: 0,
		},
		{
			name: "no solutions all the same",
			args: args{
				matrix: [][]float64{
					{1, 1},
					{1, 1},
				},
			},
			want: 0,
		},
		{
			name: "no solutions all 0s",
			args: args{
				matrix: [][]float64{
					{0, 0},
					{0, 0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberOfSolutions(tt.args.matrix); got != tt.want {
				t.Errorf("GetNumberOfSolutions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatrixSpan(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple solution",
			args: args{
				matrix: [][]float64{},
			},
			want: 0,
		},
		{
			name: "single line",
			args: args{
				matrix: [][]float64{
					{1},
				},
			},
			want: 1,
		},
		{
			name: "single line size not 1",
			args: args{
				matrix: [][]float64{
					{9},
				},
			},
			want: 1,
		},
		{
			name: "multiple single line vectors",
			args: args{
				matrix: [][]float64{
					{9},
					{9},
				},
			},
			want: 1,
		},
		{
			name: "vector in two dimensions",
			args: args{
				matrix: [][]float64{
					{1, 1},
				},
			},
			want: 1,
		},
		{
			name: "two linearly independent vectors in two dimensions",
			args: args{
				matrix: [][]float64{
					{1, 0},
					{0, 1},
				},
			},
			want: 2,
		},
		{
			name: "three vectors in two dimensions",
			args: args{
				matrix: [][]float64{
					{1, 0},
					{0, 1},
					{1, 1},
				},
			},
			want: 2,
		},
		{
			name: "three linearly indepdendent vectors in three dimensions",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: 3,
		},
		{
			name: "three dependent vectors in three dimensions",
			args: args{
				matrix: [][]float64{
					{1, 1, 1},
					{2, 2, 2},
					{3, 3, 3},
				},
			},
			want: 1,
		},
		{
			name: "two independent vectors in three dimensions",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMatrixSpan(tt.args.matrix); got != tt.want {
				t.Errorf("GetMatrixSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areVectorsLinearlyIndependentByElimination(t *testing.T) {
	type args struct {
		vectors [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty vectors",
			args: args{
				vectors: [][]float64{},
			},
			want: true,
		},
		{
			name: "identitymatrix",
			args: args{
				vectors: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: true,
		},
		{
			name: "not indep same vector",
			args: args{
				vectors: [][]float64{
					{1, 0},
					{1, 0},
				},
			},
			want: false,
		},
		{
			name: "not indep multiplied by scalar",
			args: args{
				vectors: [][]float64{
					{1, 1},
					{4, 4},
				},
			},
			want: false,
		},
		{
			name: "not indep sum of other vectors",
			args: args{
				vectors: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{1, 1, 0},
				},
			},
			want: false,
		},
		{
			name: "idependent but not in rref",
			args: args{
				vectors: [][]float64{
					{1, -1, 0},
					{0, 1, -3},
					{-2, 0, 1},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areVectorsLinearlyIndependentByGaussianElimination(tt.args.vectors); got != tt.want {
				t.Errorf("areVectorsLinearlyIndependent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areMatricesEqual(t *testing.T) {
	type args struct {
		matrixA [][]float64
		matrixB [][]float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty matrices",
			args: args{
				matrixA: [][]float64{},
				matrixB: [][]float64{},
			},
			want: true,
		},
		{
			name: "single vector",
			args: args{
				matrixA: [][]float64{{1}},
				matrixB: [][]float64{{1}},
			},
			want: true,
		},
		{
			name: "two dimensional matrices",
			args: args{
				matrixA: [][]float64{{1, 0}},
				matrixB: [][]float64{{1, 0}},
			},
			want: true,
		},
		{
			name: "two three dimensional matrices",
			args: args{
				matrixA: [][]float64{{1, 2, 3}},
				matrixB: [][]float64{{1, 2, 3}},
			},
			want: true,
		},
		{
			name: "cols is not same as rows",
			args: args{
				matrixA: [][]float64{{1, 2}},
				matrixB: [][]float64{{1, 2, 3}},
			},
			want: false,
		},
		{
			name: "one component is different",
			args: args{
				matrixA: [][]float64{{1, 0, 0, 0}},
				matrixB: [][]float64{{1, 0, 0, 1}},
			},
			want: false,
		},
		{
			name: "inverted dimensions",
			args: args{
				matrixA: [][]float64{{0, 0}},
				matrixB: [][]float64{
					{0},
					{0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areMatricesEqual(tt.args.matrixA, tt.args.matrixB); got != tt.want {
				t.Errorf("areMatricesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dotProduct(t *testing.T) {
	type args struct {
		vectorA []float64
		vectorB []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "empty vectors",
			args: args{
				vectorA: []float64{},
				vectorB: []float64{},
			},
			want: 0,
		},
		{
			name: "simple case",
			args: args{
				vectorA: []float64{1},
				vectorB: []float64{1},
			},
			want: 1,
		},
		{
			name: "simple case two dimensions",
			args: args{
				vectorA: []float64{1, 1},
				vectorB: []float64{1, 1},
			},
			want: 2,
		},
		{
			name: "vectorA is zero vector",
			args: args{
				vectorA: []float64{0, 0},
				vectorB: []float64{1, 1},
			},
			want: 0,
		},
		{
			name: "both are zero vector",
			args: args{
				vectorA: []float64{0, 0},
				vectorB: []float64{0, 0},
			},
			want: 0,
		},
		{
			name: "vectorB is zero vector",
			args: args{
				vectorA: []float64{1, 1},
				vectorB: []float64{0, 0},
			},
			want: 0,
		},
		{
			name: "vectorB is zero vector",
			args: args{
				vectorA: []float64{1, 1},
				vectorB: []float64{0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DotProduct(tt.args.vectorA, tt.args.vectorB); got != tt.want {
				t.Errorf("dotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areVectorsLinearlyIndependentByCauchySchwarz(t *testing.T) {
	type args struct {
		vectorA []float64
		vectorB []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty vectors",
			args: args{
				vectorA: []float64{},
				vectorB: []float64{},
			},
			want: true,
		},
		{
			name: "identity matrix",
			args: args{
				vectorA: []float64{1, 0},
				vectorB: []float64{0, 1},
			},
			want: true,
		},
		{
			name: "not independent, same vector",
			args: args{
				vectorA: []float64{1, 1},
				vectorB: []float64{1, 1},
			},
			want: false,
		},
		{
			name: "not independent, scalar multiplied",
			args: args{
				vectorA: []float64{1, 1},
				vectorB: []float64{5, 5},
			},
			want: false,
		},
		{
			name: "example 1",
			args: args{
				vectorA: []float64{3, 4},
				vectorB: []float64{-6, -8},
			},
			want: false,
		},
		{
			name: "identity matrix",
			args: args{
				vectorA: GenerateIdentityMatrix(100)[0],
				vectorB: GenerateIdentityMatrix(100)[0],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areVectorsLinearlyIndependentByCauchySchwarz(tt.args.vectorA, tt.args.vectorB); got != tt.want {
				t.Errorf("areVectorsLinearlyIndependentByCauchySchwarz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areVectorsLinearlyIndependentByTriangularInequality(t *testing.T) {
	type args struct {
		vectorA []float64
		vectorB []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty vector",
			args: args{
				vectorA: []float64{},
				vectorB: []float64{},
			},
			want: true,
		},
		{
			name: "1 dimension",
			args: args{
				vectorA: []float64{1},
				vectorB: []float64{1},
			},
			want: false,
		},
		{
			name: "identity matrix",
			args: args{
				vectorA: []float64{1, 0},
				vectorB: []float64{0, 1},
			},
			want: true,
		},
		{
			name: "scalar multiplication",
			args: args{
				vectorA: []float64{1, 0, 0},
				vectorB: []float64{5, 0, 0},
			},
			want: false,
		},
		{
			name: "same vector summed",
			args: args{
				vectorA: []float64{1, 2, 3},
				vectorB: []float64{2, 4, 6},
			},
			want: false,
		},
		{
			name: "linearly independent in 5 dimensions",
			args: args{
				vectorA: []float64{1, 0, 0, 0, 0},
				vectorB: []float64{0, 0, 1, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areVectorsLinearlyIndependentByTriangularInequality(tt.args.vectorA, tt.args.vectorB); got != tt.want {
				t.Errorf("areVectorsLinearlyIndependentByTriangularInequality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadMatrix(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    [][]float64
		wantErr bool
	}{
		{
			name: "single row",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("1 0 0 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{1, 0, 0},
			},
		},
		{
			name: "single number",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("1 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{1},
			},
		},
		{
			name: "single number with decimals",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("1.234 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{1.234},
			},
		},
		{
			name: "identity matrix",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("1 0 0 \n")
					buffer.WriteString("0 1 0 \n")
					buffer.WriteString("0 0 1 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			name: "empty matrix",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					return buffer
				}(),
			},
			want: [][]float64{},
		},
		{
			name: "more columns than rows",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("1 0 0 0 \n")
					buffer.WriteString("0 1 0 0 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
			},
		},
		{
			name: "big number",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("11111111 22222222 \n")
					buffer.WriteString("33333333 44444444 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{11111111, 22222222},
				{33333333, 44444444},
			},
		},
		{
			name: "big number with decimals",
			args: args{
				input: func() io.Reader {
					buffer := bytes.NewBuffer(make([]byte, 0, 1024))
					buffer.WriteString("1111.1111 2222.2222 \n")
					buffer.WriteString("3333.3333 4444.4444 \n")
					return buffer
				}(),
			},
			want: [][]float64{
				{1111.1111, 2222.2222},
				{3333.3333, 4444.4444},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadMatrix(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveMatrix(t *testing.T) {
	type args struct {
		matrix [][]float64
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "empty matrix",
			args: args{
				matrix: [][]float64{},
			},
			wantOut: func() string {
				a := ``
				return a
			}(),
			wantErr: false,
		},
		{
			name: "identity matrix",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			wantOut: func() string {
				a := "1 0 0 \n0 1 0 \n0 0 1 \n"
				return a
			}(),
			wantErr: false,
		},
		{
			name: "2x3 matrix",
			args: args{
				matrix: [][]float64{
					{1, 0, 0},
					{0, 1, 0},
				},
			},
			wantOut: func() string {
				a := "1 0 0 \n0 1 0 \n"
				return a
			}(),
			wantErr: false,
		},
		{
			name: "with decimals",
			args: args{
				matrix: [][]float64{
					{1.1, 0.1},
					{0.1, 1.1},
				},
			},
			wantOut: func() string {
				a := "1.1 0.1 \n0.1 1.1 \n"
				return a
			}(),
			wantErr: false,
		},
		{
			name: "large numbers with decimals",
			args: args{
				matrix: [][]float64{
					{9999.9999, 9999.9999},
					{9999.9999, 9999.9999},
				},
			},
			wantOut: func() string {
				a := "9999.9999 9999.9999 \n9999.9999 9999.9999 \n"
				return a
			}(),
			wantErr: false,
		},
		{
			name: "negative numbers",
			args: args{
				matrix: [][]float64{
					{-1.25, -1.99},
					{1.25, 1.99},
				},
			},
			wantOut: func() string {
				a := "-1.25 -1.99 \n1.25 1.99 \n"
				return a
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := SaveMatrix(tt.args.matrix, out); (err != nil) != tt.wantErr {
				t.Errorf("SaveMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("SaveMatrix() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1_000},
	{input: 10_000},
	{input: 100_000},
	{input: 1_000_000},
}

func BenchmarkSave(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateIdentityMatrix(v.input)
			}
		})
	}
}

func TestGetAngleBetweenVectors(t *testing.T) {
	type args struct {
		vectorA []float64
		vectorB []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example 1",
			args: args{
				vectorA: []float64{2, -1},
				vectorB: []float64{-1, 4},
			},
			want: 130.601,
		},
		{
			name: "example 2",
			args: args{
				vectorA: []float64{2, 0, -1},
				vectorB: []float64{-1, 4, 2},
			},
			want: 112.976,
		},
		{
			name: "example 3",
			args: args{
				vectorA: []float64{1, -3, 1},
				vectorB: []float64{0, 6, -2},
			},
			want: 162.451,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAngleBetweenVectors(tt.args.vectorA, tt.args.vectorB); !NearlyEqual(got, tt.want, 3) {
				t.Errorf("GetAngleBetweenVectors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_radiansToDegrees(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0",
			args: args{
				radians: 0,
			},
			want: 0,
		},
		{
			name: "30 degrees",
			args: args{
				radians: math.Pi / 6,
			},
			want: 30,
		},
		{
			name: "45 degrees",
			args: args{
				radians: math.Pi / 4,
			},
			want: 45,
		},
		{
			name: "90degrees",
			args: args{
				radians: math.Pi / 2,
			},
			want: 90.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RadiansToDegrees(tt.args.radians); !NearlyEqual(got, tt.want, 3) {
				t.Errorf("radiansToDegrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAreVectorsOrthogonal(t *testing.T) {
	type args struct {
		vectorA []float64
		vectorB []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "are orthogonal",
			args: args{
				vectorA: []float64{1, 0},
				vectorB: []float64{0, 1},
			},
			want: true,
		},
		{
			name: "are parallel",
			args: args{
				vectorA: []float64{1, 1},
				vectorB: []float64{2, 2},
			},
			want: false,
		},
		{
			name: "3 dimensions",
			args: args{
				vectorA: []float64{1, 0, 0},
				vectorB: []float64{0, 1, 0},
			},
			want: true,
		},
		{
			name: "3 dimensions not all 0",
			args: args{
				vectorA: []float64{1, 0, 0},
				vectorB: []float64{0, 1, 1},
			},
			want: true,
		},
		{
			name: "one vector is origin",
			args: args{
				vectorA: []float64{1, 1, 1},
				vectorB: []float64{0, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreVectorsOrthogonal(tt.args.vectorA, tt.args.vectorB); got != tt.want {
				t.Errorf("AreVectorsOrthogonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnitVector(t *testing.T) {
	type args struct {
		vector []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "single field",
			args: args{
				vector: []float64{1},
			},
			want: true,
		},
		{
			name: "two dimensions",
			args: args{
				vector: []float64{0, 1},
			},
			want: true,
		},
		{
			name: "three dimensions",
			args: args{
				vector: []float64{0, 0, 1},
			},
			want: true,
		},
		{
			name: "not unit, two non 0 fields",
			args: args{
				vector: []float64{1, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnitVector(tt.args.vector); got != tt.want {
				t.Errorf("IsUnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
