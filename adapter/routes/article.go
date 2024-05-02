package routes

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter/structs"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/app"
)

func ArticleRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/article",
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
				"/all",
				server.GetAllPost,
				true,
			},
			{
				"Show article with paging",
				"GET",
				"/{limit}/{offset}",
				server.GetPagedPost,
				true,
			},
			{
				"Show article with a specific id",
				"GET",
				"/{id}",
				server.GetPost,
				true,
			},
			{
				"Update article with new data",
				"PATCH",
				"/{id}",
				server.UpdatePost,
				true,
			},
			{
				"Delete article with specific id",
				"DELETE",
				"/{id}",
				server.DeletePost,
				true,
			},
		},
	}
}
