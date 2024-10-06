package model

import (
	"context"
	"github.com/looplab/fsm"
)

type CreateOrderSagaState struct {
	orderID OrderID
	fsm     *fsm.FSM
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

type CreateOrderSagaEvent = string

const (
	CreateOrderSagaEvent_CreteTicket          CreateOrderSagaEvent = "CreteTicket"
	CreateOrderSagaEvent_AuthorizeCard        CreateOrderSagaEvent = "AuthorizeCard"
	CreateOrderSagaEvent_ApproveTicket        CreateOrderSagaEvent = "ApproveTicket"
	CreateOrderSagaEvent_ApproveOrder         CreateOrderSagaEvent = "ApproveOrder"
	CreateOrderSagaEvent_OrderApprove         CreateOrderSagaEvent = "OrderApprove"
	CreateOrderSagaEvent_TicketCreationFailed CreateOrderSagaEvent = "TicketCreationFailed"
	CreateOrderSagaEvent_AuthorizeCardFailed  CreateOrderSagaEvent = "AuthorizeCardFailed"
	CreateOrderSagaEvent_RejectOrder          CreateOrderSagaEvent = "RejectOrder"
	CreateOrderSagaEvent_RejectedOrder        CreateOrderSagaEvent = "RejectedOrder"
)

func NewCreateOrderSagaState(orderID OrderID, currentStep CreateOrderSagaStep) *CreateOrderSagaState {
	e := &CreateOrderSagaState{
		orderID: orderID,
	}

	ms := fsm.NewFSM(
		CreateOrderSagaStep_ApprovalPending,
		fsm.Events{
			{
				Name: CreateOrderSagaEvent_CreteTicket,
				Src:  []string{CreateOrderSagaStep_ApprovalPending},
				Dst:  CreateOrderSagaStep_CreatingTicket,
			},
			{
				Name: CreateOrderSagaEvent_AuthorizeCard,
				Src:  []string{CreateOrderSagaStep_CreatingTicket},
				Dst:  CreateOrderSagaStep_AuthorizingCard,
			},
			{
				Name: CreateOrderSagaEvent_ApproveTicket,
				Src:  []string{CreateOrderSagaStep_AuthorizingCard},
				Dst:  CreateOrderSagaStep_ApprovingTicket,
			},
			{
				Name: CreateOrderSagaEvent_ApproveOrder,
				Src:  []string{CreateOrderSagaStep_ApprovingTicket},
				Dst:  CreateOrderSagaStep_ApprovingOrder,
			},
			{
				Name: CreateOrderSagaEvent_OrderApprove,
				Src:  []string{CreateOrderSagaStep_ApprovingOrder},
				Dst:  CreateOrderSagaStep_OrderApproved,
			},

			// Ticketの作成が失敗した場合
			{
				Name: CreateOrderSagaEvent_TicketCreationFailed,
				Src:  []string{CreateOrderSagaStep_CreatingTicket},
				Dst:  CreateOrderSagaStep_OrderRejected,
			},

			// オーソリ（仮売上）が失敗した場合
			{
				Name: CreateOrderSagaEvent_AuthorizeCardFailed,
				Src:  []string{CreateOrderSagaStep_AuthorizingCard},
				Dst:  CreateOrderSagaStep_RejectingTicket,
			},
			{
				Name: CreateOrderSagaEvent_RejectOrder,
				Src:  []string{CreateOrderSagaStep_RejectingTicket},
				Dst:  CreateOrderSagaStep_RejectingOrder,
			},
			{
				Name: CreateOrderSagaEvent_RejectedOrder,
				Src:  []string{CreateOrderSagaStep_RejectingOrder},
				Dst:  CreateOrderSagaStep_OrderRejected,
			},
		},
		fsm.Callbacks{},
	)

	ms.SetState(currentStep)
	e.fsm = ms

	return e
}

func (c *CreateOrderSagaState) GetFSMVisualize() (string, error) {
	return fsm.VisualizeWithType(c.fsm, fsm.GRAPHVIZ)
}

func (c *CreateOrderSagaState) CreteTicket() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_CreteTicket)
}

func (c *CreateOrderSagaState) AuthorizeCard() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_AuthorizeCard)
}

func (c *CreateOrderSagaState) ApproveTicket() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_ApproveTicket)
}

func (c *CreateOrderSagaState) ApproveOrder() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_ApproveOrder)
}

func (c *CreateOrderSagaState) OrderApprove() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_OrderApprove)
}

func (c *CreateOrderSagaState) TicketCreationFailed() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_TicketCreationFailed)
}

func (c *CreateOrderSagaState) AuthorizeCardFailed() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_AuthorizeCardFailed)
}

func (c *CreateOrderSagaState) RejectOrder() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_RejectOrder)
}

func (c *CreateOrderSagaState) RejectedOrder() error {
	return c.fsm.Event(context.Background(), CreateOrderSagaEvent_RejectedOrder)
}
