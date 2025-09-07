package requestresponse

import (
	"fmt"
	"net/http"
)

func greetingHandler(resp http.ResponseWriter, req *http.Request) {
	userID := req.PathValue("user")
	fmt.Fprintf(resp, "Hello : %s", userID)
}

func PrintPathVar() {
	http.HandleFunc("/greetings/{user}", greetingHandler)
	fmt.Println("server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
