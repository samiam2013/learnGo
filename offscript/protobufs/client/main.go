package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/samiam2013/learnGo/offscript/protobufs/protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	// start an http client
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	// set the content type to application/x-protobuf
	req.Header.Set("Content-Type", "application/x-protobuf")
	// make the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// read the response
	hotel := &protobuf.Hotel{}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(b, hotel); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", hotel)

}
