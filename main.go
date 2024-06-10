package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ronymmoura/goliath-national-bank/api"
	db "github.com/ronymmoura/goliath-national-bank/db/sqlc"
	"github.com/ronymmoura/goliath-national-bank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	_, router := api.NewServer(store)

	err = router.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
