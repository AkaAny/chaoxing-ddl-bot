package response

import (
	"ddl-bot/pkg/factory"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var sResponseFactory *ResponseFactory = nil

func GetInstance() *ResponseFactory {
	if sResponseFactory == nil {
		sResponseFactory = new(ResponseFactory)
	}
	return sResponseFactory
}

type ResponseFactory struct {
	Factory factory.Factorize
}

func (f *ResponseFactory) Put(creators ...factory.Creator) error {
	return f.Factory.Put(creators...)
}

func (f *ResponseFactory) DoCreateResponse(data interface{}, ex error) (*ResponseEntity, error) {
	if ex == nil {
		return &ResponseEntity{
			StatusCode: http.StatusOK,
			Body:       MakeSuccess(data),
		}, nil
	}
	conf, ok := ex.(factory.Config)
	if !ok {
		return nil, errors.New("invalid ex input")
	}
	obj, err := f.Factory.DoCreate(conf)
	if err != nil {
		return nil, err
	}
	res, ok := obj.(*ResponseEntity)
	if !ok {
		return nil, errors.New("obj must be a pointer of ResponseEntity")
	}
	if data != nil {
		res.Body.Data = data
	}
	return res, nil
}

func (f *ResponseFactory) DoResponse(c *gin.Context, data interface{}, ex error) error {
	respObj, err := f.DoCreateResponse(data, ex)
	if err != nil {
		return err
	}
	c.JSON(respObj.StatusCode, respObj.Body)
	return nil
}
