package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gyu-young-park/go_blog/db"
	"github.com/gyu-young-park/go_blog/token"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.TokenMaker
}

func NewServer(store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker("123456789012345678912345678912345")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	server.router = gin.Default()
	server.userRouterSetting()
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func sendErrorMessage(status int, message string, err error, c *gin.Context) {
	log.Printf("failed to [%v] [%v]\n", message, err.Error())
	c.JSON(status, errResponse(err))
}
