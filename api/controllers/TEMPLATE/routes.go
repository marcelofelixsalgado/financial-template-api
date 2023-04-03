package TEMPLATE

import (
	"net/http"

	"github.com/marcelofelixsalgado/financial-commons/api/controllers"
)

var TEMPLATEBasepath = "/v1/TEMPLATES"

type TEMPLATERoutes struct {
	TEMPLATEHandler ITEMPLATEHandler
}

func NewTEMPLATERoutes(TEMPLATEHandler ITEMPLATEHandler) TEMPLATERoutes {
	return TEMPLATERoutes{
		TEMPLATEHandler: TEMPLATEHandler,
	}
}

func (TEMPLATERoutes *TEMPLATERoutes) TEMPLATERouteMapping() (string, []controllers.Route) {

	return TEMPLATEBasepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               TEMPLATERoutes.TEMPLATEHandler.CreateTEMPLATE,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               TEMPLATERoutes.TEMPLATEHandler.ListTEMPLATEs,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               TEMPLATERoutes.TEMPLATEHandler.GetTEMPLATEById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               TEMPLATERoutes.TEMPLATEHandler.UpdateTEMPLATE,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               TEMPLATERoutes.TEMPLATEHandler.DeleteTEMPLATE,
			RequiresAuthentication: true,
		},
	}
}
