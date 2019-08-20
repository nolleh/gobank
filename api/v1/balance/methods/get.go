package methods

import (
	"github.com/labstack/echo"
	"gobank/common/types"
	"gobank/factory"
	"gobank/models"
	"net/http"
	"strconv"
)

func RouteGet(g *echo.Group) {
	g.GET("/:userId", Get)
}

// Get ...
func Get(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	ctx := c.Request().Context()
	traceId := factory.ApiContext(ctx).TraceId

	if err != nil {
		panic(err)
	}

	type Result struct {
		Balance models.BalanceEntity `json:"balance"`
	}

	balance := models.BalanceEntity{}
	if _, err := balance.GetById(ctx, userId); err != nil {
		panic(err)
	}

	result := Result{ balance }
	resp := types.ApiResponse{ Result: result, TraceId: traceId }

	factory.Logger(ctx).Info("get resp: ", resp)

	return c.JSON(http.StatusOK, &resp)
}

