package v1

import (
	"github.com/labstack/echo"
	"gobank/api/v1/balance"
)

func Route(g *echo.Group) {
	router := g.Group("/balance")
	balance.Route(router)
}