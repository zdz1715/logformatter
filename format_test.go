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
		//fmt.Printf("%+v", formatter)
		if err := formatter.Handle(); err != nil {
			t.Logf("err: %s", err.Error())
		}
	}
}

func TestExtraLog(t *testing.T) {
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

		// 增加自定义日志结构
		formatter.SetExtra("db_sql", []string{
			"select * from table1",
			"select * from table2",
		})

		//fmt.Printf("%+v", formatter)
		if err := formatter.Handle(); err != nil {
			t.Logf("err: %s", err.Error())
		}
	}
}
