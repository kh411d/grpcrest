package main

func main() {
	var block = make(chan struct{})
	go serveGRPC()
	go serveREST()
	<-block
}
