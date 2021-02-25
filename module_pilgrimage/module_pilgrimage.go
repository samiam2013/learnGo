package ModulePilgrimage

// FizzBuzz does the fizzbuzz thing.
func FizzBuzz(input int) string {
	var output string
	if input%3 == 0 {
		output = output + "fizz"
	}
	if input%6 == 0 {
		output = output + "buzz"
	}
	return output
}
