package superagent

import "github.com/parnurzeal/gorequest"

func WithUA(request *gorequest.SuperAgent, ua string) {
	request.Header.Set("User-Agent", ua)
}
