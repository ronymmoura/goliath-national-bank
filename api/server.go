package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/ronymmoura/goliath-national-bank/db/sqlc"
	"github.com/ronymmoura/goliath-national-bank/token"
	"github.com/ronymmoura/goliath-national-bank/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (server *Server, err error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server = &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.GET("users/:user_id", server.getUser)
	router.POST("users", server.createUser)

	router.GET("accounts", server.listAccounts)
	router.GET("accounts/:account_id", server.getAccount)
	router.POST("accounts", server.createAccount)

	router.POST("transfers", server.createTransfer)

	server.router = router
	return
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
