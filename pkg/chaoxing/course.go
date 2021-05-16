package chaoxing

import (
	"ddl-bot/pkg/cas"
	"ddl-bot/pkg/superagent"
	"ddl-bot/pkg/wrapper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"strings"
)

type Course struct {
	CourseID string `gorm:"column:course_id;primaryKey"`
	ClassID  string `gorm:"column:class_id;index"`
	Title    string `gorm:"column:title"`
	Teacher  string `gorm:"column:teacher;index"`
	URL      string `gorm:"column:url"`
}

type CourseList []Course

func (cx *ChaoXing) GetCourseList() (CourseList, error) {
	var request = cx.GetRequest()
	var url = "http://mooc1-1.chaoxing.com/visit/courses/study?isAjax=true&fileId=0&debug=false"
	request.Get(url)
	superagent.WithUA(request, cas.USER_AGENT)
	resp, body, errs := request.End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err := ioutil.WriteFile(wrapper.GetPath("resp_course.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var courses []Course
	doc.Find(".courseItem").Each(func(iCourse int, courseNode *goquery.Selection) {
		var course Course
		courseNode.ChildrenFiltered("input").Each(func(iCourseInfo int, courseMetaNode *goquery.Selection) {
			name, _ := courseMetaNode.Attr("name")
			switch name {
			case "courseId":
				course.CourseID, _ = courseMetaNode.Attr("value")
				break
			case "classId":
				course.ClassID, _ = courseMetaNode.Attr("value")
				break
			}
		})
		var courseNameNode = courseNode.Find(".courseName").First()
		course.Title, _ = courseNameNode.Attr("title")
		href, _ := courseNameNode.Attr("href")
		course.URL = "http://mooc1-1.chaoxing.com" + href
		var teacherNode = courseNode.Find("p").First()
		course.Teacher, _ = teacherNode.Attr("title")
		courses = append(courses, course)
	})
	return courses, nil
}
