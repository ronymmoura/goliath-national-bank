package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/ronymmoura/goliath-national-bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) (server *Server, router *gin.Engine) {
	server = &Server{store: store}
	router = gin.Default()

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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
