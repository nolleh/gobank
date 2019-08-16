package logger

import (
	"context"
	"fmt"
	"gobank/echoMiddlewares"
	"log"
	"runtime"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const ContextLoggerName = "ContextLogger"

func ContextLogger() echo.MiddlewareFunc {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(logger.Writer())

	// logger.Formatter = &logrus.JSONFormatter{}
	logger.Level = logrus.InfoLevel

	hooks := logrus.LevelHooks{}
	hooks.Add(&CallkerHook{})
	logger.Hooks = hooks

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logEntry := logrus.NewEntry(logger)

			id := c.Request().Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = c.Request().Context().Value(echoMiddlewares.ContextTraceId).(string)
			}
			logEntry = logEntry.WithField(echoMiddlewares.ContextTraceId, id)

			req := c.Request()
			c.SetRequest(req.WithContext(context.WithValue(req.Context(), ContextLoggerName, logEntry)))

			return next(c)
		}
	}
}

type CallkerHook struct{}

func (c *CallkerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}
func (c *CallkerHook) Fire(entry *logrus.Entry) error {
	var ok bool
	_, file, line, ok := runtime.Caller(5)
	if !ok {
		file = "???"
		line = 0
	}
	entry.Data["caller"] = fmt.Sprintf("%s:%d", file, line)
	return nil
}
