package create_order

import (
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

func (s *s) CreateOrder(ctx context.Context, items []*model.OrderItemRequest) (model.OrderID, error) {
	//order, events, err := model.NewOrder(items)
	//if err != nil {
	//	return uuid.Nil, err
	//}
	//
	//if err := s.rep.OrderSave(ctx, order); err != nil {
	//	return uuid.Nil, err
	//}
	//
	//if err := s.rep.CreateOrderSagaStateSave(ctx, &model.CreateOrderSagaState{
	//	OrderID: order.OrderID,
	//	Current: model.CreateOrderSagaStep_ApprovalPending,
	//}); err != nil {
	//	return uuid.Nil, err
	//}

	//s.pub.PublishMessages(ctx, events)

	//return order.OrderID, nil

	body := &message.EventOrderCreated{
		OrderId: uuid.New().String(),
	}

	x, err := anypb.New(body)
	if err != nil {
		log.Fatalln("failed to create anypb", err)
	}
	log.Println("x", x)

	a := message.Message{
		Subject: &message.Subject{
			Type:   message.Type_TYPE_EVENT_ORDER_CREATED,
			Source: message.Service_SERVICE_ORDER,
		},
		Envelope: x,
	}

	s.pub.PublishMessages(ctx, []*message.Message{&a})

	return uuid.Nil, nil
}
