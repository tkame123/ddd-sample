package main

import (
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/kitchen_api/di"
)

func main() {
	server, cleanup, err := di.InitializeCommandConsumer()
	defer cleanup()
	if err != nil {
		panic(err)
	}
	server.Run()
}
