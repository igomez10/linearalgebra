package linearalgebra

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPivotEntries(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPivotEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}
