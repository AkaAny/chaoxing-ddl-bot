package controller

import (
	"ddl-bot/internal/service"
	"ddl-bot/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type AssignmentController struct {
	AssignmentService *service.AssignmentService
}

func (controller *AssignmentController) Refresh(c *gin.Context) (middleware.Data, error) {
	data, err := controller.AssignmentService.Refresh()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (controller *AssignmentController) List(c *gin.Context) (middleware.Data, error) {
	return controller.AssignmentService.List()
}

func (controller *AssignmentController) ListToDo(c *gin.Context) (middleware.Data, error) {
	return controller.AssignmentService.ListToDo()
}
