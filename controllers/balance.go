package controllers

import "github.com/labstack/echo"

type BalanceController struct {

}

func (c *BalanceController) Init(g *echo.group) {
	g.GET("/:userId", c.Get)
}

func (c *BalanceController) Get(c echo.Context) {
	jsonMap := map[string]string {
		"foo": "bar",
	}
	return c.JSON(http.StatusOK, jsonMap)
}
