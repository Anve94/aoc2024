package algorithm

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := &BinaryTree{}

	// Test initial state of the tree
	// Expected tree:
	if tree.Root != nil {
		t.Fatal("Expected root to be nil for an empty tree")
	}

	// Insert root
	// Expected tree: 10
	tree.Insert(10)
	if tree.Root == nil || tree.Root.Data != 10 {
		t.Fatal("Expected root to be 10 after insertion")
	}

	// Insert left child
	// Expected tree: 10
	//               /
	//              5
	tree.Insert(5)
	if tree.Root.Left == nil || tree.Root.Left.Data != 5 {
		t.Fatal("Expected left child of root to be 5, got nil or incorrect value")
	}

	// Insert right child
	// Expected tree: 10
	//               /  \
	//              5   15
	tree.Insert(15)
	if tree.Root.Right == nil || tree.Root.Right.Data != 15 {
		t.Fatal("Expected right child of root to be 15, got nil or incorrect value")
	}

	// Insert left child of right child
	// Expected tree: 10
	//               /  \
	//              5   15
	//                 /
	//               12
	tree.Insert(12)
	if tree.Root.Right.Left == nil || tree.Root.Right.Left.Data != 12 {
		t.Fatal("Expected left child of right child of root to be 12, got nil or incorrect value")
	}
}
