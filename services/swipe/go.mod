module github.com/isOdin-l/TinderArt/services/swipe

go 1.26.0

replace github.com/isOdin-l/TinderArt/pkg/kafka => ../../pkg/kafka

replace github.com/isOdin-l/TinderArt/pkg/postgres => ../../pkg/postgres

replace github.com/isOdin-l/TinderArt/pkg/configs => ../../pkg/configs

require (
	github.com/caarlos0/env/v11 v11.4.0
	github.com/google/uuid v1.6.0
	github.com/isOdin-l/TinderArt/pkg/configs v0.0.0
	github.com/isOdin-l/TinderArt/pkg/kafka v0.0.0
	github.com/isOdin-l/TinderArt/pkg/postgres v0.0.0
	github.com/jackc/pgx/v5 v5.8.0
	github.com/labstack/echo/v5 v5.0.4
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.50 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	golang.org/x/time v0.14.0 // indirect
)
