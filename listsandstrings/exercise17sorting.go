package listsandstrings

import (
	"fmt"
	"math/rand"
	"time"
)

// Implement the following sorting algorithms:
//  Selection sort,
//  Insertion sort,
//  Merge sort,
//  Quick sort,
//  Stooge Sort.
// Check Wikipedia for descriptions.
func exercise17() {
	size := 30000
	slice := makeRandIntSlice(size)
	fmt.Println("number of elements for each sort: ", size)
	start := time.Now()

	selectionSort(slice)

	duration := time.Since(start)
	fmt.Println("selectionSort took ", duration)
	slice = makeRandIntSlice(size)
	start = time.Now()

	insertionSort(slice)

	duration = time.Since(start)
	fmt.Println("insertionSort took ", duration)
	slice = makeRandIntSlice(size)
	start = time.Now()

	mergeSort(slice)

	duration = time.Since(start)
	fmt.Println("mergeSort took ", duration)
	slice = makeRandIntSlice(size)
	start = time.Now()

	quickSort(slice, 0, len(slice)-1)

	duration = time.Since(start)
	fmt.Println("quickSort took ", duration)

	if false { // too slow don't run
		slice = makeRandIntSlice(size)
		start = time.Now()

		stoogeSort(slice, 0, len(slice)-1)

		duration = time.Since(start)
		fmt.Println("stoogeSort took (elements^(1/2))", duration)
	}
}

// copied from wikipedia
func stoogeSort(slice []int, low, high int) []int {
	var buf, third int
	if slice[low] > slice[high] {
		buf = slice[high]
		slice[high] = slice[low]
		slice[low] = buf
	}
	if (high - low + 1) > 2 {
		third = (high - low + 1) / 3
		stoogeSort(slice, low, high-third)
		stoogeSort(slice, low+third, high)
		stoogeSort(slice, low, high-third)
	}
	return slice
}

// copied from wikipedia from pseudo-code
func quickSort(slice []int, low, high int) {
	if low < high {
		p := quickPartition(slice, low, high)
		quickSort(slice, low, p-1)
		quickSort(slice, p+1, high)
	}
}

func quickPartition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low
	var buf int
	for j := low; j <= high; j++ {
		if slice[j] < pivot {
			buf = slice[j]
			slice[j] = slice[i]
			slice[i] = buf
			i++
		}
	}
	buf = slice[high]
	slice[high] = slice[i]
	slice[i] = buf
	return i
}

// copied from geeks for geeks from python
func mergeSort(slice []int) {
	if len(slice) > 1 {
		middle := len(slice) / 2
		leftSlice := make([]int, len(slice[:middle]))
		copy(leftSlice, slice[:middle])
		rightSlice := make([]int, len(slice[middle:]))
		copy(rightSlice, slice[middle:])

		mergeSort(leftSlice)
		mergeSort(rightSlice)

		i, j, k := 0, 0, 0
		for (i < len(leftSlice)) && (j < len(rightSlice)) {
			if leftSlice[i] < rightSlice[j] {
				slice[k] = leftSlice[i]
				i++
			} else {
				slice[k] = rightSlice[j]
				j++
			}
			k++
		}

		for i < len(leftSlice) {
			slice[k] = leftSlice[i]
			i++
			k++
		}

		for j < len(rightSlice) {
			slice[k] = rightSlice[j]
			j++
			k++
		}
	}
}

// I did write this one, it's slow
func insertionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := len(slice) - 1; j > 0; j-- {
			if slice[j] < slice[j-1] {
				buf := slice[j]
				slice[j] = slice[j-1]
				slice[j-1] = buf
			}
		}
	}
}

// I also wrote this one, it's also slow.
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

func makeRandIntSlice(length int) []int {
	seed := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(seed)
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = randomer.Intn(length)
	}
	return slice
}
