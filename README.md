grpc-test
===

protoc
```
protoc --go_out=pkg/time --go_opt=paths=source_relative --go-grpc_out=pkg/time --go-grpc_opt=paths=source_relative time.proto
```

server
```
go run cmd/server/main.go
```

grpcurl
```
grpcurl -plaintext localhost:50051 grpctest.time.v1.TimeService.GetCurrentTime
```

go client
```
go run cmd/client/main.go
```

---
```
grpc_tools_ruby_protoc  --ruby_out=ruby --grpc_out=ruby time.proto 
```

ruby client
```
ruby ruby/client.rb
```
