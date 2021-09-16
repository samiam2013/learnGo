package elementary

import (
	"bytes"
	"fmt"
	"strconv"
)

// Fizzbuzz takes a limit number of iterations and a test flag for returning a string rather than printing and returning ""
func FizzBuzz(limit int64, test bool) string {
	// iterate until overflow?
	var i int64
	var returnBuf bytes.Buffer
	for i = 1; i <= limit; i++ {
		var buf bytes.Buffer
		if i%3 == 0 {
			buf.WriteString("fizz")
		}
		if i%5 == 0 {
			buf.WriteString("buzz")
		} else if buf.Len() == 0 {
			buf.WriteString(strconv.FormatInt(i, 10))
		}
		if test {
			returnBuf.WriteString(buf.String() + "\n")
		} else {
			fmt.Println(buf.String())
		}
	}
	if returnBuf.Len() == 0 {
		return ""
	} else {
		return returnBuf.String()
	}

}
