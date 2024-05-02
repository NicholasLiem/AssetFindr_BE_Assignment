package user

import "github.com/NicholasLiem/GoLang_Microservice/router"

var Routes = router.RoutePrefix{
	Prefix: "/test",
	SubRoutes: []router.Route{
		{
			"Hello",
			"GET",
			"",
			HelloHandler,
			true,
		},
	},
}
