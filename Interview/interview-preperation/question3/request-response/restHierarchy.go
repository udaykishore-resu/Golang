package requestresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getUserOrderDetails(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	parts := strings.SplitN(path, "/", 2)

	if len(parts) != 2 || parts[1] != "orders" {
		http.NotFound(w, r)
		return
	}

	userID := parts[0]

	response := map[string]interface{}{
		"user_id": userID,
		"orders":  []string{"order123", "order456"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func PrintRESTHierarchy() {
	http.HandleFunc("/users/", getUserOrderDetails)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
