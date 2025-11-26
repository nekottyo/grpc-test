grpc-test
===

## Requires
- Go
- protoc
- grpcurl


## Hands on
### basic
現在時刻を返す service の実装
1. pb 用意
1. protoc で変換
1. server 実装
1. grpcurl で動作確認
1. client 実装

### advanced
現在時刻を聞くと 1秒ごとに 5 回 stream で返す serivce の実装



## Usage

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
go run cmd/client/main.go stream
```

---
```
grpc_tools_ruby_protoc  --ruby_out=ruby --grpc_out=ruby time.proto 
```

ruby client
```
ruby ruby/client.rb
```
