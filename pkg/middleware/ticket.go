package middleware

import (
	"ddl-bot/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	HEADER_TICKET = "ticket"
)

func GetTicket(c *gin.Context) string {
	staffID, exists := c.Get(HEADER_TICKET)
	if !exists {
		panic(errors.New("unexpected call, check if handler is wrapped by UnderGateway"))
	}
	return staffID.(string)
}

func WithResponse(handler ControllerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, ex := handler(c)
		err := response.GetInstance().DoResponse(c, body, ex)
		if err != nil {
			panic(err)
		}
	}
}
