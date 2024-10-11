package provider

import (
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"log"
)

func NewOrderApiDB(env *EnvConfig) (*ent.Client, func(), error) {
	client, err := ent.Open("postgres", env.OrderAPIDSN)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = client.Close(); err != nil {
			log.Println(err)
		}
	}

	return client, cleanup, nil
}
