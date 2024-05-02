package jwt

import (
	"github.com/NicholasLiem/GoLang_Microservice/router"
)

var Routes = router.RoutePrefix{
	Prefix: "/jwt",
	SubRoutes: []router.Route{
		{
			"GetJWT",
			"GET",
			"",
			GetJwt,
			false,
		},
	},
}
