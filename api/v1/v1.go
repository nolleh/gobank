package v1

import (
	"gobank/api/v1/balance"
	"github.com/labstack/echo"
)

func Route(g *echo.Group) {
	router := g.Group("/balance")
	balance.Route(router)
}