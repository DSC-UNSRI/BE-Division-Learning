package utils

import (
    "encoding/json"
    "net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if data != nil {
        json.NewEncoder(w).Encode(data)
    }
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    JSONResponse(w, statusCode, map[string]string{"error": message})
}