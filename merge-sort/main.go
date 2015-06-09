package main

import "fmt"
import "math/rand"

func main() {
	inArray := rand.Perm(10)
	fmt.Println("in:")
	fmt.Printf("%v", inArray)
	fmt.Println()

	outArray := split(inArray)
	fmt.Println()
	fmt.Println("out:")
	fmt.Printf("%v", outArray)
	fmt.Println()
}

func split(inArray []int) []int {
	length := len(inArray)
	if length <= 1 {
		return inArray
	}

	middle := length / 2
	left := inArray[0:middle]
	right := inArray[middle:length]
	left = split(left)
	right = split(right)
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	leftLength := len(left)
	rightLength := len(right)
	//sumLength := leftLength + rightLength
	var result []int
	var i int
	var j int
	for i < leftLength && j < rightLength {
		leftElement := left[i]
		rightElement := right[j]
		if leftElement < rightElement {
			result = append(result, leftElement)
			i++
		} else {
			result = append(result, rightElement)
			j++
		}
	}

	if i < leftLength {
		result = append(result, left[i:]...)
	} else {
		result = append(result, right[j:]...)
	}

	return result
}
