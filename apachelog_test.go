package apachelog

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestLogger(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.NewContext(req, echo.NewResponse(rec), e)
	b := &bytes.Buffer{}

	// Status 2xx
	h := func(c *echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	Logger(b)(h)(c)

	// Status 3xx
	rec = httptest.NewRecorder()
	c = echo.NewContext(req, echo.NewResponse(rec), e)
	h = func(c *echo.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	}
	Logger(b)(h)(c)

	// Status 4xx
	rec = httptest.NewRecorder()
	c = echo.NewContext(req, echo.NewResponse(rec), e)
	h = func(c *echo.Context) error {
		return c.String(http.StatusNotFound, "test")
	}
	Logger(b)(h)(c)

	// Status 5xx with empty path
	req, _ = http.NewRequest(echo.GET, "", nil)
	rec = httptest.NewRecorder()
	c = echo.NewContext(req, echo.NewResponse(rec), e)
	h = func(c *echo.Context) error {
		return errors.New("error")
	}
	Logger(b)(h)(c)
}
