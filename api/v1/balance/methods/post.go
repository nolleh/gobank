package methods

import (
	"net/http"
	"time"
	"strconv"
	"github.com/labstack/echo"
	// "github.com/sirupsen/logrus"

	"gobank/models"
	"gobank/factory"
)

// RoutePost ...
func RoutePost(g *echo.Group) {
	g.POST("/:userId", Post)
}

type ApiError struct {
	Code int
	Message string
}

// Get ...
func Get(c echo.Context) error {
	var content struct {
        Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
    content.Timestamp = time.Now().String()
	return c.JSON(http.StatusOK, &content)
}

// Post ...
func Post(c echo.Context) error {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {
		resp := ApiError{ Code: -1, Message: "Invalid Parameter" }
		return c.JSON(http.StatusOK, &resp)
	}

	var content struct {
		Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}

	balance := models.Balance{ UserId: userId, Amount:1000 }
	ctx := c.Request().Context()
	
	if _, err := balance.Create(ctx); err != nil {
		return err
	}
	
	factory.Logger(ctx).Info("inserted db as parameter")
	return c.JSON(http.StatusOK, &content)
}

