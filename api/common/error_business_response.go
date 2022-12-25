package common

import (
	"net/http"

	"github.com/baguseka01/golang_microservice_hexagonal/business"
)

type errorBusinessRessponseCode string

const (
	errInternalServerError errorBusinessRessponseCode = "internal_server_error"
	errHasBeenModified     errorBusinessRessponseCode = "data_has_been_modified"
	errNotFound            errorBusinessRessponseCode = "data_not_found"
	errInvalidSpec         errorBusinessRessponseCode = "invalid_spec"
)

// BusinessResponse default payload response
type BusinessResponse struct {
	Code    errorBusinessRessponseCode `json:"code"`
	Message string                     `json:"message"`
	Data    interface{}                `json:"data"`
}

// NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, BusinessResponse) {
	return errorMapping(err)
}

// errorMapping error for missing header key with given value
func errorMapping(err error) (int, BusinessResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()
	case business.ErrNotFound:
		return newNotFoundResponse()
	case business.ErrInvalidSpec:
		return newValidationResponse(err.Error())
	case business.ErrHasBeenModified:
		return newHasBeenModifiedResponse()
	}
}

// newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse() (int, BusinessResponse) {
	return http.StatusInternalServerError, BusinessResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}

// newNotFoundResponse default not found error response
func newNotFoundResponse() (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

// newValidationResponse failed to validate request payload
func newValidationResponse(message string) (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errInvalidSpec,
		"Validation failed" + message,
		map[string]interface{}{},
	}
}

// newHasBeenModifiedResponse failed to validate request payload
func newHasBeenModifiedResponse() (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errHasBeenModified,
		"Data has been modified",
		map[string]interface{}{},
	}
}