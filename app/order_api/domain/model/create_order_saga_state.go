package model

type CreateOrderSagaState struct {
	OrderID  OrderID
	Current  CreateOrderSagaStep
	TicketID TicketID
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
