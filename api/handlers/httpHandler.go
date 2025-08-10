package handlers

import (
	"net/http"
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
	return "todo/create"
}

// Handler structs

// This handler will fetch the todo list.
type TodoFetchHandler struct{}

type TodoCreateHandler struct{}

type TodoUpdateHandler struct{}

type TodoDeleteHandler struct{}

type TodoFetchSingleHandler struct{}

// Functions to build new handlers
func NewTodoFetchHandler() *TodoFetchHandler {
	return &TodoFetchHandler{}
}

func NewTodoCreateHandler() *TodoCreateHandler {
	return &TodoCreateHandler{}
}

func NewTodoUpdateHandler() *TodoUpdateHandler {
	return &TodoUpdateHandler{}
}

func NewTodoDeleteHandler() *TodoDeleteHandler {
	return &TodoDeleteHandler{}
}

func NewTodoFetchSingleHandler() *TodoFetchSingleHandler {

}
