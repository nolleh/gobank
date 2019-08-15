package methods

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
	// "github.com/sirupsen/logrus"
	"gobank/factory"
	"gobank/models"
)

// RouteDelete ...
func RouteDelete(g *echo.Group) {
	g.DELETE("/:userId", Delete)
}

// Delete ... 
func Delete(c echo.Context) error {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		resp := ApiError{ Code: -1, Message: "Invalid Parameter" }
		return c.JSON(http.StatusOK, &resp)
	}

	var content struct {
        Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
	content.Timestamp = time.Now().String()

	ctx := c.Request().Context()
	
	balance := models.Balance{ UserId: userId, Amount:1000 }
	
	if _, err := balance.Delete(ctx); err != nil {
		return err
	}
	
	factory.Logger(ctx).Info("delete db", userId)
	
	return c.JSON(http.StatusOK, &content)
}