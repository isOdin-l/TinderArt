module github.com/isOdin-l/TinderArt/pkg/middleware

replace github.com/isOdin-l/TinderArt/pkg/grpc/auth => ../grpc

replace github.com/isOdin-l/TinderArt/pkg/configs => ../configs

go 1.26.0

require (
	github.com/google/uuid v1.6.0
	github.com/isOdin-l/TinderArt/pkg/grpc/auth v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v5 v5.0.4
)

require (
	github.com/isOdin-l/TinderArt/pkg/configs v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/grpc v1.79.2 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
