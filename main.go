package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"todoApi/db"
	"todoApi/handlers"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {

	fx.New(
		fx.Module("db",
			fx.Provide(
				fx.Annotate(
					db.NewDatabase,
					fx.As(new(db.Database)),
				),
			),
		),
		fx.Provide(setupHttpServer), // Provide the SetupHttpServer function
		fx.Provide(
			handlers.NewServer,
			fx.Annotate(
				handlers.ServeMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		// Provide the ServeMux function
		fx.Provide(
			AsRoute(handlers.NewTodoCreateHandler),
			AsRoute(handlers.NewTodoFetchHandler),
			AsRoute(handlers.NewTodoDeleteHandler),
			AsRoute(handlers.NewTodoUpdateHandler),
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}), // Invoke the server to start it
		//fx.Invoke(db.InitDB),
	).Run()
}

// setupHttpServer creates a new HTTP handler for the Todo API.
func setupHttpServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	server := &http.Server{
		Addr: ":8080", Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server on", server.Addr)
			go func() {
				err := server.Serve(ln)
				if err != nil {

				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server...")
			return server.Shutdown(ctx)
		},
	})

	return server
}

// Annotates the given contructor to state that it provides and API route
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(
			new(handlers.Route),
		),
		fx.ResultTags(`group:"routes"`),
	)
}
