package controllers

import "github.com/labstack/echo"

type BalanceController struct {

}

func (b *BalanceController) Init(g *echo.Group) {
	g.GET("/:userId", b.Get)
}

func (b *BalanceController) Get(c *echo.Context) error {
	jsonMap := map[string]string {
		"foo": "bar",
	}
	return c.JSON(http.StatusOK, jsonMap)
}
