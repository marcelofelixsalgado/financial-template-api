package routes

import (
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers/health"
	"github.com/marcelofelixsalgado/financial-commons/api/controllers"
	"github.com/marcelofelixsalgado/financial-commons/api/middlewares"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	TEMPLATERoutes TEMPLATE.TEMPLATERoutes
	healthRoutes   health.HealthRoutes
}

func NewRoutes(TEMPLATERoutes TEMPLATE.TEMPLATERoutes, healthRoutes health.HealthRoutes) *Routes {
	return &Routes{
		TEMPLATERoutes: TEMPLATERoutes,
		healthRoutes:   healthRoutes,
	}
}

func (routes *Routes) RouteMapping(http *echo.Echo) {

	// TEMPLATE routes
	basePath, TEMPLATERoutes := routes.TEMPLATERoutes.TEMPLATERouteMapping()
	setupRoute(http, basePath, TEMPLATERoutes)

	// health routes
	basePath, healthRoutes := routes.healthRoutes.HealthRouteMapping()
	setupRoute(http, basePath, healthRoutes)
}

func setupRoute(http *echo.Echo, basePath string, routes []controllers.Route) {
	group := http.Group(basePath)

	for _, route := range routes {
		if route.RequiresAuthentication {
			group.Add(route.Method, route.URI, route.Function, middlewares.Authenticate)
		} else {
			group.Add(route.Method, route.URI, route.Function)
		}
	}
}
