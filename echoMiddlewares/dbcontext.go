package echoMiddlewares

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

type ContextDBType string
var ContextDBName ContextDBType = "DB"

func DbContext(db *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// by testing result, new session doesn't make new connections.
			// for user convinient, managing internal memory for context or saving lastsql and so on...
			session := db.NewSession()
			defer session.Close()

			req := c.Request()
			ctx := req.Context()
			
			c.SetRequest(req.WithContext(context.WithValue(
					ctx,
					ContextDBName,
					session,
				),
			))

			switch req.Method {
			case "POST", "PUT", "DELETE":
				if err := session.Begin(); err != nil {
					return echo.NewHTTPError(500, err.Error())
				}
				if err := next(c); err != nil {
					session.Rollback()
					return echo.NewHTTPError(500, err.Error())
				}
				if c.Response().Status >= 500 {
					session.Rollback()
					return nil
				}
				if err := session.Commit(); err != nil {
					return echo.NewHTTPError(500, err.Error())
				}
			default:
				if err := next(c); err != nil {
					return echo.NewHTTPError(500, err.Error())
				}
			}

			return nil
		}
	}
}