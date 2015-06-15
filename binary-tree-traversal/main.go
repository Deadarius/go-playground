package main

import "fmt"
import "math/rand"

func main() {
	inArray := rand.Perm(10)
	fmt.Println("in:")
	fmt.Printf("%v", inArray)
	fmt.Println()

	tree := BinaryTree{}
	for _, val := range inArray {
		tree.Add(val)
	}

	writeElement := func(value int) {
		fmt.Printf("%d ", value)
	}

	//sum := list.Sum()
	fmt.Println()
	fmt.Println("Out:")
	tree.Walk(writeElement)
	//fmt.Printf("%d", sum)
	fmt.Println()
}

type node struct {
	Value int
	Left  *node
	Right *node
}

func (current *node) add(value int) {
	if current.Value > value {
		if current.Left == nil {
			current.Left = &node{Value: value}
		} else {
			current.Left.add(value)
		}
	} else {
		if current.Right == nil {
			current.Right = &node{Value: value}
		} else {
			current.Right.add(value)
		}
	}
}

func (current *node) walk(callback func(value int)) {
	if current.Left != nil {
		current.Left.walk(callback)
	}

	callback(current.Value)

	if current.Right != nil {
		current.Right.walk(callback)
	}
}

//BinaryTree - so binary, such tree.
type BinaryTree struct {
	root *node
}

//Add - add new value to the tree
func (tree *BinaryTree) Add(value int) {
	if tree.root == nil {
		tree.root = &(node{Value: value})
	} else {
		tree.root.add(value)
	}

}

//Walk the tree
func (tree *BinaryTree) Walk(callback func(value int)) {
	tree.root.walk(callback)
}
