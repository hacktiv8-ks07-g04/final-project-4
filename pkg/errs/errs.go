package errs

import (
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"Error"`
}

func NotFound(message string) Error {
	return Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "NOT_FOUND",
	}
}

func BadRequest(message string) Error {
	return Error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BAD_REQUEST",
	}
}

func InternalServerError(message string) Error {
	return Error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_Error",
	}
}

func Unauthorized(message string) Error {
	return Error{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "UNAUTHORIZED",
	}
}

func UnprocessableEntity(message string) Error {
	return Error{
		Message: message,
		Status:  http.StatusUnprocessableEntity,
		Error:   "UNPROCESSABLE_ENTITY",
	}
}
