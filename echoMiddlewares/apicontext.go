package echoMiddlewares

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/twinj/uuid"
)

/**
1. Add Trace Id for Tracking request
2. user Context
 */

var ContextTraceId string = "trace_id"

// middlewareFunc: return (next) => HandlerFunc
// next function as parameter, and returns function that injected (middleware) procedure
func ApiContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			genUuid := uuid.NewV4()
			c.SetRequest(req.WithContext(context.WithValue(ctx, ContextTraceId, genUuid.String())))
			fmt.Println("recv request, traceId:", genUuid)
			return next(c)
		}
	}
}
