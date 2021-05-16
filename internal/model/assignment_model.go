package model

import (
	"ddl-bot/internal/entity"
	"time"
)

type AssignmentResponseItem struct {
	AssignmentID string    `json:"assignmentID"`
	AnswerID     string    `json:"answerID"`
	CourseID     string    `json:"courseID"`
	ClassID      string    `json:"classID"`
	Name         string    `json:"name"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Status       string    `json:"status"`
}

func FromEntity(item *entity.Assignment) *AssignmentResponseItem {
	var data = &AssignmentResponseItem{
		AssignmentID: item.AssignmentID,
		AnswerID:     item.AnswerID,
		CourseID:     item.CourseID,
		ClassID:      item.ClassID,
		Name:         item.Name,
		StartTime:    item.StartTime,
		EndTime:      item.EndTime,
		Status:       item.Status,
	}
	return data
}

type RefreshResponse struct {
	ErrorMap map[string]error `json:"error_map"`
}
