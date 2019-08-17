package methods

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"gobank/common/types"
	"gobank/factory"
	"gobank/models"
	"net/http"
	"strconv"
)

// RoutePost ...
func RoutePost(g *echo.Group) {
	g.POST("/:userId", Post)
}

// Post ...
// actionType: deposit / withdraw
func Post(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	ctx := c.Request().Context()
	traceId := factory.ApiContext(ctx).TraceId

	if err != nil {
		apiError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: apiError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		apiError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: apiError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	balance := models.BalanceEntity{ UserId: userId }
	mapBalance := m["diffBalance"].(map[string]interface{}) // interface{}.(jsonobject)
	var diffBalance models.Balance
	mapstructure.Decode(mapBalance, &diffBalance)

	if _, err := balance.UpdateByRelatively(ctx, diffBalance); err != nil {
		apiError := types.ApiError{ Code: -1, Message: err.Error() }
		resp := types.ApiResponse{ Error: apiError, TraceId: traceId }
		return c.JSON(http.StatusOK, &resp)
	}

	factory.Logger(ctx).Info(fmt.Sprint("modified db as parameter", diffBalance))

	type Result struct {
		balance models.BalanceEntity
	}
	result := Result { balance }
	resp := types.ApiResponse{ Result:result, TraceId:traceId }

	return c.JSON(http.StatusOK, &resp)
}

