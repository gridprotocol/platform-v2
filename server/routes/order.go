package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/platform-v2/database"
)

func ListActivedOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")
		orders, err := database.ListAllActivedOrderByUser(address)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, orders)
	}
}

func ListOrderedProviderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")
		providers, err := database.ListAllOrderedProvider(address)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, providers)
	}
}
