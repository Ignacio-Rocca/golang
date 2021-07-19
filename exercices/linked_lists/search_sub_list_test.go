package main

import "testing"

func Test_subListSearch(t *testing.T) {
	type args struct {
		headList    *Node
		headSubList *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success: has sublist",
			args: args{
				headList: &Node{Val:1, Next: &Node{Val: 2}},
				headSubList: &Node{Val:2},
			},
			want: true,
		},
		{
			name: "success: hasn't sublist",
			args: args{
				headList: &Node{Val:1, Next: &Node{Val: 2}},
				headSubList: &Node{Val:3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subListSearch(tt.args.headList, tt.args.headSubList); got != tt.want {
				t.Errorf("subListSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
