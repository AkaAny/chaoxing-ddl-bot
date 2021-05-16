package entity

import (
	"ddl-bot/pkg/chaoxing"
	"time"
)

type Assignment struct {
	AssignmentID string    `gorm:"column:assignment_id;primaryKey"`
	AnswerID     string    `gorm:"column:answer_id;index"`
	CourseID     string    `gorm:"column:course_id;index"`
	ClassID      string    `gorm:"column:class_id;index"`
	Name         string    `gorm:"column:name;index"`
	StartTime    time.Time `gorm:"column:start_time;index"`
	EndTime      time.Time `gorm:"column:end_time;index"`
	Status       string    `gorm:"column:status;index"`
}

func FromSpider(course *chaoxing.Course, assignment *chaoxing.Assignment) *Assignment {
	var item = &Assignment{
		AssignmentID: assignment.AssignmentID,
		AnswerID:     assignment.AnswerID,
		CourseID:     course.CourseID,
		ClassID:      course.ClassID,
		Name:         assignment.Name,
		StartTime:    assignment.StartTime,
		EndTime:      assignment.EndTime,
		Status:       assignment.Status,
	}
	return item
}
