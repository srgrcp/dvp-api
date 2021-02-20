package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

// IsValid validates a struct's fields
func IsValid(item interface{}) bool {
	validator := validator.New()
	err := validator.Struct(item)
	return err == nil
}

// Response sends a json response
func Response(w http.ResponseWriter, err error, data interface{}) {
	var errString *string = nil
	if err != nil {
		_errString := fmt.Sprint(err)
		errString = &_errString
	}
	response := APIResponse{
		Error: errString,
		Data:  data,
	}
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		panic(err2)
	}
}

// HeadersMiddleware cors and content-type headers
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
