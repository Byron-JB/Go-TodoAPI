package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todoApi/models"
)

// TodoCreateHandler Creates the Handler for the endpoint
type TodoCreateHandler struct {
	s *Server
}

// NewTodoCreateHandler Functions to build new handlers
func NewTodoCreateHandler(s *Server) *TodoCreateHandler {
	return &TodoCreateHandler{s: s}
}

// Pattern Defines the route of the endpoint
func (*TodoCreateHandler) Pattern() string {
	return "/create"
}

// This function digests the request for fetching multiple todo items
func (h *TodoCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	if h.s == nil || h.s.database == nil {
		http.Error(w, "Server not initialized", http.StatusInternalServerError)
		return
	}

	var todos []models.TodoDto

	// Decode JSON body into todos slice
	if err := json.NewDecoder(r.Body).Decode(&todos); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(r.Body)

	savedTodos, err := h.s.database.SaveTodosToDb(todos)

	if err != nil {
		if err.Error() != "" {
			http.Error(w, fmt.Sprintf(err.Error()), http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	// Respond back with received todos
	w.Header().Set("Content-Type", "application/json")
	encodingError := json.NewEncoder(w).Encode(savedTodos)

	if encodingError != nil {
		http.Error(w, fmt.Sprintf("DB insert error: %v", err), http.StatusInternalServerError)
	}

	return
}
