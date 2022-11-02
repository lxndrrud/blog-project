package models

import "encoding/json"

type IError interface {
	GetMessage() string
	GetCode() int
}

func NewError(code int, message string) IError {
	return &customError{
		message: message,
		code:    code,
	}
}

type customError struct {
	message string
	code    int
}

func (c *customError) GetMessage() string {
	return c.message
}

func (c *customError) GetCode() int {
	return c.code
}

func (c *customError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"message": c.message,
	})
}
