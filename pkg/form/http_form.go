package form

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"strings"
)

type Form struct {
	ActionURL string
	Method    string
	InputMap  map[string]string
}

func Parse(body string) (map[string]Form, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var resultMap = make(map[string]Form)
	doc.Find("form").Each(func(iForm int, formNode *goquery.Selection) {
		var form = Form{InputMap: make(map[string]string)}
		formID, _ := formNode.Attr("id")
		form.ActionURL, _ = formNode.Attr("action")
		form.Method, _ = formNode.Attr("method")
		form.Method = strings.ToUpper(form.Method) //统一转大写
		//解析input标签
		formNode.ChildrenFiltered("input").Each(func(iInput int, inputNode *goquery.Selection) {
			name, _ := inputNode.Attr("name")
			typ, _ := inputNode.Attr("type")
			if strings.ToLower(typ) == "submit" {
				return
			}
			val, _ := inputNode.Attr("value")
			form.InputMap[name] = val
		})
		resultMap[formID] = form
	})
	return resultMap, nil
}

func (f Form) Submit(request *gorequest.SuperAgent, onSubmit func() error) (gorequest.Response, string, error) {
	request.CustomMethod(f.Method, f.ActionURL).Type("form")
	if onSubmit != nil {
		err := onSubmit()
		if err != nil {
			return nil, "", err
		}
	}
	resp, body, errs := request.Send(f.InputMap).End()
	if errs != nil {
		return nil, "", errs[0]
	}
	fmt.Println(resp.StatusCode)
	return resp, body, nil
}
