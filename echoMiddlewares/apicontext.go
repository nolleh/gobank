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

var ApiContextName string = "api_context"

type ApiContext struct {
	TraceId string /* for log tracking id. it reads from header or if not exist, generate one. */
}

// middlewareFunc: return (next) => HandlerFunc
// next function as parameter, and returns function that injected (middleware) procedure
func InjectApiContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			genUuid := uuid.NewV4()

			apiCtx := ApiContext{ TraceId: genUuid.String() }

			c.SetRequest(req.WithContext(context.WithValue(ctx, ApiContextName, &apiCtx)))
			fmt.Println("recv request, traceId:", genUuid)
			return next(c)
		}
	}
}
