package cas

import (
	"ddl-bot/pkg/superagent"
	"ddl-bot/pkg/wrapper"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"regexp"
)

func Login(userName string, password string) (*gorequest.SuperAgent, error) {
	var request = gorequest.New()
	//request.DoNotClearSuperAgent = true
	var url = "https://cas.hdu.edu.cn/cas/login?service=https%3A%2F%2Fi.hdu.edu.cn%2Ftp_up%2F"
	request.Get(url)
	superagent.WithUA(request, USER_AGENT)
	resp, body, errs := request.Get(url).End()
	if errs != nil {
		return nil, errs[0]
	}
	err := ioutil.WriteFile(wrapper.GetPath("cas.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	ltValue, err := getLTValue(body)
	if err != nil {
		panic(err)
	}
	execValue, err := getExecutionValue(body)
	if err != nil {
		panic(err)
	}
	rsaValue, err := GetRSAValue(userName, password, ltValue)
	if err != nil {
		panic(err)
	}
	request.Post(url).Type("form")
	superagent.WithUA(request, USER_AGENT)
	request.Header.Set("Referer", url)
	request.Header.Set("Origin", "https://cas.hdu.edu.cn")
	request.Header.Set("sec-ch-ua", ";Not A Brand\";v=\"99\", \"Chromium\";v=\"88\"")
	request.Header.Set("sec-ch-ua-mobile", "?0")
	resp, body, errs = request.Send(map[string]interface{}{
		"rsa":       rsaValue,
		"ul":        len(userName),
		"pl":        len(password),
		"lt":        ltValue,
		"execution": execValue,
		"_eventId":  "submit",
	}).End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile(wrapper.GetPath("resp.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	if resp.Request.URL.Host != "i.hdu.edu.cn" {
		return nil, errors.New("fail to login by cas")
	}
	return request, nil
}

func getLTValue(body string) (string, error) {
	var lineExpr = regexp.MustCompile("<input type=\"hidden\" id=\"lt\" name=\"lt\" value=\".+\" />")
	var ltLine = lineExpr.FindString(body)
	if ltLine == "" {
		return "", errors.New("fail to find lt line")
	}
	fmt.Printf("lt line:%s\n", ltLine)
	return getValue(ltLine)
}

func getExecutionValue(body string) (string, error) {
	var lineExpr = regexp.MustCompile("<input type=\"hidden\" name=\"execution\" value=\".+\" />")
	var execLine = lineExpr.FindString(body)
	if execLine == "" {
		return "", errors.New("fail to find lt line")
	}
	fmt.Printf("execution line:%s\n", execLine)
	return getValue(execLine)
}
