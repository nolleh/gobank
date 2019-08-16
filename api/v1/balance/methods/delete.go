package methods

import (
	"github.com/labstack/echo"
	"gobank/common/types"
	"gobank/echoMiddlewares"
	"gobank/factory"
	"gobank/models"
	"net/http"
	"strconv"
)

// RouteDelete ...
func RouteDelete(g *echo.Group) {
	g.DELETE("/:userId", Delete)
}

// Delete ... 
func Delete(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	ctx := c.Request().Context()
	traceId := ctx.Value(echoMiddlewares.ContextTraceId).(string)

	if err != nil {
		respError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: respError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	balance := models.BalanceEntity{ UserId: userId }
	
	if _, err := balance.Delete(ctx); err != nil {
		respError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: respError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	factory.Logger(ctx).Info("delete db", userId)

	type Result struct {
		message string
	}
	result := Result{ "OK" }
	resp := types.ApiResponse{ Result: result, TraceId: traceId }

	return c.JSON(http.StatusOK, &resp)
}