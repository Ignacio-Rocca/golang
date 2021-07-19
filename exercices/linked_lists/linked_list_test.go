package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyNodesList(t *testing.T) {
	node3 := &Node{Val: 3, Next: nil}
	node2 := &Node{Val: 2, Next: node3}
	node1 := &Node{Val: 1, Next: node2}

	nodeWant := node1

	tests := []struct {
		name string
		args *Node
		want *Node
	}{
		{
			name: "success",
			args: node1,
			want: nodeWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CopyNodesList(tt.args)
			assert.NotEqual(t, tt.args, got)
			for tt.want != nil || got != nil {
				assert.Equal(t, tt.want.Val, got.Val)
				tt.want = tt.want.Next
				got = got.Next
			}

		})
	}
}
