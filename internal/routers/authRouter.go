package router

import (
	"partner.portal/internal/controllers"
	"partner.portal/internal/httpServer"
)

func RunAuthRouters() {
	g := httpServer.GetInstance().Group("/api/v1/auth")
	{
		g.POST("/register", controllers.RegController)
		g.POST("/login", controllers.LoginController)
	}
}
