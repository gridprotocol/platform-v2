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
	r.registCpRoute()
	r.registNodeRoute()
	r.registOrderRoute()

	return r
}

// welcome
func (r Routes) registRootRoute() {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Server")
	})
}

// cp
func (r Routes) registCpRoute() {
	r.GET("/v1/cp/:cp", GetCpInfoHandler())
	r.GET("/v1/cp/list/:start/:num", ListCpHandler())

	// get node list of a cp
	r.GET("/v1/cp/:cp/node/list", ListNodeHandler())
}

// node
func (r Routes) registNodeRoute() {
	// id = cp:id
	r.GET("/v1/node/:id", GetNodeHandler())
}

// order
func (r Routes) registOrderRoute() {
	r.GET("/v1/order/:id/info", GetOrderHandler())
	//r.GET("/v1/order/list/:user", GetOrdersHandler())
	r.GET("/v1/order/fee/:id", FeeOrderHandler())

	r.GET("/v1/user/:address/order/list", ListActivedOrderHandler())
	// list providers of an user
	r.GET("/v1/user/:address/provider/list", ListOrderedProviderHandler())

	// get order count of a provider
	r.GET("/v1/provider/:address/count", GetOrderCountHandler())
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
