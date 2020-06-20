.PHONY: protoc
protoc:
	protoc proto/*.proto --go_out=plugins=grpc:./proto
