package main

import (
	"anve/algorithm"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, world!")
	tree := algorithm.BinaryTree{}
	tree.Insert(5)
	fmt.Println(
		"The answer to the world, the universe and everything may be 42, but the root in this tree is " +
			strconv.Itoa(tree.Root.Data),
	)
}
