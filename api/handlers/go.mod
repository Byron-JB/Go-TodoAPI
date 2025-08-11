module todoApi/handlers

go 1.24.5

require (
	go.uber.org/zap v1.27.0
	todoApi/db v0.0.0-00010101000000-000000000000
	todoApi/models v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	gorm.io/driver/mysql v1.6.0 // indirect
	gorm.io/gorm v1.30.1 // indirect
)

replace todoApi/models => ../../internal/models

replace todoApi/db => ../../internal/db
