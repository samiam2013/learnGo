package main

import "fmt"

// Write functions that add, subtract, and multiply two numbers in their
//  digit-list representation (and return a new digit list).
// Try different bases --- how about no?
func exercise15() {
	fmt.Println("this one should actually be fun.")
	x := []int{2, 1, 0, 0, 1}
	y := []int{0, 9, 9, 9}

	fmt.Println(addDigitSlice(x, y))
	fmt.Println(subtractDigitSlice(x, y))
	fmt.Println(multiplyDigitSlice(x, y))
}

func addDigitSlice(x, y []int) []int {
	carry := 0
	longestLen := maxOf(len(x), len(y))
	x = padSlice(x, longestLen)
	y = padSlice(y, longestLen)
	sum := make([]int, longestLen+1)
	for i := maxOf(len(x), len(y)) - 1; i >= 0; i-- {
		buf := 0
		// add values, including carry from previous add
		if carry != 0 {
			buf = carry
		}
		if x[i] > 0 {
			buf += x[i]
		}
		if y[i] > 0 {
			buf += y[i]
		}
		// do the carry check
		if buf >= 10 {
			sum[i+1] = buf % 10
			carry = 1
		} else {
			sum[i+1] = buf
			carry = 0
		}
	}
	if carry == 1 {
		sum[0] = 1
	}
	return sum
}

func subtractDigitSlice(x, y []int) []int {
	longestLen := maxOf(len(x), len(y))
	difference := make([]int, longestLen)
	x = padSlice(x, longestLen)
	y = padSlice(y, longestLen)
	// honestly, this is magic. I'm surprised I got it to work
	for i := longestLen - 1; i >= 0; i-- {
		difference[i] = x[i] - y[i]
		if difference[i] < 0 {
			difference[i] += 10
			for j := i - 1; j >= 0; j-- {
				if x[j] > 0 {
					x[j]--
					break
				} else {
					x[j] += 9
				}
			}
		}
	}
	return difference
}

func multiplyDigitSlice(x, y []int) []int {
	longestLen := maxOf(len(x), len(y))
	x = padSlice(x, longestLen)
	y = padSlice(y, longestLen)
	product := make([]int, (2*longestLen)+2)

	// for each number in the factor
	offset := 0
	for i := longestLen - 1; i >= 0; i-- {
		buf := make([]int, (2*longestLen)+2)
		bufIMax := len(buf) - 1
		carry := 0
		count := 0
		// multiply by each of the other number, doing the carry
		for j := longestLen - 1; j >= 0; j-- {
			value := x[i]*y[j] + carry
			if value > 10 {
				carry = (value - (value % 10)) / 10
			} else {
				carry = 0
			}
			buf[(bufIMax-offset)-count] = (value % 10)
			count++
		}
		if carry > 0 {
			buf[(bufIMax-offset)-count] = carry
		}
		// increase the offset for the sum to build product
		offset++
		product = addDigitSlice(product, buf)
	}
	return product
}

func padSlice(x []int, length int) []int {
	paddedSlice := make([]int, length)
	sizeDiff := length - len(x)
	for i := 0; i < length; i++ {
		if i >= sizeDiff {
			paddedSlice[i] = x[i-sizeDiff]
		} else {
			paddedSlice[i] = 0
		}
	}
	return paddedSlice
}

func maxOf(x, y int) int {
	if x > y {
		return x
	}
	return y
}
