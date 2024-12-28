package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithError(rw http.ResponseWriter, code int, msg string) {
	fmt.Println("ResponsWithError : Message: ", msg)
	RespondWithJson(rw, code, map[string]string{"error": msg})
}

func RespondWithJson(rw http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(response)

}
