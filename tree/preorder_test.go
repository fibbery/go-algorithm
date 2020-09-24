package tree

import (
	"reflect"
	"testing"
)

var (
	data          = []int{5, 9, 4, 2, 6, 9}
	root          = CreateTree(data, 0, len(data))
	preorderWant  = []int{5, 9, 2, 6, 4, 9}
	inorderWant   = []int{2, 9, 6, 5, 9, 4}
	postorderWant = []int{2, 6, 9, 9, 4, 5}
)

func TestPreorder_Recursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试递归前序遍历",
			args: args{
				root: root,
			},
			wantData: preorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := PreorderRecursive(tt.args.root); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Preorder_Recursive() = %v, preorderWant %v", gotData, tt.wantData)
			}
		})
	}
}

func TestPreorder_Stack(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试栈前序遍历",
			args: args{
				root: root,
			},
			wantData: preorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := PreorderStack(tt.args.root); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Preorder_Stack() = %v, preorderWant %v", gotData, tt.wantData)
			}
		})
	}
}

func TestPreorder_Morris(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试Morris前序遍历",
			args: args{
				root: root,
			},
			wantData: preorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := PreorderMorris(tt.args.root); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Preorder_Morris() = %v, preorderWant %v", gotData, tt.wantData)
			}
		})
	}
}
