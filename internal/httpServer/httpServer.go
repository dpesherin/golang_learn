package httpServer

import "github.com/gin-gonic/gin"

var RouterInstance *gin.Engine

func GetInstance() *gin.Engine {
	if RouterInstance == nil {
		RouterInstance = gin.Default()
	}
	return RouterInstance
}
