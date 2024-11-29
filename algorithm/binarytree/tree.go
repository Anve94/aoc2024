package binarytree

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(number int) {
	if t.Root == nil {
		// No root yet so we create one
		t.Root = &Node{number, nil, nil}
	} else {
		// There is a root so we insert on the node
		t.Root.insert(number)
	}
}

func (n *Node) insert(number int) {
	// Lower numbers compared to current node go left, higher or equal go right
	if number < n.Data {
		if n.Left == nil {
			// No left node so we insert a left node
			n.Left = &Node{number, nil, nil}
		} else {
			// Continue looking down from the already existing left node
			n.Left.insert(number)
		}
	} else {
		if n.Right == nil {
			// No node yet so insert to the right here
			n.Right = &Node{number, nil, nil}
		} else {
			// Continue looking down from right node
			n.Right.insert(number)
		}
	}
}
