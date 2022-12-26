package errors

import (
	"fmt"
)

const (
	findOneMessage         = "Expected to find one (and only one) value from %v table"
	unexpectedErrorMessage = "Unexpected error!"
)

type queryErrorType int

const (
	FindOneErr queryErrorType = iota
)

type queryError struct {
	errorType   queryErrorType
	messageArgs []interface{}
}

func NewQueryError(errorType queryErrorType, messageArgs ...interface{}) *queryError {
	return &queryError{
		errorType:   errorType,
		messageArgs: messageArgs,
	}
}

func errorTypeToMessage(errorType queryErrorType) string {
	switch errorType {
	case FindOneErr:
		return findOneMessage
	default:
		return unexpectedErrorMessage
	}
}

func (err *queryError) Error() string {
	return fmt.Sprintf(errorTypeToMessage(err.errorType), err.messageArgs...)
}

func (err *queryError) Is(target error) bool {
	targetQueryError, ok := target.(*queryError)
	if !ok {
		return false
	}

	return err.errorType == targetQueryError.errorType
}
