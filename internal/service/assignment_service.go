package service

import (
	"ddl-bot/internal/config"
	"ddl-bot/internal/entity"
	"ddl-bot/internal/mapper"
	"ddl-bot/internal/model"
	"ddl-bot/pkg/cas"
	"ddl-bot/pkg/chaoxing"
	"fmt"
	"sync"
)

type AssignmentService struct {
	Config *config.Config
	Mapper *mapper.AssignmentMapper
}

func (s *AssignmentService) Refresh() (*model.RefreshResponse, error) {
	request, err := cas.Login(s.Config.CAS.UserName, s.Config.CAS.Password)
	if err != nil {
		return nil, err
	}
	cx, err := chaoxing.Login(request)
	if err != nil {
		return nil, err
	}
	courses, err := cx.GetCourseList()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(courses))
	var errMap = make(map[string]error)
	for _, course := range courses {
		var courseLocal = course
		go func() {
			fmt.Println(courseLocal.Title, courseLocal.Teacher)
			var assignments []chaoxing.Assignment
			assignments, err = cx.GetAssignmentList(&courseLocal)
			if err != nil {
				errMap[courseLocal.Title] = err
				return
			}
			fmt.Println(assignments)
			for _, assignment := range assignments {
				var item = entity.FromSpider(&courseLocal, &assignment)
				err := s.Mapper.Save(s.Mapper.GetDB(), item)
				if err != nil {
					errMap[courseLocal.Title] = err
				}
			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("finished")
	return &model.RefreshResponse{ErrorMap: errMap}, nil
}

func (s *AssignmentService) List() ([]model.AssignmentResponseItem, error) {
	items, err := s.Mapper.List(s.Mapper.GetDB())
	if err != nil {
		return nil, err
	}
	return s.wrapForResponse(items), nil
}

func (s *AssignmentService) ListToDo() ([]model.AssignmentResponseItem, error) {
	items, err := s.Mapper.ListToDo(s.Mapper.GetDB())
	if err != nil {
		return nil, err
	}
	return s.wrapForResponse(items), nil
}

func (s *AssignmentService) wrapForResponse(items []entity.Assignment) []model.AssignmentResponseItem {
	var dataList []model.AssignmentResponseItem
	for _, item := range items {
		var data = model.FromEntity(&item)
		dataList = append(dataList, *data)
	}
	return dataList
}
