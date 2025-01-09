package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/dumper/database"
	"github.com/gridprotocol/platform-v2/lib/utils"
)

// get order by id
// handler for get order by id
// GetOrderHandler godoc
//
//	@Summary		get order by id
//	@Description	get order by id
//	@Tags			GetOrderHandler
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"order id"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/order/{id}/info [get]
func GetOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		id64, _ := utils.StringToUint64(id)

		order, err := database.GetOrderById(id64)
		if err != nil {
			logger.Error(err.Error())
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
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
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
		}

		c.JSON(200, orders)
	}
}

// orders of an user
// list all orders of an user
// handler for list user orders
// ListUserOrderHandler godoc
//
//	@Summary		List an user orders
//	@Description	list an user orders
//	@Tags			ListUserOrders
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"user address"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/user/{address}/order/list [get]
func ListUserOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("address")
		orders, err := database.ListAllOrderByUser(user)
		if err != nil {
			logger.Error(err.Error())
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
		}

		c.JSON(200, orders)
	}
}

// list a provider's all orders
func LisProviderOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		provider := c.Param("address")
		orders, err := database.ListAllOrderByProvider(provider)
		if err != nil {
			logger.Error(err.Error())
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
		}

		c.JSON(200, orders)
	}
}

// active orders of an user
func ListUserActivedOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("address")
		orders, err := database.ListAllActivedOrderByUser(user)
		if err != nil {
			logger.Error(err.Error())
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
		}

		c.JSON(200, orders)
	}
}

// list cp of an user with active orders
// ListOrderedProviderHandler godoc
//
//	@Summary		list cp of an user with active orders
//	@Description	list cp of an user with active orders
//	@Tags			ListOrderedProviderHandler
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"user address"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/user/{address}/provider/list [get]
func ListOrderedProviderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user address
		address := c.Param("address")
		providers, err := database.ListAllOrderedProvider(address)
		if err != nil {
			logger.Error(err.Error())
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
		}

		c.JSON(200, providers)
	}
}

// get order count of a provider
// GetOrderCountHandler godoc
//
//	@Summary		get order count of a provider
//	@Description	get order count of a provider
//	@Tags			GetOrderCountHandler
//	@Accept			json
//	@Produce		json
//	@Param			address	path		string	true	"provider address"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/provider/{address}/count [get]
func GetOrderCountHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// provider address
		address := c.Param("address")
		cnt, err := database.GetOrderCount(address)
		if err != nil {
			logger.Error(err.Error())
			if err.Error() == "record not found" {
				c.AbortWithStatusJSON(200, gin.H{})
				return
			} else {
				c.AbortWithStatusJSON(400, err.Error())
				return
			}
		}

		c.JSON(200, cnt)
	}
}

// increase cp num by 1
func IncCpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := database.IncCp()
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func IncNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mem := c.Param("mem")
		disk := c.Param("disk")
		mem64, _ := utils.StringToInt64(mem)
		disk64, _ := utils.StringToInt64(disk)
		err := database.IncNode(mem64, disk64)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func IncUsedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mem := c.Param("mem")
		disk := c.Param("disk")
		mem64, _ := utils.StringToInt64(mem)
		disk64, _ := utils.StringToInt64(disk)
		err := database.IncUsed(mem64, disk64)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func DecUsedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mem := c.Param("mem")
		disk := c.Param("disk")
		mem64, _ := utils.StringToInt64(mem)
		disk64, _ := utils.StringToInt64(disk)
		err := database.DecUsed(mem64, disk64)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"result": "success",
		})
	}
}

// get the fee of an order
// FeeOrderHandler godoc
//
//	@Summary		get the fee of an order by id
//	@Description	get the fee of an order by id
//	@Tags			FeeOrderHandler
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"order id"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/order/fee/{id} [get]
func FeeOrderHandlerID() gin.HandlerFunc {
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

// order info for calc order fee
type OrderLite struct {
	Provider string `json:"provider"`
	// the node where the app is deployed
	Nid uint64 `json:"node_id"`
	// the service duration
	Duration int64 `json:"duration"`
}

// calc order fee
func FeeOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order OrderLite
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("order in body:", order)

		// 根据 order 中的 provider 和 nid 查找对应的 node
		node, err := database.GetNodeByCpAndId(order.Provider, order.Nid)
		if err != nil {
			c.AbortWithStatusJSON(400, fmt.Sprintf("get node error: %s", err.Error()))
			return
		}

		// 计算费用
		fee := calcOrderFee(node, order.Duration)

		// 返回费用
		c.JSON(http.StatusOK, gin.H{"fee": fee})
	}
}

// calcOrderFee 计算订单费用
func calcOrderFee(node database.Node, duration int64) float64 {
	// 将价格字符串转换为浮点数
	cpuPrice, _ := strconv.ParseFloat(node.CPUPrice.String(), 64)
	gpuPrice, _ := strconv.ParseFloat(node.GPUPrice.String(), 64)
	memPrice, _ := strconv.ParseFloat(node.MemPrice.String(), 64)
	diskPrice, _ := strconv.ParseFloat(node.DiskPrice.String(), 64)

	// 计算总费用
	totalFee := (cpuPrice + gpuPrice + memPrice + diskPrice) * float64(duration)
	return totalFee
}
