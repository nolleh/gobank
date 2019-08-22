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
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	ctx := c.Request().Context()
	traceId := factory.ApiContext(ctx).TraceId

	if err != nil {
		panic(err)
	}

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		panic(err)
	}

	mapBalance := m["diffBalance"].(map[string]interface{}) // interface{}.(jsonobject)
	var diffBalance models.Balance
	if err := mapstructure.Decode(mapBalance, &diffBalance); err != nil {
		return err
	}

	mapAction := m["action"].(string)
	var strAction string
	if err := mapstructure.Decode(mapAction, &strAction); err != nil {
		return err
	}

	action := models.Deposit; if strAction == "withdraw" {
		action = models.Withdraw
	}

	balance, err := models.BalanceDatastore.UpdateByRelatively(ctx, userId, &diffBalance, action); if err != nil {
		return err
	}

	factory.Logger(ctx).Info(fmt.Sprint("modified db as parameter", diffBalance))

	type Result struct {
		Balance interface{}
	}
	result := Result { balance }
	resp := types.ApiResponse{ Result: result, TraceId:traceId }

	return c.JSON(http.StatusOK, &resp)
}

