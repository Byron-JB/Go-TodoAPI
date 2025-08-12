package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "todoApi/db"
	"todoApi/models"
)

// TodoUpdateHandler handles updating multiple Todo items via HTTP requests.
// It uses the zap.Logger for logging errors or operations.
type TodoUpdateHandler struct {
	s *Server
}

// NewTodoUpdateHandler Functions to build new handlers
func NewTodoUpdateHandler(s *Server) *TodoUpdateHandler {
	return &TodoUpdateHandler{s: s}
}

// Pattern Defines the route of the endpoint
func (*TodoUpdateHandler) Pattern() string {
	return "/update"
}

// This function digests the request for fetching multiple todo items
func (h *TodoUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Only PATCH method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the server is initialized
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

	updatedTodos, err := h.s.database.UpdateTodosInDB(todos)

	if err != nil {
		if err.Error() != "" {
			http.Error(w, fmt.Sprintf(err.Error()), http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	// Respond back with received todos
	w.Header().Set("Content-Type", "application/json")
	encodingError := json.NewEncoder(w).Encode(updatedTodos)

	if encodingError != nil {
		http.Error(w, fmt.Sprintf("DB insert error: %v", err), http.StatusInternalServerError)
	}

	return
}
