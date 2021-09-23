package logformatter

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	Msg        = "auto-log"
	ContextKey = "context"
)

type FormatterParams struct {
	Logger *zerolog.Logger
	Level  zerolog.Level
	Context
}

type Context struct {
	DbSql       []DbSql `json:"db_sql,omitempty"`
	HttpRequest `json:",inline"`
	Exec        `json:",inline"`
	Extra       map[string]interface{} `json:"extra"`
}

type Exec struct {
	ExecMs             float64 `json:"exec_ms"`
	ExecException      string  `json:"exec_exception"`
	ExecExceptionStack string  `json:"exec_exception_stack"`
}

type DbSql struct {
	ConnectionName string `json:"connection_name"`
	Sql            string `json:"sql"`
	Bindings       string `json:"bindings"`
	Ms             string `json:"ms"`
}

type HttpRequest struct {
	FullUrl            string      `json:"full_url"`
	PathInfo           string      `json:"path_info"`
	ClientIp           string      `json:"client_ip"`
	RequestMethod      string      `json:"request_method"`
	RequestHeader      http.Header `json:"request_header"`
	RequestParams      string      `json:"request_params"`
	ResponseHeader     http.Header `json:"response_header"`
	ResponseBody       string      `json:"response_body"`
	ResponseStatusCode int         `json:"response_status_code"`
}

func (u *FormatterParams) SetError(err error, stacks ...string) *FormatterParams {
	if err != nil {
		u.ExecException = err.Error()
		if len(stacks) > 0 {
			u.ExecExceptionStack = stacks[0]
		}
	}
	return u
}

func (u *FormatterParams) resetExtra() {
	u.Context.Extra = make(map[string]interface{})
}

func NewContext(context Context) *FormatterParams {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	u := &FormatterParams{
		Context: context,
		Level:   zerolog.InfoLevel,
	}
	return u
}

func (u *FormatterParams) SetLevel(level zerolog.Level) *FormatterParams {
	u.Level = level
	return u
}

func (u *FormatterParams) SetExtra(key string, context interface{}) *FormatterParams {
	if u.Context.Extra == nil {
		u.resetExtra()
	}
	u.Context.Extra[key] = context
	return u
}

func (u *FormatterParams) Handle(message ...string) error {
	jsonA, err := json.Marshal(u.Context)
	if err != nil {
		return err
	}
	msg := Msg
	if len(message) > 0 {
		msg = message[0]
	}
	log.WithLevel(u.Level).RawJSON(ContextKey, jsonA).Msg(msg)
	return nil
}
