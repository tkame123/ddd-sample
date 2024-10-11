package connect

import (
	"connectrpc.com/connect"
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	order_apiv1 "github.com/tkame123/ddd-sample/proto/order_api/v1"
	"github.com/tkame123/ddd-sample/proto/order_api/v1/order_apiv1connect"
	"log"
)

type orderServiceServer struct {
	rep repository.Repository
	order_apiv1connect.UnimplementedOrderServiceHandler
}

func (s *orderServiceServer) FindOrder(
	ctx context.Context,
	req *connect.Request[order_apiv1.FindOrderRequest],
) (*connect.Response[order_apiv1.FindOrderResponse], error) {
	id := req.Msg.GetId()
	log.Printf("Got a request  a %v ", id)
	order, err := s.rep.OrderFindOne(ctx, uuid.Nil)
	if err != nil {
		return nil, err
	}
	log.Printf("Got a order %v ", order)
	return connect.NewResponse(&order_apiv1.FindOrderResponse{}), nil
}
