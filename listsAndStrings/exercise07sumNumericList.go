package listsAndStrings

import (
	"container/list"
	"fmt"
)

// Write three functions that compute the sum of the numbers in a list:
//  using a for-loop, a while-loop (these are the same construct in go)
//  recursion
func exercise07() {
	numbers := list.New()
	for _, v := range []int64{1, 130, 1920, 72302918, 2828, 17183, 81, 1038} {
		numbers.PushBack(v)
	}
	fmt.Println("sum result of for loop: ", sumFor(numbers))
	fmt.Println("sum result of recursion: ", sumRecursion(numbers, 0))
}

func sumFor(l *list.List) int64 {
	sum := int64(0)
	for e := l.Front(); e != nil; e = e.Next() {
		sum = sum + e.Value.(int64)
	}
	return sum
}

func sumRecursion(l *list.List, sum int64) int64 {
	e := l.Front()
	sum = sum + e.Value.(int64)
	if e.Next() != nil {
		l.Remove(e)
		return sumRecursion(l, sum)
	}
	return sum
}
