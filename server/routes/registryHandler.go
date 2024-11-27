package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/dumper/database"
	"github.com/gridprotocol/platform-v2/lib/utils"
	"github.com/gridprotocol/platform-v2/logs"
	"golang.org/x/xerrors"
)

var (
	// blockNumber = big.NewInt(0)
	logger = logs.Logger("routes")
)

// get cp info
func GetCpInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cp := c.Param("cp")

		provider, err := database.GetProviderByAddress(cp)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(400, err.Error())
			return
		}

		c.JSON(200, provider)
	}
}

// list all cp
func ListCpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := c.Param("start")
		num := c.Param("num")

		iStart, _ := utils.StringToInt64(start)
		iNum, _ := utils.StringToInt64(num)

		providers, err := database.ListAllProviders(int(iStart), int(iNum))
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, providers)
	}
}

// get node, nodeID = cp:id
func GetNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeID := c.Param("id")

		// parse cp and id
		cp, id, err := decodeNodeID(nodeID)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(400, err.Error())
			return
		}

		node, err := database.GetNodeByCpAndId(cp, id)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(400, err.Error())
			return
		}

		c.JSON(200, node)
	}
}

// get node list of a cp
func ListNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cp := c.Query("cp")

		var nodes []database.NodeStore
		var err error
		if cp == "" {
			nodes, err = database.ListAllNodes()
		} else {
			nodes, err = database.ListAllNodesByCp(cp)
		}

		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, nodes)
	}
}

func decodeNodeID(nodeID string) (string, uint64, error) {
	results := strings.Split(nodeID, ":")

	if len(results) != 2 {
		return "", 0, xerrors.Errorf("node id(%s) Format Error, it should be {address}:{id}", nodeID)
	}

	// id, err := strconv.Atoi(results[1])
	// if err != nil {
	// 	return "", 0, xerrors.Errorf("can't parse %s to int, %s", results[1], err.Error())
	// }

	// string to uint64
	id, err := utils.StringToUint64(results[1])
	if err != nil {
		return "", 0, xerrors.Errorf("can't parse %s to uint64, %s", results[1], err.Error())
	}

	return results[0], id, nil
}
