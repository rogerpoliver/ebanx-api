package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rogerpoliver/ebanx-api/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func CreateServer() Server {
	return Server{
		port:   "9000",
		server: gin.Default(),
	}
}

func (s *Server) StartServer() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
