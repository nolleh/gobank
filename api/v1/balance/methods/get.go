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
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	ctx := c.Request().Context()
	traceId := factory.ApiContext(ctx).TraceId

	if err != nil {
		panic(err)
	}

	type Result struct {
		Balance interface{} `json:"balance"`
	}

	balance, err := models.BalanceDatastore.GetById(ctx, userId); if err != nil {
		panic(err)
	}

	result := Result{ balance }
	resp := types.ApiResponse{ Result: result, TraceId: traceId }

	factory.Logger(ctx).Info("get resp: ", resp)

	return c.JSON(http.StatusOK, &resp)
}

