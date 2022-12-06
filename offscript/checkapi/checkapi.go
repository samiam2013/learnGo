package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func main() {
	// call the mock server
	reader := mockServerCall()
	// create a new decoder and point it at our reader
	decoder := json.NewDecoder(reader)
	// create a variable to hold our JSON response
	var response interface{}
	// decode the JSON into our exchangeRates struct
	err := decoder.Decode(&response)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// print out the type of the response
	fmt.Printf("response type: %T\n", response)

	// mock another call
	reader = mockServerCall()
	// create a new decoder and point it at our reader
	decoder = json.NewDecoder(reader)
	// create a variable to hold our JSON response
	response2 := make(map[string]interface{})
	// decode the JSON into our exchangeRates struct
	err = decoder.Decode(&response2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// print out the type of the response
	fmt.Printf("response2 type: %T\n", response2)

	// check if the types are the same
	if fmt.Sprintf("%T", response) == fmt.Sprintf("%T", response2) {
		fmt.Println("response types are the same")
	} else {
		fmt.Println("response types are different")
	}

	// check if the keys are the same
	r1 := printKeys(response, "")
	r2 := printKeys(response2, "")
	// sort both
	sort.Strings(r1)
	sort.Strings(r2)
	if reflect.DeepEqual(r1, r2) {
		fmt.Println("response keys are the same")
	} else {
		fmt.Println("response keys are different")
	}
	fmt.Printf("\tresponse keys: '%v'\n", r1)
	fmt.Printf("\tresponse2 keys: '%v'\n", r2)

}

func printKeys(v interface{}, prefix string) (result []string) {
	switch v := v.(type) {
	case map[string]interface{}:
		for k := range v {
			// fmt.Println(k)
			result = append(result, prefix+"."+k)
			result = append(result, printKeys(v[k], prefix+"."+k)...)
		}
	case []interface{}:
		for _, u := range v {
			result = printKeys(u, "")
		}
	}
	return
}

func mockServerCall() io.Reader {
	// seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())
	// pick a random number between 1 and 3
	selection := rand.Intn(3) + 1
	fmt.Println("selection:", selection)
	switch selection {
	case 1:
		// return status okay and a set of exchange rates
		return bytes.NewBuffer([]byte(`{"success":true,"data":{"timestamp":158000,"base":"EUR","date":"2020-01-29","rates":{"USD":1.123456}}}`))
	case 2:
		// return status okay and a set of exchange rates
		return bytes.NewBuffer([]byte(`{"success":true,"data":{"timestamp":138000,"base":"CNY","date":"2020-04-29","rates":{"USD":0.0237}}}`))
	case 3:
		// return a different structure with "Success" instead of "success"
		return bytes.NewBuffer([]byte(`{"Success":true,"data":{"timestamp":138000,"base":"CNY","date":"2020-04-29","rates":{"USD":0.0237}}}`))
	default:
		return nil
	}
}
