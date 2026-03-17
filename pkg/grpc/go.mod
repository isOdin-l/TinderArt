module github.com/isOdin-l/TinderArt/pkg/grpc

go 1.26.0

replace github.com/isOdin-l/TinderArt/pkg/configs => ../configs

require (
	github.com/isOdin-l/TinderArt/pkg/configs v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.79.2
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)
