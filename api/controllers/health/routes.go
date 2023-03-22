package health

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers"
)

var basepath = "/v1/health"

type HealthRoutes struct {
	healthHandler IHealthHandler
}

func NewHealthRoutes(healthHandler IHealthHandler) HealthRoutes {
	return HealthRoutes{
		healthHandler: healthHandler,
	}
}

func (healthRoutes *HealthRoutes) HealthRouteMapping() (string, []controllers.Route) {
	return basepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               healthRoutes.healthHandler.Health,
			RequiresAuthentication: false,
		},
	}
}
