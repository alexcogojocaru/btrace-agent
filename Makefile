build:
	@go install

run:
	@btrace-agent.exe

generate-proto:
	@protoc -I=btrace-idl/proto --go_out=proto-gen .\btrace-idl\proto\agent.proto
	@protoc -I=btrace-idl/proto --go-grpc_out=proto-gen .\btrace-idl\proto\agent.proto

	@protoc -I=btrace-idl/proto --go_out=proto-gen .\btrace-idl\proto\v2\proxy.proto
	@protoc -I=btrace-idl/proto --go-grpc_out=proto-gen .\btrace-idl\proto\v2\proxy.proto

proto-collector:
	protoc -I=btrace-idl/proto --go_out=proto-gen ./btrace-idl/proto/collector.proto
	protoc -I=btrace-idl/proto --go-grpc_out=proto-gen ./btrace-idl/proto/collector.proto
