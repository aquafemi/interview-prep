package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	quickSort([]int{4, 2, 13, 5, 0, -10, 44, 2, 2, 2, 6, 3, 1})
}

func quickSort(list []int) {
	fmt.Println(list)
	quickSortPivot(list, 0, len(list)-1)

	fmt.Println(list)
}

func quickSortPivot(list []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := list[(lo+hi)/2]

	index := partition(list, pivot, lo, hi)

	quickSortPivot(list, lo, index-1)
	quickSortPivot(list, index, hi)
}

func partition(list []int, pivot, lo, hi int) int {
	for lo <= hi {
		for list[lo] < pivot {
			lo++
		}
		for list[hi] > pivot {
			hi--
		}
		if lo <= hi {
			swap(list, lo, hi)
			lo++
			hi--
		}
		fmt.Println("lo", lo)
		fmt.Println("hi", hi)
	}
	return lo
}

func swap(list []int, index1, index2 int) {
	temp := list[index1]
	list[index1] = list[index2]
	list[index2] = temp
}
