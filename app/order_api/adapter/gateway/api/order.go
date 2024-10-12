package connect

import (
	"connectrpc.com/connect"
	"context"

	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	order_apiv1 "github.com/tkame123/ddd-sample/proto/order_api/v1"
	"github.com/tkame123/ddd-sample/proto/order_api/v1/order_apiv1connect"
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
	parsedId, err := model.OrderIdParse(id)
	if err != nil {
		return nil, err
	}
	order, err := s.rep.OrderFindOne(ctx, *parsedId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&order_apiv1.FindOrderResponse{
		Order: &order_apiv1.Order{
			Id: order.OrderID.String(),
		},
	}), nil
}
