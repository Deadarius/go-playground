package main

import "fmt"
import "math/rand"

func main() {
	inArray := rand.Perm(10)

	fmt.Println("in:")
	fmt.Printf("%v", inArray)

	var length = len(inArray)
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			var current = inArray[j]
			var next = inArray[j+1]
			if current > next {
				inArray[j+1] = current
				inArray[j] = next
			}
		}
	}
	fmt.Println()
	fmt.Println("out:")
	fmt.Printf("%v", inArray)
	fmt.Println()

}
