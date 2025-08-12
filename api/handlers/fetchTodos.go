package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todoApi/models"
)

// TodoFetchHandler This handler will fetch the todo list.
type TodoFetchHandler struct {
	s *Server
}

// NewTodoFetchHandler Functions to build new handlers
func NewTodoFetchHandler(s *Server) *TodoFetchHandler {
	return &TodoFetchHandler{s: s}
}

// Pattern Defines the route of the endpoint
func (*TodoFetchHandler) Pattern() string {
	return "/todos"
}

// This function digests the request for fetching multiple todo items
func (h *TodoFetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the server is initialized
	if h.s == nil || h.s.database == nil {
		http.Error(w, "Server not initialized", http.StatusInternalServerError)
		return
	}

	// Get the host url
	hostUrl := r.Host

	// Get the skip value from the query string
	skipStr := r.URL.Query().Get("skip")
	skip, err := strconv.Atoi(skipStr)
	// If the skip value is not a number, set it to 0
	if err != nil {
		skip = 0 // your fallback value
	}

	// Get the take value from the query string
	takeStr := r.URL.Query().Get("take")
	take, err := strconv.Atoi(takeStr)

	// If the take value is not a number, set it to 10
	if err != nil {
		take = 10
	}

	// Calculate the next page skip value
	nextPageSkip := skip + take

	// Build the next page url
	nextPageUrl := hostUrl + "/todos?skip=" + strconv.Itoa(nextPageSkip) + "&take" + strconv.Itoa(take)

	metaData := models.MetaData{
		Skip:     skip,
		Take:     take,
		NextPage: nextPageUrl,
	}

	todosFromDb, err := h.s.database.FetchTodosFromDb(skip, take)

	if err != nil {
		if err.Error() != "" {
			http.Error(w, fmt.Sprintf(err.Error()), http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	response := models.FetchTodoResponseDTO{
		MetaData: metaData,
		Todos:    todosFromDb,
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Encode the response and return it
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	return
}
