package algorithm

type BinaryTree struct {
	Root *TreeNode
}

func (t *BinaryTree) Insert(number int) {
	if t.Root == nil {
		// No root yet so we create one
		t.Root = &TreeNode{number, nil, nil}
	} else {
		// There is a root so we insert on the node
		t.Root.insert(number)
	}
}

func (n *TreeNode) insert(number int) {
	// Lower numbers compared to current node go left, higher or equal go right
	if number < n.Data {
		if n.Left == nil {
			// No left node so we insert a left node
			n.Left = &TreeNode{number, nil, nil}
		} else {
			// Continue looking down from the already existing left node
			n.Left.insert(number)
		}
	} else {
		if n.Right == nil {
			// No node yet so insert to the right here
			n.Right = &TreeNode{number, nil, nil}
		} else {
			// Continue looking down from right node
			n.Right.insert(number)
		}
	}
}
