package requestresponse

import (
	"fmt"
	"net/http"
	"strings"
)

func getProdDetails(resp http.ResponseWriter, req *http.Request) {
	path := req.PathValue("path")
	params := make(map[string]string)
	pathVars := strings.Split(path, ";")
	for _, varb := range pathVars {
		v := strings.Split(varb, "=")
		params[v[0]] = v[1]
	}

	fmt.Fprintf(resp, "The matrix params are: \n")
	for key, val := range params {
		fmt.Fprintf(resp, "%s = %s \n", key, val)
	}
}
func PrintMatrixParam() {
	http.HandleFunc("/prod-details/{path}", getProdDetails)
	http.ListenAndServe(":8080", nil)
}
