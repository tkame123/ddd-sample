package create_order_saga

import "github.com/tkame123/ddd-sample/app/order_api/domain/model"

type CreateOrderSagaState struct {
	Current CreateOrderSagaStep

	OrderID  model.OrderID
	TicketID model.TicketID
}

type CreateOrderSagaStep = string

const (
	CreateOrderSagaStep_ApprovalPending CreateOrderSagaStep = "ApprovalPending"
	CreateOrderSagaStep_CreatingTicket  CreateOrderSagaStep = "CreatingTicket"
	CreateOrderSagaStep_AuthorizingCard CreateOrderSagaStep = "AuthorizingCard"
	CreateOrderSagaStep_ApprovingTicket CreateOrderSagaStep = "ApprovingTicket"
	CreateOrderSagaStep_ApprovingOrder  CreateOrderSagaStep = "ApprovingOrder"
	CreateOrderSagaStep_OrderApproved   CreateOrderSagaStep = "OrderApproved"

	CreateOrderSagaStep_RejectingTicket CreateOrderSagaStep = "RejectingTicket"
	CreateOrderSagaStep_RejectingOrder  CreateOrderSagaStep = "RejectingOrder"
	CreateOrderSagaStep_OrderRejected   CreateOrderSagaStep = "OrderRejected"
)
