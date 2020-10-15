package main

import "fmt"

// Implement the following sorting algorithms:
//  Selection sort,
//  Insertion sort,
//  Merge sort,
//  Quick sort,
//  Stooge Sort.
// Check Wikipedia for descriptions.
func main() {
	slice := []int{10, 72, 2, 820, -71, 492, 7, 20138, 7232}
	selectionSort(slice)
	fmt.Println(slice)
}

func selectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		minI := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minI] {
				minI = j
			}
		}
		if minI != i {
			buf := slice[i]
			slice[i] = slice[minI]
			slice[minI] = buf
		}
	}
}
