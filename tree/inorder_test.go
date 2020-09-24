package tree

import (
	"reflect"
	"testing"
)


func TestInorderRecursive(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试递归中序遍历",
			args: args{
				node: root,
			},
			wantData: inorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := InorderRecursive(tt.args.node); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("InorderRecursive() = %v, preorderWant %v", gotData, tt.wantData)
			}
		})
	}
}

func TestInorderStack(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试堆栈中序遍历",
			args: args{
				node: root,
			},
			wantData: inorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := InorderStack(tt.args.node); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("InorderStack() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestInorderMorris(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantData []int
	}{
		{
			name: "测试Morris中序遍历",
			args: args{
				node: root,
			},
			wantData: inorderWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := InorderMorris(tt.args.node); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("InorderMorris() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}