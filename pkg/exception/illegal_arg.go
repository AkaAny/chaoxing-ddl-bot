package exception

import (
	"ddl-bot/pkg/factory"
	"ddl-bot/pkg/response"
	"net/http"
)

func init() {
	err := response.GetInstance().Put(IllegalArgumentException{})
	if err != nil {
		panic(err)
	}
}

type IllegalArgumentException struct {
	ArgNames []string `json:"arg_names"`
	msg      string
}

func NewIllegalArgumentException(msg string, argNames ...string) IllegalArgumentException {
	return IllegalArgumentException{ArgNames: argNames, msg: msg}
}

func (e IllegalArgumentException) Error() string {
	return e.msg
}

func (IllegalArgumentException) GetType() string {
	return "bad_request"
}

func (IllegalArgumentException) Create(exObj factory.Config) (factory.Object, error) {
	ex, ok := exObj.(IllegalArgumentException)
	if !ok {
		return nil, NewIllegalArgumentException("invalid ex input", "exObj")
	}
	return &response.ResponseEntity{
		StatusCode: http.StatusBadRequest,
		Body: response.BaseResponse{
			Error: 40000,
			Msg:   ex.Error(),
			Data:  ex,
		},
	}, nil
}
