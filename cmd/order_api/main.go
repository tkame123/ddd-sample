package main

import (
	connect "connectrpc.com/connect"
	"context"
	"fmt"
	order_apiv1 "github.com/tkame123/ddd-sample/proto/order_api/v1"
	"github.com/tkame123/ddd-sample/proto/order_api/v1/order_apiv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := order_apiv1connect.NewOrderServiceHandler(&orderServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type orderServiceServer struct {
	order_apiv1connect.UnimplementedOrderServiceHandler
}

func (s *orderServiceServer) PutPet(
	ctx context.Context,
	req *connect.Request[order_apiv1.FindOrderRequest],
) (*connect.Response[order_apiv1.FindOrderResponse], error) {
	id := req.Msg.GetId()
	log.Printf("Got a request  a %v ", id)
	return connect.NewResponse(&order_apiv1.FindOrderResponse{}), nil
}
