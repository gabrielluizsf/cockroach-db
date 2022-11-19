package server

import (
	"log"
	//"os"
	"github.com/GabrielLuizSF/cockroach-db/server/routes"
	"github.com/gin-gonic/gin"

)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	localPort := "5000";
	//serverPort := os.Getenv("PORT");
	return Server{
		port:  localPort,
		server: gin.Default(),
	}

}

func (s *Server) Run() {
        gin.SetMode(gin.ReleaseMode)
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
