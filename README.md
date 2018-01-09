# grpcrest
GRPC and Rest testing

This is a test case for using GRPC as a great alternative for REST


### Running Benchmark
```
MMDCs-MacBook-Air-3:grpcrest mmdc$ go test ./... -bench Bench
BenchmarkGRPC-4   	 1000000	      1105 ns/op
BenchmarkREST-4   	    2000	  15518352 ns/op
```

### Running a standalone server grpcrest
```
MMDCs-MacBook-Air-3:grpcrest mmdc$ go build
MMDCs-MacBook-Air-3:grpcrest mmdc$ ./grpcrest
Listening on  :9090
Listening on  :50051
```

### Run a GRPC Client
```
MMDCs-MacBook-Air-3:grpcrest mmdc$ go run client/grpc/main.go
2018/01/09 20:15:59 Greeting: Hello John Doe
```

### Run a REST Client
```
MMDCs-MacBook-Air-3:grpcrest mmdc$ go run client/rest/main.go
```

```
MMDCs-MacBook-Air-3:grpcrest mmdc$ ./grpcrest
Listening on  :9090
Listening on  :50051
{"message":"Greeting: Hello John Doe"}
```
