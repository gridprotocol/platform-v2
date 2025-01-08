package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/dumper/database"
	"github.com/gridprotocol/platform-v2/lib/utils"
	"github.com/gridprotocol/platform-v2/logs"
)

var (
	// blockNumber = big.NewInt(0)
	logger = logs.Logger("routes")
)

// get cp info
// get a cp info
// handler for get a cp info
// GetCpInfoHandler godoc
//
//	@Summary		get a cp info
//	@Description	get a cp info
//	@Tags			GetCpInfoHandler
//	@Accept			json
//	@Produce		json
//	@Param			cp	path		string	true	"cp address"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/cp/{cp} [get]
func GetCpInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cp := c.Param("cp")

		provider, err := database.GetProviderByAddress(cp)
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

		c.JSON(200, provider)
	}
}

// list all cp
// handler for list cp
// ListCpHandler godoc
//
//	@Summary		list cp info
//	@Description	list cp info
//	@Tags			ListCpHandler
//	@Accept			json
//	@Produce		json
//	@Param			start	path		string	true	"start"
//	@Param			num	path		string	true	"num"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/cp/list/{start}/{num} [get]
func ListCpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := c.Param("start")
		num := c.Param("num")

		iStart, _ := utils.StringToInt64(start)
		iNum, _ := utils.StringToInt64(num)

		providers, err := database.ListAllProviders(int(iStart), int(iNum))
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

// list all nodes by specify node start and num
// handler for list nodes
// ListNodeHandler godoc
//
//	@Summary		List all nodes
//	@Description	list all nodes
//	@Tags			Listnodes
//	@Accept			json
//	@Produce		json
//	@Param			start	path		string	true	"start"
//	@Param			num		path		string	true	"number"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/node/list/{start}/{num} [get]
func ListNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := c.Param("start")
		num := c.Param("num")

		iStart, _ := utils.StringToInt64(start)
		iNum, _ := utils.StringToInt64(num)

		nodes, err := database.ListAllNodes(int(iStart), int(iNum))
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

		c.JSON(200, nodes)
	}
}

// get node
// handler for get a node
// GetNodeHandler godoc
//
//	@Summary		get a node
//	@Description	get a node
//	@Tags			GetNodeHandler
//	@Accept			json
//	@Produce		json
//	@Param			cp	path		string	true	"cp address"
//	@Param			id		path		string	true	"node id"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/node/{cp}/{id} [get]
func GetNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpAddr := c.Param("cp")
		nodeID := c.Param("id")

		id, _ := utils.StringToUint64(nodeID)

		// // parse cp and id
		// cp, id, err := decodeNodeID(nodeID)
		// if err != nil {
		// 	logger.Error(err.Error())
		// 	c.AbortWithStatusJSON(400, err.Error())
		// 	return
		// }

		node, err := database.GetNodeByCpAndId(cpAddr, id)
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

		c.JSON(200, node)
	}
}

// set node's status
func SetNodeStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpAddr := c.Param("cp")
		nodeID := c.Param("id")
		status := c.Param("status")
		in := c.Param("in")
		var bVal bool

		switch in {
		case "true":
			bVal = true
		case "false":
			bVal = false
		default:
			c.AbortWithStatusJSON(400, fmt.Errorf("in must be true or false"))
			return
		}

		id, _ := utils.StringToUint64(nodeID)

		switch status {
		case "exist":
			database.SetExist(cpAddr, id, bVal)
		case "sold":
			database.SetSold(cpAddr, id, bVal)
		case "avail":
			database.SetAvail(cpAddr, id, bVal)
		case "online":
			database.SetOnline(cpAddr, id, bVal)
		default:
			c.AbortWithStatusJSON(400, fmt.Errorf("status must be exist, sold, avail, online"))
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "Success",
		})
	}
}

// set an order's status
func SetOrderStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		oid := c.Param("oid")
		st := c.Param("st")

		oid64, _ := utils.StringToUint64(oid)
		st64, _ := utils.StringToUint64(st)

		database.SetOrderStatus(oid64, st64)

		c.JSON(200, gin.H{
			"code":    200,
			"message": "Success",
		})
	}
}

// check a provider's all orders, set status=4 if an order is end
func CheckProviderOrderStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		provider := c.Param("provider")

		database.CheckProviderOrders(provider)

		c.JSON(200, gin.H{
			"code":    200,
			"message": "Success",
		})
	}
}

// get node list of a cp
// handler for get node list of a cp
// CpNodeHandler godoc
//
//	@Summary		list cp nodes
//	@Description	list cp nodes
//	@Tags			CpNodeHandler
//	@Accept			json
//	@Produce		json
//	@Param			cp	path		string	true	"cp address"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/cp/{cp}/node/list [get]
func CpNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cp := c.Param("cp")

		var nodes []database.NodeStore
		var err error

		nodes, err = database.ListAllNodesByCp(cp)
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

		c.JSON(200, nodes)
	}
}

// list all nodes of an user
// handler for list user nodes
// ListUserNodesHandler godoc
//
//	@Summary		List an user nodes
//	@Description	list an usernodes
//	@Tags			ListUserNodes
//	@Accept			json
//	@Produce		json
//	@Param			user	path		string	true	"user"
//	@Success		200		{object}	int
//	@Failure		404		{object}	string	"page not found"
//	@Router			/v1/user/node/list/{user} [get]
//
// get node list by an user
func ListUserNodesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Param("user")

		var nodes []database.NodeAdaptor
		var err error

		nodes, err = database.ListAllNodesByUser(user)
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

		c.JSON(200, nodes)
	}
}

// func decodeNodeID(nodeID string) (string, uint64, error) {
// 	results := strings.Split(nodeID, ":")

// 	if len(results) != 2 {
// 		return "", 0, xerrors.Errorf("node id(%s) Format Error, it should be {address}:{id}", nodeID)
// 	}

// 	// id, err := strconv.Atoi(results[1])
// 	// if err != nil {
// 	// 	return "", 0, xerrors.Errorf("can't parse %s to int, %s", results[1], err.Error())
// 	// }

// 	// string to uint64
// 	id, err := utils.StringToUint64(results[1])
// 	if err != nil {
// 		return "", 0, xerrors.Errorf("can't parse %s to uint64, %s", results[1], err.Error())
// 	}

// 	return results[0], id, nil
// }
