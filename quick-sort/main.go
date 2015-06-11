package main

import "fmt"
import "math/rand"

func main() {
	inArray := rand.Perm(10)
	fmt.Println("in:")
	fmt.Printf("%v", inArray)
	fmt.Println()

	quicksort(inArray, 0, len(inArray)-1)
	fmt.Println()
	fmt.Println("out:")
	fmt.Printf("%v", inArray)
	fmt.Println()
}

func quicksort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIndex := partition(arr, lo, hi)

	quicksort(arr, lo, pivotIndex)
	quicksort(arr, pivotIndex+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivotIndex := arr[(lo+hi)/2]
	pivotValue := arr[pivotIndex]

	swap(arr, pivotIndex, hi)

	left := lo
	right := hi - 1

	for left < right {
		for arr[left] < pivotValue {
			left++
		}

		for right >= left && arr[right] >= pivotValue {
			right--
		}

		if right > left {
			swap(arr, left, right)
		}
	}

	swap(arr, hi, left)
	return left
}

func swap(arr []int, first int, second int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}
