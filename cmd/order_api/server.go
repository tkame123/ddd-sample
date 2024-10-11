package main

import (
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/di"
)

func main() {
	server, cleanup, err := di.InitializeOrderAPIServer()
	defer cleanup()
	if err != nil {
		panic(err)
	}
	server.Run()
}
