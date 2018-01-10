# grpcrest
GRPC and Rest testing 

> note: Install all requirement before proceed, "go get"
> see grpc-go docs https://github.com/grpc/grpc-go to install grpc

This is a test case for using GRPC as a great alternative for REST


### Running Benchmark
```
MacBook-Air-3:grpcrest khalid$ go test ./... -bench Bench
BenchmarkGRPC-4   	 1000000	      1105 ns/op
BenchmarkREST-4   	    2000	  15518352 ns/op
```

> The average run time of the function under test for the final value of b.N iterations. 

> In this case, the loop ran 10000000 times, my laptop can execute  GRPC request in 1105 nanoseconds per loop

### Running a standalone server grpcrest
```
MacBook-Air-3:grpcrest khalid$ go build
MacBook-Air-3:grpcrest khalid$ ./grpcrest
Listening on  :9090
Listening on  :50051
```

### Run a GRPC Client
```
MacBook-Air-3:grpcrest khalid$ go run client/grpc/main.go
2018/01/09 20:15:59 Greeting: Hello John Doe
```

### Run a REST Client
```
MacBook-Air-3:grpcrest khalid$ go run client/rest/main.go
```

```
MacBook-Air-3:grpcrest khalid$ ./grpcrest
Listening on  :9090
Listening on  :50051
{"message":"Greeting: Hello John Doe"}
```
