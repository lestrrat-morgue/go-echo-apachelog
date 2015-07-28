package apachelog

// apachelog is a middleware to emit Apache HTTP Web Server style access
// logging.

import (
	"io"
	"time"

	"github.com/labstack/echo"
	logformat "github.com/lestrrat/go-apache-logformat"
)

// ApacheLog contains the basic information we need to log the access
type ApacheLog struct {
	LogFormat *logformat.ApacheLog
}

type ctxWithElapsedTime struct {
	*echo.Context
	start time.Time
}
func (c *ctxWithElapsedTime) ElapsedTime() time.Duration {
	return time.Since(c.start)
}
func (c *ctxWithElapsedTime) Response() logformat.Response {
	return c.Context.Response()
}

var newLine = []byte{'\n'}
func (l *ApacheLog) Wrap(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		start := time.Now()

		if err := h(c); err != nil {
			c.Error(err)
		}
		l.LogFormat.FormatCtx(
			l.LogFormat.Output(),
			&ctxWithElapsedTime{c, start},
		)
		l.LogFormat.Output().Write(newLine)
		return nil
	}
}

func Logger(dst io.Writer) echo.MiddlewareFunc {
	l := &ApacheLog{}
	l.LogFormat = logformat.CombinedLog.Clone()
	l.LogFormat.SetOutput(dst)
	return l.Wrap
}