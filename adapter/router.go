package adapter

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/middleware"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/routes"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/structs"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/app"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(server app.MicroserviceServer) *mux.Router {
	router := mux.NewRouter()

	structs.AppRoutes = append(structs.AppRoutes, routes.ArticleRoutes(server))
	for _, route := range structs.AppRoutes {

		//create sub route
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//for each sub route
		for _, subRoute := range route.SubRoutes {

			var handler http.Handler
			handler = subRoute.HandlerFunc

			if subRoute.JSONRequest {
				handler = middleware.JSONMiddleware(subRoute.HandlerFunc) // use middleware
			}

			//register the route
			routePrefix.Path(subRoute.Pattern).Handler(handler).Methods(subRoute.Method).Name(subRoute.Name)
		}

	}

	return router
}
