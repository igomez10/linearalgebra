package linearalgebra

import (
	"fmt"
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
			if got := MultiplyScalarByRow(tt.args.matrix, tt.args.rowIndex, tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
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
