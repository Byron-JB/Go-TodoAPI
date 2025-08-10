package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"todoApi/models"

	"go.uber.org/zap"
)

func ServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}

// This function digests the request for fetching multiple todo items
func (*TodoFetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := models.SeedTodoData()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
	//io.Copy(w, r.Body);

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	err2 := json.NewEncoder(w).Encode(todos)

	if err2 != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}

}

// This function digests the request for fetching multiple todo items
func (h *TodoCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := io.Copy(w, r.Body)

	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
