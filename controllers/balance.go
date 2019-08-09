package controllers

import (
	"time"
	"net/http"
	"github.com/labstack/echo"
	"github.com/nolleh/gobank/models"
)

type BalanceController struct {

}

func (b BalanceController) Init(g *echo.Group) {
	g.GET("/:userId", b.Get)
	g.POST("/:userId", b.Post)
}

func (BalanceController) Get(c echo.Context) error {
	var content struct {
        Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
    content.Timestamp = time.Now().String()
	return c.JSON(http.StatusOK, &content)
}

func (BalanceController) Post(c echo.Context) error {
	var content struct {
		Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}

	balance := models.Balance{}

	if _, err := balance.Create(c.Request().Context()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &content)
}