package routes

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/structs"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/app"
)

func PostRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/api/posts",
		SubRoutes: []structs.Route{
			{
				"Create a new article",
				"POST",
				"",
				server.CreatePost,
				true,
			},
			{
				"Get all posts data",
				"GET",
				"",
				server.GetAllPost,
				false,
			},
			{
				"Show article with a specific id",
				"GET",
				"/:id",
				server.GetPost,
				false,
			},
			{
				"Update article with new data",
				"PUT",
				"/:id",
				server.UpdatePost,
				true,
			},
			{
				"Delete article with specific id",
				"DELETE",
				"/:id",
				server.DeletePost,
				false,
			},
		},
	}
}
