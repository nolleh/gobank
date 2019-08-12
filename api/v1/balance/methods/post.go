package methods

import (
	"time"
	"net/http"
	"github.com/labstack/echo"
	// "github.com/sirupsen/logrus"

	"gobank/models"
	"gobank/factory"
)

type BalanceController struct {

}

func RoutePost(g *echo.Group) {
	g.POST("/:userId", Post)
}

func Get(c echo.Context) error {
	var content struct {
        Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
    content.Timestamp = time.Now().String()
	return c.JSON(http.StatusOK, &content)
}

func Post(c echo.Context) error {
	var content struct {
		Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}

	balance := models.Balance{ Amount:1 }
	// balance2 := models.Balance{ Amount: 2}

	if _, err := balance.Create(c.Request().Context()); err != nil {
		return err
	}

	// if _, err := balance2.Create(c.Request().Context()); err != nil {
	// 	return err
	// }
	
	ctx := c.Request().Context()
	factory.Logger(ctx).Info("inserted db as parameter")
	factory.Logger(ctx).Info("inserted db as parameter2")
	return c.JSON(http.StatusOK, &content)
}