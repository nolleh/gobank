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

// RoutePost ...
func RoutePost(g *echo.Group) {
	g.POST("/:userId", Post)
}

// struct tag meaning
// `marshaling struct type: name, [omitempty]`
type ApiError struct {
	Code int `json:"code"`
	Message string `json:"message"`
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

