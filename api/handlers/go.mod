module todoApi/handlers

go 1.24.5

require todoApi/models v0.0.0-00010101000000-000000000000

require (
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)

replace todoApi/models => ../../internal/models
