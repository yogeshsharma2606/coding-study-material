package trees

import (
	"reflect"
	"testing"
)

// build a sample BST:      5
//
//	   /   \
//	  3     8
//	 / \   / \
//	2  4  7  9
func sampleBST() *TreeNode {
	return &TreeNode{5,
		&TreeNode{3, &TreeNode{Val: 2}, &TreeNode{Val: 4}},
		&TreeNode{8, &TreeNode{Val: 7}, &TreeNode{Val: 9}},
	}
}

func TestMaxDepth(t *testing.T) {
	if got := MaxDepth(sampleBST()); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := MaxDepth(nil); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestInorderTraversal(t *testing.T) {
	got := InorderTraversal(sampleBST())
	if !reflect.DeepEqual(got, []int{2, 3, 4, 5, 7, 8, 9}) {
		t.Errorf("got %v", got)
	}
}

func TestLevelOrder(t *testing.T) {
	got := LevelOrder(sampleBST())
	want := [][]int{{5}, {3, 8}, {2, 4, 7, 9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestIsValidBST(t *testing.T) {
	if !IsValidBST(sampleBST()) {
		t.Error("expected valid")
	}
	bad := &TreeNode{5, &TreeNode{Val: 6}, &TreeNode{Val: 8}}
	if IsValidBST(bad) {
		t.Error("expected invalid")
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	root := sampleBST()
	p := root.Left.Left  // 2
	q := root.Left.Right // 4
	if LowestCommonAncestor(root, p, q).Val != 3 {
		t.Error("LCA of 2,4 should be 3")
	}
	if LowestCommonAncestor(root, root.Left, root.Right).Val != 5 {
		t.Error("LCA of 3,8 should be 5")
	}
}

func TestDiameterOfBinaryTree(t *testing.T) {
	if got := DiameterOfBinaryTree(sampleBST()); got != 4 {
		t.Errorf("got %d", got)
	}
}

func TestInvertTree(t *testing.T) {
	got := InorderTraversal(InvertTree(sampleBST()))
	if !reflect.DeepEqual(got, []int{9, 8, 7, 5, 4, 3, 2}) {
		t.Errorf("got %v", got)
	}
}

func TestKthSmallest(t *testing.T) {
	if got := KthSmallest(sampleBST(), 3); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := KthSmallest(sampleBST(), 1); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestRightSideView(t *testing.T) {
	got := RightSideView(sampleBST())
	if !reflect.DeepEqual(got, []int{5, 8, 9}) {
		t.Errorf("got %v", got)
	}
}

func TestCodec(t *testing.T) {
	c := &Codec{}
	root := sampleBST()
	got := InorderTraversal(c.Deserialize(c.Serialize(root)))
	if !reflect.DeepEqual(got, []int{2, 3, 4, 5, 7, 8, 9}) {
		t.Errorf("roundtrip got %v", got)
	}
	if c.Deserialize(c.Serialize(nil)) != nil {
		t.Error("nil roundtrip failed")
	}
}
