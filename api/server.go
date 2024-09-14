package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/simplebank/api/routes"
	db "github.com/simplebank/db/sqlc"
	"github.com/simplebank/token"
	"github.com/simplebank/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	routes.AccountRoutes(
		router,
		server.CreateAccount,
		server.GetAccount,
		server.ListAccount,
		server.UpdateAccount,
		server.DeleteAccount,
	)
	routes.TransferRoutes(router, server.createTransfer)

	routes.UserRoutes(router, server.createUser, server.loginUser)
	server.router = router
}

func (server Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func deleteResponse() gin.H {
	return gin.H{"message": "the record got deleted successfully"}
}
