package api

import (
	"github.com/Pedroxer/avito/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	db     *sqlc.Queries
}

func NewServer(db *sqlc.Queries) *Server {
	server := &Server{db: db}

	server.setUpRoutes()
	return server
}

func (serv *Server) setUpRoutes() {
	router := gin.Default()
	router.POST("user", serv.createUser)
	router.POST("user/delete", serv.deleteUser)
	serv.router = router
}

func (serv *Server) Start(address string) error {
	return serv.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}