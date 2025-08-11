module todoApi/db

go 1.24.5

replace todoApi/models => ../models

require (
	github.com/go-sql-driver/mysql v1.9.3
	todoApi/models v0.0.0-00010101000000-000000000000
)

require filippo.io/edwards25519 v1.1.0 // indirect
