package tree

import (
	"reflect"
	"testing"
)

func TestPostorderRecursive(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试递归后序遍历",
			args: args{
				node: root,
			},
			wantData: postorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := PostorderRecursive(tt.args.node); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("PostorderRecursive() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestPostorderStack(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试Stack后序遍历",
			args: args{
				node: root,
			},
			wantData: postorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := PostorderStack(tt.args.node); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("PostorderStack() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}