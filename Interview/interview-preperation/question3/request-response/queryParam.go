package requestresponse

import (
	"fmt"
	"net/http"
)

func searchUserHandler(resp http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	age := query.Get("age")
	fmt.Fprintf(resp, "Search User with Name: %s and Age: %s", name, age)
}
func PrintQueryParam() {
	http.HandleFunc("/search", searchUserHandler)
	http.ListenAndServe(":8080", nil)
}
