package exception

import (
	"ddl-bot/pkg/factory"
	"ddl-bot/pkg/response"
	"net/http"
)

func init() {
	err := response.GetInstance().Put(InternalException{})
	if err != nil {
		panic(err)
	}
}

type InternalException struct {
	Type string `json:"type"`
	err  error
	msg  string `json:"-"`
}

func NewInternalException(msg string, exType string, err error) InternalException {
	return InternalException{
		Type: exType,
		err:  err,
		msg:  msg,
	}
}

func (e InternalException) Error() string {
	return e.msg
}

func (InternalException) GetType() string {
	return "internal_server_error"
}

func (InternalException) Create(exObj factory.Config) (factory.Object, error) {
	ex, ok := exObj.(InternalException)
	if !ok {
		return nil, NewIllegalArgumentException("invalid ex input", "exObj")
	}
	return &response.ResponseEntity{
		StatusCode: http.StatusInternalServerError,
		Body: response.BaseResponse{
			Error: 50000,
			Msg:   ex.Error(),
			Data:  ex,
		},
	}, nil
}
