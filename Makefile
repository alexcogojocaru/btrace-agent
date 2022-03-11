generate-proto:
	@protoc -I=btrace-idl/proto --go_out=proto-gen .\btrace-idl\proto\agent.proto
	@protoc -I=btrace-idl/proto --go-grpc_out=proto-gen .\btrace-idl\proto\agent.proto

build:
	@go install

run:
	@btrace-agent.exe