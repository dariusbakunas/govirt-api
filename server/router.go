package server

import (
	"github.com/gin-gonic/gin"
	"github.com/dariusbakunas/govirt-api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	router.Use(gin.Recovery())

	health := &controllers.HealthController{}
	domains := &controllers.DomainsController{}

	router.GET("/health", health.Status)

	users := router.Group("/domains")
	{
		users.GET("/", domains.GetAllDomains)
	}

	return router
}