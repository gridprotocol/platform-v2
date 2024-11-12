package routes

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/platform-v2/database"
	"golang.org/x/xerrors"
)

func GetNodeInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeID := c.Param("id")

		address, id, err := decodeNodeID(nodeID)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(400, err.Error())
			return
		}

		node, err := database.GetNodeByAddressAndId(address, id)
		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(400, err.Error())
			return
		}

		c.JSON(200, node)
	}
}

func ListNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Query("address")

		var node []database.NodeStore
		var err error
		if address == "" {
			node, err = database.ListAllNodes()
		} else {
			node, err = database.ListAllNodesByProvider(address)
		}

		if err != nil {
			logger.Error(err.Error())
			c.AbortWithStatusJSON(500, err.Error())
			return
		}

		c.JSON(200, node)
	}
}

func decodeNodeID(nodeID string) (string, int, error) {
	results := strings.Split(nodeID, ":")
	if len(results) != 2 {
		return "", 0, xerrors.Errorf("node id(%s) Format Error, it should be {address}:{id}", nodeID)
	}

	id, err := strconv.Atoi(results[1])
	if err != nil {
		return "", 0, xerrors.Errorf("can't parse %s to int, %s", results[1], err.Error())
	}

	return results[0], id, nil
}
