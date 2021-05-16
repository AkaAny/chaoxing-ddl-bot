package middleware

import "github.com/gin-gonic/gin"

type Data interface{}
type ControllerFunc func(c *gin.Context) (Data, error)
