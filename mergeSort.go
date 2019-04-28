package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(mergeSort([]int{3, 2, 1, 7, 0, 6, 4, 4, 5, -3}))
}

func mergeSort(values []int) []int {	
	if len(values) == 1 {
		return values
	} else {
		midpoint := len(values) / 2

		// Sort the left
		left := mergeSort(values[0:midpoint])

		// Sort the right
		right := mergeSort(values[midpoint:len(values)])

		// Merge the result		
		return merge(left, right)
	}

}

func merge(left, right []int) []int {
	mergedList := []int{}
	leftIndex, rightIndex := 0, 0
	
	// Add items from each list in order until one list is empty
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			leftIndex, mergedList = appendAtIndex(mergedList, left, leftIndex)
		} else if left[leftIndex] >= right[rightIndex] {
			rightIndex, mergedList = appendAtIndex(mergedList, right, rightIndex)
		}
	}

	// Add remaining items if there are any
	for leftIndex < len(left) {
		leftIndex, mergedList = appendAtIndex(mergedList, left, leftIndex)
	}

	// Add remaining items if there are any
	for rightIndex < len(right) {
		rightIndex, mergedList = appendAtIndex(mergedList, right, rightIndex)
	}

	return mergedList
}

func appendAtIndex(mergedList, subList []int, index int) (int, []int) {
	mergedList = append(mergedList, subList[index])
	index++
	return index, mergedList
}
