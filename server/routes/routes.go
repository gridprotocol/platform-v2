package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	*gin.Engine
}

type NodeInfo struct {
	Name     string `json:"name"`
	Entrance string `json:"entrance"`
	Resource string `json:"resource"`
	Price    string `json:"price"`
}

type OrderInfo struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
	Duration string `json:"duration"`
	Price    string `json:"price"`
}

func init() {

}

// register all routes for server
func RegistRoutes() Routes {

	router := gin.Default()

	router.Use(cors())

	r := Routes{
		router,
	}

	// for test
	r.registRootRoute()

	// for functions
	r.registListRoute()
	r.registOrderRoute()

	return r
}

// welcome
func (r Routes) registRootRoute() {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Server")
	})
}

// gateway list, list all nodes' gateway info
func (r Routes) registListRoute() {

	// read nodes from config
	nodes := ReadList()

	// response node list
	r.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, nodes)
	})
}

// order operation in platform
func (r Routes) registOrderRoute() {
	or := OrderInfo{ID: "123", Resource: "res", Duration: "dur", Price: "100"}
	// response order
	r.GET("/order", func(c *gin.Context) {
		c.JSON(http.StatusOK, or)
	})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
