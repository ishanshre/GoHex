package main

import (
	"log"
	"os"

	"github.com/ishanshre/GoHex/internal/adapters/app/api"
	"github.com/ishanshre/GoHex/internal/adapters/core/arithmetic"
	rpc "github.com/ishanshre/GoHex/internal/adapters/framework/left/grpc"
	"github.com/ishanshre/GoHex/internal/adapters/framework/right/db"
)

func main() {

	dbaseDriver := os.Getenv("dbaseDriver")
	dsn := os.Getenv("dsn")
	dbAdapter, err := db.NewAdapter(dbaseDriver, dsn)
	if err != nil {
		log.Fatalf("error in connecting to database: %v", err)
	}
	defer dbAdapter.CloseDbConnection()
	core := arithmetic.NewAdapter()
	applicationAPI := api.NewAdapter(dbAdapter, core)
	grpcAdapter := rpc.NewAdapater(applicationAPI)
	grpcAdapter.Run()
}
