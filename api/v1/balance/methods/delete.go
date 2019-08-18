package methods

import (
	"github.com/labstack/echo"
	"gobank/common/types"
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
	traceId := factory.ApiContext(ctx).TraceId

	if err != nil {
		respError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: &respError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	balance := models.BalanceEntity{ UserId: userId }
	
	if _, err := balance.Delete(ctx); err != nil {
		respError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: &respError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	type Result struct {
		message string
	}
	result := Result{ "OK" }
	resp := types.ApiResponse{ Result: result, TraceId: traceId }

	factory.Logger(ctx).Info("delete db.. ", userId, ", resp:", resp)

	return c.JSON(http.StatusOK, &resp)
}