package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gridprotocol/platform-v2/server/routes"
)

type ServerOption struct {
	Endpoint       string
	Chain_Endpoint string
}

func NewServer(opt ServerOption) *http.Server {

	log.Println("Server Start")
	gin.SetMode(gin.ReleaseMode)

	// register routes
	router := routes.RegistRoutes()

	// start server
	srv := &http.Server{
		Addr:    opt.Endpoint,
		Handler: router,
	}

	return srv
}
