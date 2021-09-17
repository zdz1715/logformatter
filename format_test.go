package logformatter

import (
	"net/http"
	"runtime/debug"
	"testing"
)

func TestLog(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		formatter := NewContext(Context{
			HttpRequest: HttpRequest{
				FullUrl:       "http://c.biancheng.net/view/124.html",
				PathInfo:      "view/124.html",
				ClientIp:      "127.0.0.1",
				RequestMethod: "POST",
				RequestHeader: http.Header{
					"Content-Type": []string{"application/json"},
				},
				RequestParams: "",
				ResponseHeader: http.Header{
					"Content-Type": []string{"application/json"},
				},
				ResponseBody:       "",
				ResponseStatusCode: 200,
			},
			Exec: Exec{
				ExecMs:             1.2,
				ExecException:      nil,
				ExecExceptionStack: string(debug.Stack()),
			},
		})
		//fmt.Printf("%+v", data)
		if err := formatter.Handle(); err != nil {
			t.Logf("err: %s", err.Error())
		}
	}
}
