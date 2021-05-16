package chaoxing

import (
	"ddl-bot/pkg/cas"
	"ddl-bot/pkg/superagent"
	"ddl-bot/pkg/wrapper"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	url2 "net/url"
	"regexp"
	"strings"
	"time"
)

type Assignment struct {
	AssignmentID string `gorm:"column:assignment_id;primaryKey"`
	AnswerID     string `gorm:"column:answer_id;index"`
	Name         string `gorm:"column:name;index"`
	//URL       string    `gorm:"column:url"`
	StartTime time.Time `gorm:"column:start_time;index"`
	EndTime   time.Time `gorm:"column:end_time;index"`
	Status    string    `gorm:"column:status;index"`
}

func (cx *ChaoXing) GetAssignmentList(course *Course) ([]Assignment, error) {
	var request = cx.GetRequest()
	var url = course.URL
	request.Get(url)
	superagent.WithUA(request, cas.USER_AGENT)
	resp, body, errs := request.End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err := ioutil.WriteFile(wrapper.GetPath("resp_course_item.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var assignmentURL string = ""
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		title, _ := selection.Attr("title")
		if title != "作业" {
			return
		}
		data, _ := selection.Attr("data")
		assignmentURL = "http://mooc1-1.chaoxing.com" + data
	})
	if assignmentURL == "" {
		return nil, errors.New("failed to get assignment url")
	}
	request.Get(assignmentURL)
	superagent.WithUA(request, cas.USER_AGENT)
	resp, body, errs = request.End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile(wrapper.GetPath("resp_assignment.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var assignments []Assignment
	doc.Find(".titTxt").Each(func(iAssignment int, assignmentNode *goquery.Selection) {
		var assignment Assignment
		var metaNode = assignmentNode.Find("a").First()
		assignment.Name, _ = metaNode.Attr("title")
		fmt.Println(assignment.Name)
		href, _ := metaNode.Attr("href")
		var expr = regexp.MustCompile("workId=(.+)&")
		if !expr.MatchString(href) { //未完成作业没有直接url
			assignment.AssignmentID, _ = metaNode.Attr("data")
			assignment.AnswerID, _ = metaNode.Attr("data2")
		} else {
			url := "http://mooc1-1.chaoxing.com" + href
			uri, err := url2.Parse(url)
			if err != nil {
				fmt.Println("unexpected url for assignment:" + assignment.Name)
				return
			}
			var queryMap = uri.Query()
			assignment.AssignmentID = queryMap.Get("workId")
			assignment.AnswerID = queryMap.Get("workAnswerId")
		}

		assignmentNode.Find(".pt5").Each(func(iInfo int, infoNode *goquery.Selection) {
			var text = Trim(infoNode.Text())
			if strings.Contains(text, "：") {
				text = strings.Split(text, "：")[1]
			}
			switch iInfo {
			case 0: //开始时间
				startTime, err := ParseTime(text)
				if err != nil {
					fmt.Println("invalid start time")
				}
				assignment.StartTime = startTime
				break
			case 1: //截止时间
				deadline, err := ParseTime(text)
				if err != nil {
					fmt.Println("invalid end time")
				}
				assignment.EndTime = deadline
				break
			case 2: //作业状态
				assignment.Status = text
				break
			}
		})
		assignments = append(assignments, assignment)
	})
	return assignments, nil
}
