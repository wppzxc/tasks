package errors

import (
	"fmt"
)

type TaskError struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func (e TaskError) Error() string {
	return fmt.Sprintf("%#v", e)
}

func NewTaskServerError(code int, err error) (int, TaskError) {
	return code, TaskError{
		Code: code,
		Msg:  err.Error(),
	}
}