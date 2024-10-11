package main

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"log"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// WARN: for development environment only
func main() {
	envCfg, err := provider.NewENV()
	if err != nil {
		log.Fatalf("failed loading env: %v", err)
	}
	client, closeFunc, err := provider.NewOrderApiDB(envCfg)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer closeFunc()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
