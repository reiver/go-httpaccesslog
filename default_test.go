package httpaccesslog

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"testing"
)

func TestDefaultAccessLogWriter(t *testing.T) {

	tests := []struct {
		Method     string
		URL        string
		Proto      string
		RemoteAddr string

		ResponseStatusCode int
		ResponseSize       int

		ExpectedPrefix string
	}{
		{
			Method:     "DELETE",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "1.2.3.4:5678",

			ResponseStatusCode: 200,
			ResponseSize:       9012,

			ExpectedPrefix: `"request"."client-address"="1.2.3.4:5678"` + ` ` +
				`"request"."method"="DELETE"`                       + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="200"`                    + ` ` +
				`"response"."content-length"="9012"`,
		},
		{
			Method:     "GET",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "2.3.4.5:6789",

			ResponseStatusCode: 401,
			ResponseSize:       43210,

			ExpectedPrefix: `"request"."client-address"="2.3.4.5:6789"` + ` ` +
				`"request"."method"="GET"`                          + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="401"`                    + ` ` +
				`"response"."content-length"="43210"`,
		},
		{
			Method:     "PATCH",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "3.4.5.6:7890",

			ResponseStatusCode: 404,
			ResponseSize:       11223,

			ExpectedPrefix: `"request"."client-address"="3.4.5.6:7890"` + ` ` +
				`"request"."method"="PATCH"`                        + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="404"`                    + ` ` +
				`"response"."content-length"="11223"`,
		},
		{
			Method:     "POST",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "4.5.6.7:8901",

			ResponseStatusCode: 500,
			ResponseSize:       7,

			ExpectedPrefix: `"request"."client-address"="4.5.6.7:8901"` + ` ` +
				`"request"."method"="POST"`                         + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="500"`                    + ` ` +
				`"response"."content-length"="7"`,
		},
		{
			Method:     "PUT",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "5.6.7.8:9012",

			ResponseStatusCode: 302,
			ResponseSize:       13,

			ExpectedPrefix: `"request"."client-address"="5.6.7.8:9012"` + ` ` +
				`"request"."method"="PUT"`                          + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="302"`                    + ` ` +
				`"response"."content-length"="13"`,
		},

		{
			Method:     "KICK",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "6.7.8.9:0123",

			ResponseStatusCode: 208,
			ResponseSize:       918273645,

			ExpectedPrefix: `"request"."client-address"="6.7.8.9:0123"` + ` ` +
				`"request"."method"="KICK"`                         + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="208"`                    + ` ` +
				`"response"."content-length"="918273645"`,
		},
		{
			Method:     "PUNCH",
			URL:        "/apple/banana?cherry",
			Proto:      "HTTP/1.5",
			RemoteAddr: "7.8.9.0:1234",

			ResponseStatusCode: 222,
			ResponseSize:       2121212,

			ExpectedPrefix: `"request"."client-address"="7.8.9.0:1234"` + ` ` +
				`"request"."method"="PUNCH"`                        + ` ` +
				`"request"."uri"="/apple/banana?cherry"`            + ` ` +
				`"request"."protocol"="HTTP/1.5"`                   + ` ` +
				`"request"."host"=""`                               + ` ` +
				`"request"."content-length"="0"`                    + ` ` +
				`"request"."transfer-encoding"=[]`                  + ` ` +
				`"response"."status-code"="222"`                    + ` ` +
				`"response"."content-length"="2121212"`,
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var trace Trace
		trace.BeginTime = time.Now()
		trace.EndTime = trace.BeginTime.Add(5 * time.Minute)

		uri, err := url.Parse(test.URL)
		if nil != err {
			t.Errorf("For test #%d, should not have received an error when parsing URL string, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		r := http.Request{
			Method:     test.Method,
			URL:        uri,
			Proto:      test.Proto,
			RemoteAddr: test.RemoteAddr,
		}

		w := ResponseWriter{
			StatusCode: test.ResponseStatusCode,
			BodySize:   test.ResponseSize,
		}

		if err := DefaultAccessLogWriter(&buffer, &w, &r, &trace); nil != err {
			t.Errorf("For test #%d, did not expect to get an error, but actually got one: (%T) %v", testNumber, err, err)
		}

		expected := test.ExpectedPrefix +
			` "trace"."begin-time"=` + strconv.Quote(trace.BeginTime.String()) +
			` "trace"."end-time"=` + strconv.Quote(trace.EndTime.String()) +
			"\n"

		if actual := buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected:\n%q\nbut actually got\n%q.", testNumber, expected, actual)
			continue
		}
	}
}
