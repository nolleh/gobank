package echoMiddlewares

import (
	"fmt"
	"github.com/labstack/echo"
	"gobank/common/types"
	"gobank/utils"
	"net/http"
)

func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			v := ctx.Value(ApiContextName).(*ApiContext)

			// handling error, before http/net has come to catch
			// this is why running next function before return this function.
			// defer functions are run at the end of a function call, right before that function return

			defer func() {
				if err := recover(); err != nil {
					detail := utils.Stringify(err)
					fmt.Println("caught error!", err)
					apiError := types.ApiError{Code: -1, Message: detail}
					resp := types.ApiResponse{Error: &apiError, TraceId: v.TraceId}
					c.JSON(http.StatusOK, &resp)
				}
			}()

			res := next(c)
			return res
		}
	}
}
