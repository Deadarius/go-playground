package main

import "fmt"
import "math/rand"

func main() {
	inArray := rand.Perm(10)
	fmt.Println("in:")
	fmt.Printf("%v", inArray)
	fmt.Println()

	list := List{}
	for _, val := range inArray {
		list.AddToEnd(val)
	}

	sum := list.Sum()
	fmt.Println()
	fmt.Println("Sum:")
	fmt.Printf("%d", sum)
	fmt.Println()
}

type node struct {
	Value    int
	Previous *node
	Next     *node
}

// List - Bidirectional linked list
type List struct {
	//First - first node of the list
	First *node
	//Last - last node of the list
	Last *node
}

//AddToEnd - Add new node to the end
func (list *List) AddToEnd(value int) {
	newNode := node{Previous: list.Last, Value: value}
	if list.Last != nil {
		list.Last.Next = &newNode
	}

	if list.First == nil {
		list.First = list.Last
	}
	list.Last = &newNode
}

//AddToBegin - Add new node to the begin
func (list *List) AddToBegin(value int) {
	newNode := node{Next: list.First, Value: value}
	list.First.Previous = &newNode
	list.First = &newNode
}

//Sum - sum of all nodes
func (list *List) Sum() int {
	sum := 0
	cur := list.First
	for cur != nil {
		sum += cur.Value
		cur = cur.Next
	}

	return sum
}
