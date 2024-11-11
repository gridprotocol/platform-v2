package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rockiecn/platform/database"
	"github.com/rockiecn/platform/logs"
)

var (
	// blockNumber = big.NewInt(0)
	logger = logs.Logger("routes")
)

func GetProviderInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")

		provider, err := database.GetProviderByAddress(address)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(400, err.Error())
			return
		}

		c.JSON(200, provider)
	}
}

func ListProviderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		providers, err := database.ListAllProviders()
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, providers)
	}
}
