package mir

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	for _, data := range []struct {
		code int
		err  error
		msg  string
	}{
		{
			code: http.StatusInternalServerError,
			err:  errors.New(http.StatusText(http.StatusInternalServerError)),
			msg:  http.StatusText(http.StatusInternalServerError),
		},
		{
			code: http.StatusMethodNotAllowed,
			err:  errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			msg:  http.StatusText(http.StatusMethodNotAllowed),
		},
	} {
		err := NewError(data.code, data.err)
		code, msg := err.StatusCode(), err.Error()
		if code != data.code {
			t.Errorf("expect error code: %d but got: %d", data.code, code)
		}
		if msg != data.msg {
			t.Errorf("expect error msg: %s but got: %s", data.msg, msg)
		}
	}
}

func TestErrorf(t *testing.T) {
	for _, data := range []struct {
		code   int
		format string
		a      []any
	}{
		{
			code:   http.StatusInternalServerError,
			format: `internal server error: %s:%s`,
			a:      []any{"host", "localhost"},
		},
		{
			code:   http.StatusMethodNotAllowed,
			format: `method not allowed: %s:%s`,
			a:      []any{"method", "POST"},
		},
	} {
		err := Errorf(data.code, data.format, data.a...)
		fmtErrMsg := fmt.Errorf(data.format, data.a...).Error()
		code, msg := err.StatusCode(), err.Error()
		if code != data.code {
			t.Errorf("expect error code: %d but got: %d", data.code, code)
		}
		if msg != fmtErrMsg {
			t.Errorf("expect error msg: %s but got: %s", fmtErrMsg, msg)
		}
	}
}

func TestErrorln(t *testing.T) {
	for _, data := range []struct {
		code int
		text string
		msg  string
	}{
		{
			code: http.StatusInternalServerError,
			text: http.StatusText(http.StatusInternalServerError),
			msg:  http.StatusText(http.StatusInternalServerError),
		},
		{
			code: http.StatusMethodNotAllowed,
			text: http.StatusText(http.StatusMethodNotAllowed),
			msg:  http.StatusText(http.StatusMethodNotAllowed),
		},
	} {
		err := Errorln(data.code, data.text)
		code, msg := err.StatusCode(), err.Error()
		if code != data.code {
			t.Errorf("expect error code: %d but got: %d", data.code, code)
		}
		if msg != data.msg {
			t.Errorf("expect error msg: %s but got: %s", data.msg, msg)
		}
	}
}
