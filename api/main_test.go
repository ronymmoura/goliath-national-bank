package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/goliath-national-bank/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server, _ := NewServer(store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
