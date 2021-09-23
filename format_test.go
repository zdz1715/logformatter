package logformatter

import (
	"errors"
	"net/http"
	"runtime/debug"
	"testing"
)

func Benchmark_Log(b *testing.B) {
	for i := 0; i <= b.N; i++ {
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
				ExecMs: 1.2,
			},
		})
		// 增加错误记录
		formatter.SetError(errors.New("it is a bug. "), string(debug.Stack()))

		if err := formatter.Handle(); err != nil {
			b.Logf("err: %s", err.Error())
		}
	}
}

func Benchmark_ExtraLog(b *testing.B) {
	for i := 0; i <= b.N; i++ {
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
				ExecException:      errors.New("it is a bug. ").Error(),
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
			b.Logf("err: %s", err.Error())
		}
	}
}
