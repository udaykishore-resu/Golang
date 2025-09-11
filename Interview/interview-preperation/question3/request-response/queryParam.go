package requestresponse

import (
	"fmt"
	"net/http"
)

func searchHandler(resp http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	keyword := query.Get("keyword")
	page := query.Get("page")
	fmt.Fprintf(resp, "Search with keyword: %s and page: %s", keyword, page)
}
func PrintQueryParam() {
	http.HandleFunc("/search", searchHandler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
