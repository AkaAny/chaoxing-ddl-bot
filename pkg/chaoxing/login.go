package chaoxing

import (
	"ddl-bot/pkg/cas"
	"ddl-bot/pkg/form"
	"ddl-bot/pkg/superagent"
	"ddl-bot/pkg/wrapper"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

type ChaoXing struct {
	request *gorequest.SuperAgent
}

func (cx *ChaoXing) GetRequest() *gorequest.SuperAgent {
	var request = superagent.CopySuperAgent(cx.request)
	return request
}

func Login(casRequest *gorequest.SuperAgent) (*ChaoXing, error) {
	var request = superagent.CopySuperAgent(casRequest)
	request.RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
		fmt.Println(req.URL)
		return nil
	})
	fmt.Println(request)
	var url = "https://hdu.fanya.chaoxing.com/sso/hdu"
	request.Get(url)
	superagent.WithUA(request, cas.USER_AGENT)
	resp, body, errs := request.End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err := ioutil.WriteFile(wrapper.GetPath("resp.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	userLoginForm, err := form.Parse(body)
	if err != nil {
		return nil, err
	}
	resp, body, err = userLoginForm["userLogin"].Submit(request, func() error {
		superagent.WithUA(request, cas.USER_AGENT)
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(wrapper.GetPath("resp_2.html"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	if resp.Request.URL.Host != "hdu.fanya.chaoxing.com" {
		return nil, errors.New("failed to login to chaoxing")
	}
	request.Get("http://i.mooc.chaoxing.com/space/index.shtml")
	superagent.WithUA(request, cas.USER_AGENT)
	resp, body, errs = request.End()
	if errs != nil {
		return nil, errs[0]
	}
	fmt.Println(resp.StatusCode)
	err = ioutil.WriteFile(wrapper.GetPath("resp_index.shtml"), []byte(body), 0644)
	if err != nil {
		panic(err)
	}
	return &ChaoXing{
		request: request,
	}, nil
}
