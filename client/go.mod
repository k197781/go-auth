module github.com/k197781/go-auth-client

go 1.12

require (
	github.com/k197781/go-auth-proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)

replace github.com/k197781/go-auth-proto => ../proto/proto
