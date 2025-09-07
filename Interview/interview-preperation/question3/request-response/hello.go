package requestresponse

import (
	"fmt"
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello")
}
func PrintHello() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
