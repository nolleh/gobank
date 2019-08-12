package api

import (
	"gobank/api/v1"
	"github.com/labstack/echo"
)

func Route(r *echo.Group) {
	router := r.Group("/v1")
	v1.Route(router)
}