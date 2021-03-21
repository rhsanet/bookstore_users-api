package response

import "net/http"

type RestErr struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func BadRequestResponse(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "bad_request",
		StatusCode: http.StatusBadRequest,
	}
}
