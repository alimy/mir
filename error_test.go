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

func TestHttpStatusCode(t *testing.T) {
	for _, data := range []struct {
		idx        int
		err        error
		statusCode int
		ok         bool
	}{
		{
			idx:        0,
			err:        errors.New("just a error"),
			statusCode: 0,
			ok:         false,
		},
		{
			idx:        1,
			err:        NewError(500, errors.New("internal error")),
			statusCode: 500,
			ok:         true,
		},
		{
			idx: 2,
			err: errors.Join(
				errors.New("error 1"),
				errors.New("error 2"),
				Errorln(404, "not found"),
			),
			statusCode: 404,
			ok:         true,
		},
		{
			idx: 3,
			err: errors.Join(
				errors.New("error 1"),
				errors.New("error 2"),
				Errorln(404, "not found"),
				Errorln(401, "auth error"),
				Errorln(501, "server error"),
			),
			statusCode: 404,
			ok:         true,
		},
		{
			idx: 4,
			err: errors.Join(
				errors.New("error 1"),
				errors.New("error 2"),
				errors.Join(
					errors.New("error 3"),
					errors.New("error 4"),
					Errorln(404, "not found"),
					Errorln(401, "auth error"),
					Errorln(501, "server error"),
				),
				errors.Join(
					errors.New("error 5"),
					errors.New("error 6"),
					Errorln(101, "error 7"),
					Errorln(201, "error 8"),
					Errorln(301, "error 9"),
				),
			),
			statusCode: 404,
			ok:         true,
		},
		{
			idx: 5,
			err: errors.Join(
				errors.New("error 1"),
				errors.New("error 2"),
				Errorf(20001, "custom error %d", 1),
				errors.Join(
					errors.New("error 3"),
					errors.New("error 4"),
					Errorln(404, "not found"),
					Errorln(401, "auth error"),
					Errorln(501, "server error"),
				),
				errors.Join(
					errors.New("error 5"),
					errors.New("error 6"),
					Errorln(101, "error 7"),
					Errorln(201, "error 8"),
					Errorln(301, "error 9"),
				),
			),
			statusCode: 20001,
			ok:         true,
		},
		{
			idx: 7,
			err: errors.Join(
				errors.New("error 1"),
				errors.New("error 2"),
				errors.Join(
					errors.New("error 3"),
					errors.New("error 4"),
				),
				errors.Join(
					errors.New("error 5"),
					errors.New("error 6"),
				),
			),
			statusCode: 0,
			ok:         false,
		},
	} {
		code, ok := HttpStatusCode(data.err)
		if code != data.statusCode || ok != data.ok {
			t.Errorf("expect HttpStatusCode[%d] result: (%d, %t) but got: (%d, %t)", data.idx, data.statusCode, data.ok, code, ok)
		}
	}
}
