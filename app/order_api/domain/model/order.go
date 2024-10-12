package model

import (
	"errors"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
)

const APPOVAL_LIMIT = 10000

type OrderStatus = string

const (
	OrderStatus_ApprovalPending OrderStatus = "Pending"
	OrderStatus_Approved        OrderStatus = "Approved"
	OrderStatus_Rejected        OrderStatus = "Rejected"
)

// 集約ルート
type Order struct {
	OrderID    OrderID
	OrderItems []*OrderItem
	Status     OrderStatus
}

type OrderItem struct {
	OrderItemID OrderItemID
	OrderID     OrderID
	SortNo      int
	ItemID      ItemID
	Price       int
	Quantity    int
}

type OrderItemRequest struct {
	Item
	Quantity int
}

func NewOrder(items []*OrderItemRequest) (*Order, []*message.Message, error) {
	orderID := generateID()

	orderItems := make([]*OrderItem, 0, len(items))
	for i, item := range items {
		orderItems = append(orderItems, &OrderItem{
			OrderItemID: generateID(),
			OrderID:     orderID,
			SortNo:      i + 1,
			ItemID:      item.ItemID,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}

	order := &Order{
		OrderID:    orderID,
		OrderItems: orderItems,
		Status:     OrderStatus_ApprovalPending,
	}

	if !order.validateApprovalLimit() {
		return nil, nil, errors.New("approval limit over")
	}

	mes, err := event_helper.CreateMessage(
		message.Type_TYPE_EVENT_ORDER_CREATED,
		message.Service_SERVICE_ORDER,
		&message.EventOrderCreated{
			OrderId: order.OrderID.String(),
			// TODO: 詳細のアイテム対応
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return order, []*message.Message{mes}, nil
}

func (o *Order) ApproveOrder() ([]*message.Message, error) {
	if o.Status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.Status = OrderStatus_Approved

	mes, err := event_helper.CreateMessage(
		message.Type_TYPE_EVENT_ORDER_APPROVED,
		message.Service_SERVICE_ORDER,
		&message.EventOrderApproved{
			OrderId: o.OrderID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}

func (o *Order) RejectOrder() ([]*message.Message, error) {
	if o.Status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.Status = OrderStatus_Rejected

	mes, err := event_helper.CreateMessage(
		message.Type_TYPE_EVENT_ORDER_REJECTED,
		message.Service_SERVICE_ORDER,
		&message.EventOrderRejected{
			OrderId: o.OrderID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}

func (o *Order) validateApprovalLimit() bool {
	sum := 0
	for _, v := range o.OrderItems {
		sum += v.Price * v.Quantity
	}

	return sum <= APPOVAL_LIMIT
}
