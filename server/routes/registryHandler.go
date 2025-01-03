package routes

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/dumper/database"
	"github.com/gridprotocol/platform-v2/lib/utils"
	"github.com/gridprotocol/platform-v2/logs"
	"github.com/gridprotocol/platform-v2/sqldb"
	"golang.org/x/xerrors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// blockNumber = big.NewInt(0)
	logger = logs.Logger("routes")
)

// register cp
func RegCPHandler() gin.HandlerFunc {

	// handler func
	return func(c *gin.Context) {
		address := c.Request.Header.Get("address")
		name := c.Request.Header.Get("name")
		ip := c.Request.Header.Get("ip")
		port := c.Request.Header.Get("port")
		domain := c.Request.Header.Get("domain")

		// 插入记录的SQL语句
		insertSQL := `
	INSERT INTO providers (address, name, ip, port, domain, created_at) 
	VALUES (?, ?, ?, ?, ?, ?);`

		// current time
		createdAt := time.Now().Format("2006-01-02 15:04:05") // 格式化当前时间为YYYY-MM-DD HH:MM:SS

		// 执行插入操作
		_, err := sqldb.GRID_DB.Exec(insertSQL, address, name, ip, port, domain, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"msg": "[ACK] Record inserted successfully"})
	}
}

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
