package balance

import (
	"github.com/labstack/echo"
	"gobank/api/v1/balance/methods"
)

func Route(g *echo.Group) {
	methods.RouteGet(g)
	methods.RoutePost(g)
	methods.RouteDelete(g)
}