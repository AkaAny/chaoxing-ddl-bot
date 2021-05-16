package internal

import (
	"ddl-bot/internal/config"
	"ddl-bot/internal/controller"
	"ddl-bot/internal/mapper"
	"ddl-bot/internal/service"
	"ddl-bot/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type BotApplication struct {
	Config            *config.Config
	AssignmentService *service.AssignmentService
}

func Create(conf *config.Config) *BotApplication {
	db, err := gorm.Open(mysql.Open(conf.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var assignmentMapper = mapper.AssignmentMapper{DB: db}
	err = assignmentMapper.Init()
	if err != nil {
		panic(err)
	}
	var assignmentService = service.AssignmentService{
		Config: conf,
		Mapper: &assignmentMapper,
	}
	return &BotApplication{
		Config:            conf,
		AssignmentService: &assignmentService,
	}
}

func (app *BotApplication) Run() {
	var engine = gin.Default()
	var assignmentController = controller.AssignmentController{AssignmentService: app.AssignmentService}
	assignmentGroup := engine.Group("/assignment")
	{
		assignmentGroup.Handle(http.MethodPut, "/refresh",
			middleware.WithResponse(assignmentController.Refresh))
		assignmentGroup.Handle(http.MethodGet, "/list",
			middleware.WithResponse(assignmentController.List))
		assignmentGroup.Handle(http.MethodGet, "/list/todo",
			middleware.WithResponse(assignmentController.ListToDo))
	}
	err := engine.Run(fmt.Sprintf(":%d", app.Config.Port))
	if err != nil {
		panic(err)
	}
}
