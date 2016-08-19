package httpaccesslog


import (
	"net/url"
)


func mustParseUrl(s string) *url.URL {

	u, err := url.Parse(s)
	if nil != err {
		panic(err)
	}

	return u
}
