package main

import (
	"net/http"

	"github.com/samiam2013/learnGo/offscript/protobufs/protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	// start an http server on port 8080
	hotel := protobuf.Hotel{
		Id:        1,
		Name:      "Hilton",
		Address:   "1234 Main St",
		Lattitude: 123.456,
		Longitude: 45.678,
	}

	// listen on port 8080
	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// write the protobuf to the response
			w.Header().Set("Content-Type", "application/x-protobuf")
			w.WriteHeader(http.StatusOK)
			out, err := proto.Marshal(&hotel)
			if err != nil {
				panic(err)
			}
			if _, err := w.Write(out); err != nil {
				panic(err)
			}
		}),
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
