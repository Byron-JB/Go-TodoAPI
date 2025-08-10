package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

// The Route interface will handle HTTP requests and define a pattern for the route.
type Route interface {
	http.Handler
	Pattern() string
}

func (*TodoFetchHandler) Pattern() string {
	return "/todos"
}

func (*TodoCreateHandler) Pattern() string {
	return "/create"
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

// Handler structs

// This handler will fetch the todo list.
type TodoFetchHandler struct {
	log *zap.Logger
}

type TodoCreateHandler struct {
	log *zap.Logger
}

type TodoUpdateHandler struct {
	log *zap.Logger
}

type TodoDeleteHandler struct {
	log *zap.Logger
}

type TodoFetchSingleHandler struct {
	log *zap.Logger
}

// Functions to build new handlers
func NewTodoFetchHandler(log *zap.Logger) *TodoFetchHandler {
	return &TodoFetchHandler{log: log}
}

func NewTodoCreateHandler(log *zap.Logger) *TodoCreateHandler {
	return &TodoCreateHandler{log: log}
}

func NewTodoUpdateHandler(log *zap.Logger) *TodoUpdateHandler {
	return &TodoUpdateHandler{log: log}
}

func NewTodoDeleteHandler(log *zap.Logger) *TodoDeleteHandler {
	return &TodoDeleteHandler{log: log}
}

func NewTodoFetchSingleHandler(log *zap.Logger) *TodoFetchSingleHandler {
	return &TodoFetchSingleHandler{log: log}
}

// HelloHandler is an HTTP handler that
// prints a greeting to the user.
type HelloHandler struct {
	log *zap.Logger
}

// NewHelloHandler builds a new HelloHandler.
func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}
