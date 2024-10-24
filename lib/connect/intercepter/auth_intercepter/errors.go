package auth_intercepter

import (
	"errors"
	"fmt"
)

type AuthenticationError struct {
	cause error
}

func (e *AuthenticationError) Error() string {
	if e.cause == nil {
		return "Authentication error"
	}
	return fmt.Sprintf("Authentication error: %v", e.cause)
}

func IsAuthenticationError(err error) bool {
	if err == nil {
		return false
	}
	var e *AuthenticationError
	return errors.As(err, &e)
}

type PermissionError struct {
	cause error
}

func (e *PermissionError) Error() string {
	if e.cause == nil {
		return "Permission error"
	}
	return fmt.Sprintf("Permission error: %v", e.cause)
}

func IsPermissionError(err error) bool {
	if err == nil {
		return false
	}
	var e *PermissionError
	return errors.As(err, &e)
}
