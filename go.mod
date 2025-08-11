module todoApi/main

go 1.24.5

replace todoApi/handlers => ./api/handlers

replace todoApi/models => ./internal/models

replace todoApi/db => ./internal/db

require (
	go.uber.org/fx v1.24.0
	go.uber.org/zap v1.27.0
	todoApi/db v0.0.0-00010101000000-000000000000
	todoApi/handlers v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.uber.org/dig v1.19.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	gorm.io/driver/mysql v1.6.0 // indirect
	gorm.io/gorm v1.30.1 // indirect
	todoApi/models v0.0.0-00010101000000-000000000000 // indirect
)
