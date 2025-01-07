package routers

import (
	"ADMgmtSystem/controllers"

	"github.com/gin-gonic/gin"
)

// Route struct
type Route struct {
	Method     string
	Pattern    string
	Handler    gin.HandlerFunc
	Middleware gin.HandlerFunc
}

var routes []Route

func init() {
	register("GET", "/api/SayHello", controllers.SayHello, nil)
}

// NewRouter
// @return *gin.Engine
func NewRouter() *gin.Engine {
	router := gin.Default() // Use gin default engine
	for _, route := range routes {
		// Has middleware use this router
		if route.Middleware != nil {
			router.Handle(route.Method, route.Pattern, route.Middleware, route.Handler)
		} else {
			router.Handle(route.Method, route.Pattern, route.Handler)
		}
	}
	return router
}

// register
func register(method, pattern string, handler gin.HandlerFunc, middleware gin.HandlerFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
