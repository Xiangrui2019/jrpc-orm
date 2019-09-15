package routers

import (
	"jrpc-orm/api"
	"jrpc-orm/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	//router.Use(middlewares.SentryReportor())
	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)
	}

	return router
}
