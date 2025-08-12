package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

// Pattern Defines the route of the endpoint
func (*TodoDeleteHandler) Pattern() string {
	return "/delete"
}

// TodoDeleteHandler Creates the Handler for the endpoint
type TodoDeleteHandler struct {
	s *Server
}

// NewTodoDeleteHandler Functions to build new handlers
func NewTodoDeleteHandler(s *Server) *TodoDeleteHandler {
	return &TodoDeleteHandler{s: s}
}

// This function digests the request
func (h *TodoDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method allowed", http.StatusMethodNotAllowed)
		return
	}

	if h.s == nil || h.s.database == nil {
		http.Error(w, "Server not initialized", http.StatusInternalServerError)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.s.database.DeleteTodoFromDb(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Write the success message to the response body
	_, err = fmt.Fprintf(w, "todo deleted successfully")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	return
}
