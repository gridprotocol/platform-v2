package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/dumper/database"
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

func ListActivedOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("address")
		orders, err := database.ListAllActivedOrderByUser(user)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, orders)
	}
}

// list cp of an user
func ListOrderedProviderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user address
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

// get the fee of an order
func FeeOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// order id
		id := c.Param("id")
		id64, _ := utils.StringToUint64(id)
		fee, err := database.CalcOrderFee(id64)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, fee)
	}
}
