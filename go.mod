module todoApi/main

go 1.24.5

replace todoApi/handlers => ../goTodoAPI/api/handlers

replace todoApi/models => ../goTodoAPI/internal/models

require (
	go.uber.org/fx v1.24.0
	todoApi/handlers v0.0.0-00010101000000-000000000000
)

require (
	go.uber.org/dig v1.19.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	todoApi/models v0.0.0-00010101000000-000000000000 // indirect
)
