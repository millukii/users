package errors

import "errors"

var ErrResourceNotFound = errors.New("resource not found")
var ErrInvalidPayload = errors.New("invalid payload")
var ErrServiceUnavailable = errors.New("the service is unavailable")
var ErrUnauthenticated = errors.New("unauthenticated")
var ErrUnauthorized = errors.New("unauthorized")