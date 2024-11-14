package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/platform-v2/database"
	"github.com/gridprotocol/platform-v2/lib/utils"
)

// get order by id
func GetOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		id64, _ := utils.StringToUint64(id)

		order, err := database.GetOrderById(id64)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, order)
	}
}

// get order list by user
func GetOrdersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("user")

		orders, err := database.GetOrdersByUser(user)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, orders)
	}
}

// get provider list by user
func GetProsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("user")

		pros, err := database.GetProsByUser(user)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, pros)
	}
}

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
