package balance

import (
	"github.com/labstack/echo"
	"gobank/api/v1/balance/methods"
)

func Route(g *echo.Group) {
	methods.RoutePost(g)
}