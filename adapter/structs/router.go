package structs

import (
	"github.com/gin-gonic/gin"
)

var AppRoutes []RoutePrefix

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
	JSONRequest bool
}
