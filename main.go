package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"todoApi/handlers"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {

	fx.New(
		fx.Provide(setupHttpServer), // Provide the SetupHttpServer function
		fx.Provide(
			fx.Annotate(
				handlers.ServeMux,
				fx.ParamTags(`group:"routes"`),
			),
		), // Provide the ServeMux function
		fx.Provide(
			AsRoute(handlers.NewHelloHandler),
			AsRoute(handlers.NewTodoCreateHandler),
			AsRoute(handlers.NewTodoFetchHandler),
			//AsRoute(handlers.NewTodoDeleteHandler),
			//AsRoute(handlers.NewTodoCreateHandler),
			//AsRoute(handlers.NewTodoFetchSingleHandler),
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}), // Invoke the server to start it
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
			go server.Serve(ln)
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
		fx.As(new(handlers.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
