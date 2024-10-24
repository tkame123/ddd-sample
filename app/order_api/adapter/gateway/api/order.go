package connect

import (
	"connectrpc.com/connect"
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/usecase"
	"github.com/tkame123/ddd-sample/proto/order_api/v1/order_apiv1connect"

	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	order_apiv1 "github.com/tkame123/ddd-sample/proto/order_api/v1"
)

type orderServiceServer struct {
	order_apiv1connect.UnimplementedOrderServiceHandler

	rep repository.Repository
	pub domain_event.Publisher
}

func (s *orderServiceServer) CreateOrder(
	ctx context.Context, req *connect.Request[order_apiv1.CreateOrderRequest],
) (*connect.Response[order_apiv1.CreateOrderResponse], error) {
	svc := usecase.NewOrderService(s.rep, s.pub)

	items, err := toModelOrderItemRequest(req.Msg.GetItems())
	if err != nil {
		return nil, err
	}

	orderId, err := svc.CreateOrder(ctx, items)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&order_apiv1.CreateOrderResponse{
		Id: orderId.String(),
	}), nil
}

func (s *orderServiceServer) CancelOrder(
	ctx context.Context, req *connect.Request[order_apiv1.CancelOrderRequest],
) (*connect.Response[order_apiv1.CancelOrderResponse], error) {
	id := req.Msg.GetId()
	parsedId, err := model.OrderIdParse(id)
	if err != nil {
		return nil, err
	}

	svc := usecase.NewOrderService(s.rep, s.pub)

	orderId, err := svc.CancelOrder(ctx, *parsedId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&order_apiv1.CancelOrderResponse{
		Id: orderId.String(),
	}), nil
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

func toModelOrderItemRequest(items []*order_apiv1.CreateOrderRequest_RequestItem) ([]*model.OrderItemRequest, error) {
	orderItemRequests := make([]*model.OrderItemRequest, 0, len(items))
	for _, item := range items {
		id, err := model.ItemIdParse(item.ProductId)
		if err != nil {
			return nil, err
		}
		orderItemRequests = append(orderItemRequests, &model.OrderItemRequest{
			Item: model.Item{
				ItemID: *id,
				Price:  int(item.Price),
			},
			Quantity: int(item.Quantity),
		})
	}
	return orderItemRequests, nil
}
