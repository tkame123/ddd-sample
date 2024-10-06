package model

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
		status:     ApprovalPending,
	}

	createdEvent := NewOrderCreated(order.OrderID())
	return order, []OrderEvent{createdEvent}, nil
}

func (o *Order) OrderID() OrderID {
	return o.orderID
}
