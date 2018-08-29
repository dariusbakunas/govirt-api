package server

import (
	"github.com/gin-gonic/gin"
	"github.com/dariusbakunas/govirt-api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	health := &controllers.HealthController{}
	domains := &controllers.DomainsController{}

	router.GET("/health", health.Status)

	users := router.Group("/domains")
	{
		users.GET("/", domains.GetAllDomains)
	}

	return router
}