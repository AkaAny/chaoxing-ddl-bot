package superagent

import (
	"github.com/parnurzeal/gorequest"
	"net/url"
)

func mustParseURL(rawURL string) *url.URL {
	uri, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return uri
}

func CopySuperAgent(cur *gorequest.SuperAgent) *gorequest.SuperAgent {
	var request = gorequest.New()
	request.DoNotClearSuperAgent = cur.DoNotClearSuperAgent
	request.Transport.Proxy = cur.Transport.Proxy
	for key, values := range cur.Header {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
	//for _,cookieUrl:=range cookieUrls{
	//	uri:=mustParseURL(cookieUrl)
	//	request.Client.Jar.SetCookies(uri,cur.Client.Jar.Cookies(uri))
	//}
	copyJar(cur.Client.Jar, request.Client.Jar)
	return request
}
