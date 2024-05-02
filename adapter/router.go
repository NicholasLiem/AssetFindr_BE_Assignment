package adapter

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/middleware"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/routes"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/structs"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/app"
	"github.com/gin-gonic/gin"
)

func NewRouter(server app.MicroserviceServer) *gin.Engine {
	router := gin.Default()

	structs.AppRoutes = append(structs.AppRoutes, routes.PostRoutes(server))
	for _, routePrefix := range structs.AppRoutes {
		group := router.Group(routePrefix.Prefix)

		for _, route := range routePrefix.SubRoutes {
			ginHandler := route.HandlerFunc

			if route.JSONRequest {
				ginHandler = middleware.ApplyJSONMiddleware(ginHandler)
			}

			switch route.Method {
			case "GET":
				group.GET(route.Pattern, ginHandler)
			case "POST":
				group.POST(route.Pattern, ginHandler)
			case "PATCH":
				group.PATCH(route.Pattern, ginHandler)
			case "DELETE":
				group.DELETE(route.Pattern, ginHandler)
			case "PUT":
				group.PUT(route.Pattern, ginHandler)
			}
		}
	}
	return router
}
