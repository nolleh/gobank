package controllers

import "github.com/labstack/echo"

type BalanceController struct {

}

func (c *BalanceController) Init(g *echo.Group) {
	g.GET("/:userId", c.Get)
}

func (b *BalanceController) Get(c echo.Context) {
	jsonMap := map[string]string {
		"foo": "bar",
	}
	return c.JSON(http.StatusOK, jsonMap)
}
