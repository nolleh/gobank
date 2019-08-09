package controllers

import (
	"time"
	"github.com/labstack/echo"
)

type BalanceController struct {

}

func (b *BalanceController) Init(g *echo.Group) {
	g.GET("/:userId", b.Get)
}

func (b *BalanceController) Get(c *echo.Context) error {
	var content struct {
        Response  string `json:"response"`
        Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
    content.Timestamp = getNow()
	return c.JSON(http.StatusOK, &content)
}
