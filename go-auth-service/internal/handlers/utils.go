package handlers

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
