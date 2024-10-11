package provider

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"log"
)

func NewOrderApiDB(env *EnvConfig) (*ent.Client, func(), error) {
	client, err := ent.Open("postgres", env.OrderAPIDSN)
	if err != nil {
		return nil, nil, fmt.Errorf("failed opening connection to mysql: %w", err)
	}
	cleanup := func() {
		if err = client.Close(); err != nil {
			log.Printf("failed closing connection to mysql: %v\n", err)
		}
	}

	return client, cleanup, nil
}
