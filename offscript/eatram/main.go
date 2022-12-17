package main

func main() {
	// create a huge slice of pointers to strings
	strings := make([]*string, 5_000_000_000)
	// range over the slice
	for i := range strings {
		// create a new string of "bruh"s
		bruh := new(string)
		*bruh = "bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh bruh "
		strings[i] = bruh
	}
}
