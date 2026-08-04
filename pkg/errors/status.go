package errors

import (
	"errors"
	"fmt"
	"yatter-backend-go/pkg/errors/code"
)

type Status struct {
	code       code.StatusCode
	uiMessage  string
	devMessage string
}

func New(code code.StatusCode, uiMessage string, devMessage string) *Status {
	return &Status{
		code:       code,
		uiMessage:  uiMessage,
		devMessage: devMessage,
	}
}

func FromError(err error) *Status {
	if err == nil {
		return nil
	}

	if v, ok := errors.AsType[*Status](err); ok {
		return v
	}
	return &Status{
		code:      code.Internal,
		uiMessage: err.Error(),
	}
}

func (e *Status) Error() string {
	return fmt.Sprintf("code: %s, uiMessage: %s, devMessage: %v", e.code, e.uiMessage, e.devMessage)
}

func (e *Status) Code() code.StatusCode {
	return e.code
}

func (e *Status) UIMessage() string {
	return e.uiMessage
}

func (e *Status) DevMessage() string {
	return e.devMessage
}

func (e *Status) WithDevMessage(devMessage string) *Status {
	return &Status{
		code:       e.code,
		uiMessage:  e.uiMessage,
		devMessage: devMessage,
	}
}
