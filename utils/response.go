package utils

type MOResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Payload    interface{} `json:"payload"`
}

// Responsse construct ManyOption's response payload format
func Response(statusCode int, message interface{}, payload interface{}) (int, *MOResponse) {
	err, ok := message.(error)
	if ok {
		message = err.Error()
	}

	return statusCode, &MOResponse{
		statusCode,
		message,
		payload,
	}
}
