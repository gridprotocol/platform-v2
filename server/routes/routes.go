package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gridprotocol/platform-v2/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// for swagger
	fmt.Print("register swagger")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// for test
	r.registRootRoute()

	// for functions
	r.registCpRoute()
	r.registNodeRoute()
	r.registOrderRoute()
	r.registGlobalRoute()

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
	// get a cp info by address
	r.GET("/v1/cp/:cp", GetCpInfoHandler())

	// get cp list by start and num
	r.GET("/v1/cp/list/:start/:num", ListCpHandler())

	// get node list of a cp
	r.GET("/v1/cp/:cp/node/list", CpNodeHandler())

	// get cp count
	r.GET("/v1/cp/count", CpCountHandler())
}

// node
func (r Routes) registNodeRoute() {
	// get all nodelist by node start and num
	r.GET("/v1/node/list/:start/:num", ListNodeHandler())

	// get nodelist of a user
	r.GET("/v1/user/node/list/:user", ListUserNodesHandler())

	//todo: node statics, currently get from contracts

	// id = cp:id
	r.GET("/v1/node/:cp/:id", GetNodeHandler())

	// set node status, in must be true or false
	r.POST("/v1/node/:cp/:id/:status/:in", SetNodeStatusHandler())
}

// order
func (r Routes) registOrderRoute() {
	// order by id
	r.GET("/v1/order/:id/info", GetOrderHandler())

	//r.GET("/v1/order/list/:user", GetOrdersHandler())
	// order fee by id
	r.GET("/v1/order/fee/:id", FeeOrderHandler())

	// orders of an user
	r.GET("/v1/user/:address/order/list", ListUserOrderHandler())

	// orders of a provider
	r.GET("/v1/provider/:address/order/list", LisProviderOrderHandler())

	// list providers of an user with active orders
	r.GET("/v1/user/:address/provider/list", ListOrderedProviderHandler())

	// get order count of a provider
	r.GET("/v1/provider/:address/count", GetOrderCountHandler())

	// set order status
	r.POST("/v1/order/:oid/:st", SetOrderStatusHandler())

	// check orders of a provider, if order is end, set status=4
	r.POST("/v1/check/order/:provider", CheckProviderOrderStatusHandler())

}

// global
func (r Routes) registGlobalRoute() {
	// get global info
	r.GET("/v1/global/", GetGlobalHandler())

	// increase cp num
	r.POST("v1/global/inccp/", IncCpHandler())
	r.POST("v1/global/incnode/:mem/:disk", IncNodeHandler())
	r.POST("v1/global/incused/:mem/:disk", IncUsedHandler())
	r.POST("v1/global/decused/:mem/:disk", DecUsedHandler())
}

// cors operation
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
