package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dariusbakunas/govirt-api/models"
)

type DomainsController struct{}

var domainModel = new(models.Domain)

func (d DomainsController) GetAllDomains(c *gin.Context) {
	domains, err := domainModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve domains", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"domains": domains})
	return
}