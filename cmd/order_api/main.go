package main

import (
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	connect "github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/api"
	"log"
)

func main() {
	//client, err := ent.Open("mysql", "root:Number@5@tcp(localhost:3306)/ddl_sample?parseTime=True")
	client, err := ent.Open("postgres", "host=localhost port=5432 user=root dbname=ddl_sample password=Number@5 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	repo := database.NewRepository(client)

	server := connect.NewServer(repo)
	server.Run()
}
