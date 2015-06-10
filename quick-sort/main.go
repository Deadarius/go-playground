package main

import "fmt"
import "math/rand"

func main() {
	inArray := rands.Perm(10)
	fmt.Println("in:")
	fmt.Printf("%v", inArray)
	fmt.Println()

	quicksort(inArray)
	fmt.Println()
	fmt.Println("out:")
	fmt.Printf("%v", outArray)
	fmt.Println()
}

func quicksort(arr []int, lo int, hi int) {
  pivot := partition(arr, lo, hi)
  quicksort(arr, lo, pivot)
  quicksort(arr, pivot + 1, hi)
}

func partition(arr []int, lo int, hi int) {
  storedIndex:= lo
  storedValue:= arr[storedIndex]
  pivotValue:= arr[hi]

  for index := lo + 1;  <= hi; ++ {
    cur := arr[index]
    if cur < pivotValue {
      arr[index] = storedValue
      storedValue = arr[storedIndex] = cur
      storedIndex = index
    }
  }
  return storedIndex
}
