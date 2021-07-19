package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *Node
		l2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "success",
			args: args{
				l1: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 4}}},
				l2: &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 4}}},
			},
			want: &Node{Val: 1, Next: &Node{Val: 1, Next: &Node{Val: 2,  Next: &Node{Val: 3,  Next: &Node{Val: 4, Next: &Node{Val: 4}}}}}},
		},
		{
			name: "success",
			args: args{
				l1: &Node{},
				l2: &Node{},
			},
			want: &Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
