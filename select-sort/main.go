package main

import "fmt"
import "math/rand"

func main() {
	inArray := rand.Perm(10)
	fmt.Println("in:")
	fmt.Printf("%v", inArray)
	fmt.Println()

	selectSort(inArray)
	fmt.Println()
	fmt.Println("out:")
	fmt.Printf("%v", inArray)
	fmt.Println()
}

func selectSort(inArray []int) {
	length := len(inArray)
	for marker := 0; marker < length; marker++ {
		min := marker
		minValue := inArray[min]
		origMinValue := minValue
		for i := marker + 1; i < length; i++ {
			cur := inArray[i]
			if cur < minValue {
				min = i
				minValue = inArray[min]
			}
		}
		if minValue != origMinValue {
			inArray[marker] = minValue
			inArray[min] = origMinValue
		}
	}
}
