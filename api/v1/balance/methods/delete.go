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
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	ctx := c.Request().Context()
	traceId := factory.ApiContext(ctx).TraceId

	if err != nil {
		panic(err)
	}

	if err := models.BalanceDatastore.Delete(ctx, userId); err != nil {
		panic(err)
	}

	type Result struct {
		Message string `json:"message"`
	}
	result := Result{ "OK" }
	resp := types.ApiResponse{ Result: result, TraceId: traceId }

	factory.Logger(ctx).Info("delete db.. ", userId, ", resp:", resp)

	return c.JSON(http.StatusOK, &resp)
}