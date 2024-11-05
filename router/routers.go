package router

import (
	"github.com/Takina-Space/backend-go/app/middleware"
	route "github.com/Takina-Space/backend-go/router/v1"
	"github.com/gin-gonic/gin"
)

// InitRoutesJWT function route that use JWT midlleware
func InitRoutesJWT(g *gin.RouterGroup) {
	// Initialize Midlleware
	g.Use(middleware.ErrorHandler())
	g.Use(middleware.JSONMiddleware())
	//g.Use(middleware.AuthorizeJWT())
	// Initialize route
	route.ExampleRoute(g)

}

// InitRoutes function route for home or some url not using a JWT Auth
func InitRoutes(g *gin.RouterGroup) {
	g.Use(middleware.ErrorHandler())
	//g.Use(middleware.JSONMiddleware())
	// Initialize route
	route.SetHomeRoutes(g)
}
