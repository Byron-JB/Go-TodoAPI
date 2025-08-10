package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"todoApi/models"
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
