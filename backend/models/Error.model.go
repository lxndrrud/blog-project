package models

import "encoding/json"

type IError interface {
	GetMessage() string
	GetCode() int
}

func NewError(code int, message string) IError {
	return &CustomError{
		message: message,
		code:    code,
	}
}

type CustomError struct {
	message string
	code    int
}

func (c *CustomError) GetMessage() string {
	return c.message
}

func (c *CustomError) GetCode() int {
	return c.code
}

func (c *CustomError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"message": c.message,
	})
}
