package model

type CreateOrderSagaState struct {
	orderID OrderID
	current CreateOrderSagaStep
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

func NewCreateOrderSagaState(orderID OrderID, currentStep CreateOrderSagaStep) *CreateOrderSagaState {
	return &CreateOrderSagaState{
		orderID: orderID,
		current: currentStep,
	}
}

func (s *CreateOrderSagaState) OrderID() OrderID {
	return s.orderID
}

func (s *CreateOrderSagaState) Current() CreateOrderSagaStep {
	return s.current
}
