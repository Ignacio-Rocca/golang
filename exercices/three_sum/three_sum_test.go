package three_sum

import (
	"reflect"
	"testing"
)

func Test_getTripletsThatSumZero(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want [][3]int
	}{
		{
			name: "without duplicate cases",
			args: []int{-1, 0, 1, 2, -4, 4},
			want: [][3]int{{-4, 0, 4}, {-1, 0, 1}},
		},
		{
			name: "with duplicate cases",
			args: []int{-1, 0, 1, 2, -1, -4},
			want: [][3]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "without triplets",
			args: []int{-10, 0, 1, 2, 5, -4},
			want: [][3]int{},
		},
		{
			name: "with empty data entry",
			args: []int{},
			want: [][3]int{},
		},
		{
			name: "without sufficient values",
			args: []int{1,2},
			want: [][3]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTripletsThatSumZero(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTripletsThatSumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
