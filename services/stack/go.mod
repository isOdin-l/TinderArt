module github.com/isOdin-l/TinderArt/services/stack

go 1.26.0

replace github.com/isOdin-l/TinderArt/pkg/postgres => ../../pkg/postgres

replace github.com/isOdin-l/TinderArt/pkg/configs => ../../pkg/configs

replace github.com/isOdin-l/TinderArt/pkg/grpc => ../../pkg/grpc

replace github.com/isOdin-l/TinderArt/pkg/db_models => ../../pkg/db_models

replace github.com/isOdin-l/TinderArt/pkg/redis => ../../pkg/redis

require (
	github.com/caarlos0/env/v11 v11.4.0
	github.com/isOdin-l/TinderArt/pkg/configs v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/db_models v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/postgres v0.0.0-00010101000000-000000000000
	github.com/isOdin-l/TinderArt/pkg/redis v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.8.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cridenour/go-postgis v1.0.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/redis/go-redis/v9 v9.18.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.35.0 // indirect
)
