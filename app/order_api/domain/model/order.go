package model

import "errors"

// 集約ルート
type Order struct {
	orderID       OrderID
	approvalLimit int64
	orderItems    []*OrderItem
	status        OrderStatus
}

type OrderItem struct {
	OrderID  OrderID
	SortNo   int32
	ItemID   ItemID
	Price    int64
	Quantity int64
}

type OrderStatus = string

const (
	OrderStatus_ApprovalPending OrderStatus = "ApprovalPending"
	OrderStatus_OrderApproved   OrderStatus = "OrderApproved"
	OrderStatus_OrderRejected   OrderStatus = "OrderRejected"
)

type OrderItemRequest struct {
	Item
	quantity int64
}

func NewOrder(items []*OrderItemRequest) (*Order, []OrderEvent, error) {
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
		orderID:    orderID,
		orderItems: orderItems,
		status:     OrderStatus_ApprovalPending,
	}

	createdEvent := NewOrderCreatedEvent(order.OrderID())
	return order, []OrderEvent{createdEvent}, nil
}

func (o *Order) OrderID() OrderID {
	return o.orderID
}

func (o *Order) ApproveOrder() ([]OrderEvent, error) {
	if o.status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.status = OrderStatus_OrderApproved
	return []OrderEvent{NewOrderApprovedEvent(o.OrderID())}, nil
}

func (o *Order) RejectOrder() ([]OrderEvent, error) {
	if o.status != OrderStatus_ApprovalPending {
		return nil, errors.New("order is not in approval pending status")
	}

	o.status = OrderStatus_OrderRejected
	return []OrderEvent{NewOrderRejectedEvent(o.OrderID())}, nil
}
