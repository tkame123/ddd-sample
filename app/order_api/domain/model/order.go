package model

import (
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/event"
)

const (
	APPROVAL_LIMIT int64 = 10000 // 10,000円
)

// 集約ルート
type Order struct {
	OrderID       OrderID
	approvalLimit int64
	orderItems    []*OrderItem
}

type OrderItem struct {
	OrderID  OrderID
	SortNo   int32
	ItemID   ItemID
	Price    int64
	Quantity int64
}

type OrderItemRequest struct {
	Item
	quantity int64
}

func NewOrder(items []*OrderItemRequest) (*Order, []event.OrderEvent, error) {
	orderID := generateID()

	orderItems := make([]*OrderItem, 0, len(items))
	for i, item := range items {
		orderItems = append(orderItems, &OrderItem{
			OrderID:  orderID,
			SortNo:   int32(i + 1),
			ItemID:   item.ItemID,
			Price:    item.Price,
			Quantity: item.quantity,
		})
	}

	order := &Order{
		OrderID:       orderID,
		approvalLimit: APPROVAL_LIMIT,
		orderItems:    orderItems,
	}

	if !order.validateApprovalLimit() {
		return nil, nil, errors.New("approval limit over")
	}

	createdEvent := event.NewOrderCreated(order.OrderID)
	return order, []event.OrderEvent{createdEvent}, nil
}

func (o *Order) UpdateOrderItems(items []*OrderItemRequest) ([]event.OrderEvent, error) {
	orderItems := make([]*OrderItem, 0, len(items))
	for i, item := range items {
		orderItems = append(orderItems, &OrderItem{
			OrderID:  o.OrderID,
			SortNo:   int32(i + 1),
			ItemID:   item.ItemID,
			Price:    item.Price,
			Quantity: item.quantity,
		})
	}

	o.orderItems = orderItems

	if !o.validateApprovalLimit() {
		return nil, errors.New("approval limit over")
	}

	updatedEvent := event.NewOrderItemsUpdated(o.OrderID)
	return []event.OrderEvent{updatedEvent}, nil
}

func (o *Order) validateApprovalLimit() bool {
	sum := 0
	for _, v := range o.orderItems {
		sum += int(v.Price * v.Quantity)
	}

	return sum <= int(o.approvalLimit)
}
