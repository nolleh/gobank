package api

import (
	"github.com/labstack/echo"
	"gobank/api/v1"
)

func Route(r *echo.Group) {
	router := r.Group("/v1")
	v1.Route(router)
}