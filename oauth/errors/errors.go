package errors

import "net/http"

type RESTErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestErr(message string) *RESTErr {
	return &RESTErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundErr(message string) *RESTErr {
	return &RESTErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func InternalServerErr(message string) *RESTErr {
	return &RESTErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
