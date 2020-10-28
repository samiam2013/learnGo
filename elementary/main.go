package main

func main() {
	/*
		exercise01()
		exercise02()
		exercise03()
		exercise04()
		exercise05()
		exercise06()
		exercise07()
		exercise08()
		exercise09()
		exercise10()
		exercise11()
		fizzBuzz()
	*/
	testAnagram()
}

func triangular(input int64) int64 {
	sum := int64(0)
	for input > 0 {
		sum = sum + input
		input = input - 1
	}
	return sum
}
