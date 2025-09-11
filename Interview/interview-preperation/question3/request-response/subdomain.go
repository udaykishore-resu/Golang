package requestresponse

import (
	"fmt"
	"net/http"
	"strings"
)

// http://api.customer1.localhost:8080/

func getSubdomain(resp http.ResponseWriter, req *http.Request) {
	subdomainHandlers := map[string]string{
		"customer1": "Customer 1 API",
		"customer2": "Customer 2 API",
		"admin":     "Admin API",
	}

	host := req.Host
	subdomain := strings.Split(host, ".")[1]
	if msg, ok := subdomainHandlers[subdomain]; ok {
		fmt.Fprintln(resp, msg)
	} else {
		fmt.Fprintln(resp, "Default API")
	}
}

func PrintSubdomain() {
	http.HandleFunc("/", getSubdomain)
	http.ListenAndServe(":8080", nil)
}
