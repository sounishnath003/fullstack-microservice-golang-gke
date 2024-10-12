package utils

import (
	"log"
	"os"
)

// GetEnv function to get environment variable.
// If the environment variable is not found, it will return the fallback value.
func GetEnv(key string, fallback any) any {
	if val, ok := os.LookupEnv(key); ok {
		log.Printf("key %s environment value found\n", key)
		return val
	}
	log.Printf("no environment value found. setting fallback value key=%s\n", key)
	return fallback
}

// ApiResponse struct to handle response data.
type ApiResponse struct {
	Data   any    `json:"data"`
	Status int    `json:"statusCode"`
	Error  string `json:"error,omitempty"`
}

// Create a new ApiResponse object.
func NewApiResponse(status int, data any, err error) *ApiResponse {
	if err != nil {
		return &ApiResponse{
			Data:   data,
			Status: status,
			Error:  err.Error(),
		}
	}

	return &ApiResponse{
		Data:   data,
		Status: status,
	}
}
