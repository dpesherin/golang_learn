package main

import (
	"partner.portal/internal/config"
	"partner.portal/internal/database"
	"partner.portal/internal/httpServer"
	"partner.portal/internal/routers"
)

func main() {
	cfg := config.MustLoad()
	database.GetInstance()
	engine := httpServer.GetInstance()
	router.InitializeRouting()
	err := engine.Run(cfg.HttpServer.Address)
	if err != nil {
		panic(err)
	}
}
