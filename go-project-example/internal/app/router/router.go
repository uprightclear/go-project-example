package router

import (
	"github.com/gin-gonic/gin"
	"go-project-example/internal/app/handler"
	"go-project-example/internal/app/middleware"
)

func Router() error {
	router := gin.New()
	router.Use(gin.Recovery())
	apiGroup := router.Group("/api")
	{
		// 健康检查
		var healthCheck = apiGroup.Group("/healthcheck")
		{
			healthCheck.GET("")
			healthCheck.HEAD("")
			healthCheck.POST("/start")
			healthCheck.POST("/shutdown")
		}
		apiV1Group := apiGroup.Group("/v1", middleware.AuthMiddileware())
		{
			// 业务部分
			apiV1Group.POST("/test", handler.TestHandler)
		}
	}

	return router.Run()
}
