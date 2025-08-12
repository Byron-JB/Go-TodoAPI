package handlers

import (
	"net/http"
	"todoApi/db"
)

type Server struct {
	database db.Database
}

func NewServer(db db.Database) *Server {
	return &Server{database: db}
}

// The Route interface will handle HTTP requests and define a pattern for the route.
type Route interface {
	http.Handler
	Pattern() string
}

// ServeMux creates and returns an HTTP ServeMux configured with the provided routes.
func ServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}
