package exception

import (
	"ddl-bot/pkg/factory"
	"ddl-bot/pkg/response"
	"net/http"
)

func init() {
	err := response.GetInstance().Put(IllegalAccessException{})
	if err != nil {
		panic(err)
	}
}

type IllegalAccessException struct {
	msg string
}

func NewIllegalAccessException(msg string) IllegalAccessException {
	return IllegalAccessException{msg: msg}
}

func (e IllegalAccessException) Error() string {
	return e.msg
}

func (IllegalAccessException) GetType() string {
	return "forbidden"
}

func (IllegalAccessException) Create(exObj factory.Config) (factory.Object, error) {
	ex, ok := exObj.(IllegalAccessException)
	if !ok {
		return nil, NewIllegalArgumentException("invalid ex input", "exObj")
	}
	return &response.ResponseEntity{
		StatusCode: http.StatusForbidden,
		Body: response.BaseResponse{
			Error: 40300,
			Msg:   ex.Error(),
			Data:  nil,
		},
	}, nil
}
