module github.com/isOdin-l/TinderArt/services/profile

go 1.26.0

replace github.com/isOdin-l/TinderArt/pkg/configs => ../../pkg/configs

replace github.com/isOdin-l/TinderArt/pkg/grpc => ../../pkg/grpc

replace github.com/isOdin-l/TinderArt/pkg/db_models => ../../pkg/db_models

replace github.com/isOdin-l/TinderArt/pkg/s3 => ../../pkg/s3

replace github.com/isOdin-l/TinderArt/pkg/middleware => ../../pkg/middleware

replace github.com/isOdin-l/TinderArt/pkg/postgres => ../../pkg/postgres

require (
	github.com/caarlos0/env/v11 v11.4.0
	github.com/google/uuid v1.6.0
	github.com/isOdin-l/TinderArt/pkg/configs v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/db_models v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/grpc v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/middleware v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/postgres v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/s3 v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.8.0
	github.com/labstack/echo/v5 v5.0.4
	golang.org/x/crypto v0.49.0
)

require (
	github.com/aws/aws-sdk-go-v2 v1.41.4 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.7 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.20 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.20 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.97.1 // indirect
	github.com/aws/smithy-go v1.24.2 // indirect
	github.com/cridenour/go-postgis v1.0.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	golang.org/x/time v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/grpc v1.79.2 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
