package router

import (
	"partner.portal/internal/handlers"
	"partner.portal/internal/httpServer"
)

func RunAuthRouters() {
	g := httpServer.GetInstance().Group("/api/v1/auth")
	{
		g.POST("/register", handlers.RegHandler)
		g.POST("/login", handlers.LoginHandler)
	}
}
