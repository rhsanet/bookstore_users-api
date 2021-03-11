package response

import "net/http"

type ErrorMessage struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func BadRequestResponse(message string, err error) *ErrorMessage {
	return &ErrorMessage{
		Message:    message,
		Error:      err.Error(),
		StatusCode: http.StatusBadRequest,
	}
}
