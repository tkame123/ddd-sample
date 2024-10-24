package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
)

type s struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewOrderService(rep repository.Repository, pub domain_event.Publisher) service.OrderService {
	return &s{
		rep: rep,
		pub: pub,
	}
}

func (s *s) CreateOrder(ctx context.Context, items []*model.OrderItemRequest) (model.OrderID, error) {
	order, events, err := model.NewOrder(items)
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.CreateOrderSagaStateSave(ctx, &create_order_saga.CreateOrderSagaState{
		OrderID: order.OrderID,
		Current: create_order_saga.CreateOrderSagaStep_ApprovalPending,
	}); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}

func (s *s) ApproveOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return uuid.Nil, err
	}

	events, err := order.ApproveOrder()
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}

func (s *s) RejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return uuid.Nil, err
	}

	events, err := order.RejectOrder()
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}

func (s *s) CancelOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return uuid.Nil, err
	}

	events, err := order.CancelOrder()
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.CancelOrderSagaStateSave(ctx, &cancel_order_saga.CancelOrderSagaState{
		OrderID: order.OrderID,
		Current: cancel_order_saga.CancelOrderSagaStep_CancelPending,
	}); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}

func (s *s) CancelConfirmOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return uuid.Nil, err
	}

	events, err := order.CancelConfirm()
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}

func (s *s) CancelRejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error) {
	order, err := s.rep.OrderFindOne(ctx, orderID)
	if err != nil {
		return uuid.Nil, err
	}

	events, err := order.CancelReject()
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.rep.OrderSave(ctx, order); err != nil {
		return uuid.Nil, err
	}

	s.pub.PublishMessages(ctx, events)

	return order.OrderID, nil
}
