package model

import (
	"errors"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/proto/message"
)

const APPOVAL_LIMIT = 10000

type OrderStatus = string

const (
	OrderStatus_ApprovalPending OrderStatus = "Pending"
	OrderStatus_Approved        OrderStatus = "Approved"
	OrderStatus_Rejected        OrderStatus = "Rejected"
	OrderStatus_CancelPending   OrderStatus = "CancelPending"
	OrderStatus_Canceled        OrderStatus = "Canceled"
)

// 集約ルート
type Order struct {
	OrderID    OrderID
	OrderItems []*OrderItem
	Status     OrderStatus
	Version    int
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

	mes, err := CreateMessage(
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
	if !o.StatusFSM().Can("authorized") {
		return nil, errors.New("order is not in approval pending status")
	}

	o.Status = OrderStatus_Approved

	mes, err := CreateMessage(
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
	if !o.StatusFSM().Can("rejected") {
		return nil, errors.New("order is not in rejected status")
	}

	o.Status = OrderStatus_Rejected

	mes, err := CreateMessage(
		&message.EventOrderRejected{
			OrderId: o.OrderID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}

func (o *Order) CancelOrder() ([]*message.Message, error) {
	if !o.StatusFSM().Can("cancel") {
		return nil, errors.New("order is not in cancel status")
	}

	o.Status = OrderStatus_CancelPending

	mes, err := CreateMessage(
		&message.EventOrderCanceled{
			OrderId: o.OrderID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}

func (o *Order) CancelConfirm() ([]*message.Message, error) {
	if !o.StatusFSM().Can("cancelConfirmed") {
		return nil, errors.New("order is not in cancel status")
	}

	o.Status = OrderStatus_Canceled

	mes, err := CreateMessage(
		&message.EventOrderCancellationConfirmed{
			OrderId: o.OrderID.String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []*message.Message{mes}, nil
}

func (o *Order) CancelReject() ([]*message.Message, error) {
	if !o.StatusFSM().Can("cancelRejected") {
		return nil, errors.New("order is not in cancel status")
	}

	o.Status = OrderStatus_Canceled

	mes, err := CreateMessage(
		&message.EventOrderCancellationRejected{
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

func (o *Order) StatusFSM() *fsm.FSM {
	ms := fsm.NewFSM(
		o.Status,
		fsm.Events{
			{
				Name: "authorized",
				Src:  []string{OrderStatus_ApprovalPending},
				Dst:  OrderStatus_Approved,
			},
			{
				Name: "rejected",
				Src:  []string{OrderStatus_ApprovalPending},
				Dst:  OrderStatus_Rejected,
			},
			{
				Name: "cancel",
				Src:  []string{OrderStatus_Approved},
				Dst:  OrderStatus_CancelPending,
			},
			{
				Name: "cancelConfirmed",
				Src:  []string{OrderStatus_CancelPending},
				Dst:  OrderStatus_Canceled,
			},
			{
				Name: "cancelRejected",
				Src:  []string{OrderStatus_CancelPending},
				Dst:  OrderStatus_Approved,
			},
		},
		fsm.Callbacks{},
	)
	return ms
}
