package main

func main() {
	//FizzBuzz(0)
}

// FizzBuzz implements the world renowned technical interview boo - boo baby keys barrier
func FizzBuzz(input int) string {
	output := ""
	if input%3 == 0 {
		output = output + "fizz"
	}
	if input%5 == 0 {
		output = output + "buzz"
	}
	return output
}
