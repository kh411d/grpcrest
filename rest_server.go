package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/kh411d/grpcrest/protobuff/helloworld"
)

const (
	restPort = ":9090"
)

func SayHello(w http.ResponseWriter, r *http.Request) {

	var input pb.HelloRequest
	var response pb.HelloReply
	// decode input
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&input)
	r.Body.Close()

	response.Message = "Greeting: Hello " + input.Name
	respBytes, _ := json.Marshal(response)

	w.WriteHeader(200)
	w.Write(respBytes)

	fmt.Println(string(respBytes))
}

func serveREST() {
	http.HandleFunc("/", SayHello) // set router

	fmt.Println("Listening on ", restPort)
	err := http.ListenAndServe(restPort, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
