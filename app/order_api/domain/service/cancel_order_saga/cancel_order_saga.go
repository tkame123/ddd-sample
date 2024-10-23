package cancel_order_saga

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
)

type CancelOrderSaga struct {
	fsm *fsm.FSM
}

func NewCancelOrderSaga(
	currentState *CancelOrderSagaState,
) (*CancelOrderSaga, error) {
	if currentState == nil {
		return nil, fmt.Errorf("current state is nil")
	}

	c := &CancelOrderSaga{}

	ms := fsm.NewFSM(
		CancelOrderSagaStep_CancelPending,
		fsm.Events{
			{
				Name: message.Type_TYPE_EVENT_ORDER_CANCELED.String(),
				Src:  []string{CancelOrderSagaStep_CancelPending},
				Dst:  CancelOrderSagaStep_CancelingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_CANCELED.String(),
				Src:  []string{CancelOrderSagaStep_CancelingTicket},
				Dst:  CancelOrderSagaStep_CancelingCard,
			},
			{
				Name: message.Type_TYPE_EVENT_CARD_CANCELED.String(),
				Src:  []string{CancelOrderSagaStep_CancelingCard},
				Dst:  CancelOrderSagaStep_CancellationConfirmingOrder,
			},
			{
				Name: message.Type_TYPE_EVENT_ORDER_CANCELLATION_CONFIRMED.String(),
				Src:  []string{CancelOrderSagaStep_CancellationConfirmingOrder},
				Dst:  CancelOrderSagaStep_OrderCanceled,
			},

			// Ticketのキャンセルが拒否された場合
			{
				Name: message.Type_TYPE_EVENT_TICKET_CANCELLATION_REJECTED.String(),
				Src:  []string{CancelOrderSagaStep_CancelingTicket},
				Dst:  CancelOrderSagaStep_CancellationRejectingOrder,
			},
			{
				Name: message.Type_TYPE_EVENT_ORDER_CANCELLATION_REJECTED.String(),
				Src:  []string{CancelOrderSagaStep_CancellationRejectingOrder},
				Dst:  CancelOrderSagaStep_OrderCancellationRejected,
			},

			// CARDの仮売上のキャンセルの失敗は想定されない
		},
		fsm.Callbacks{
			//"enter_CreatingTicket": func(ctx context.Context, e *fsm.Event) {
			//	e.Err = c.createTicket(ctx)
			//},
			//"enter_ApprovingTicket": func(ctx context.Context, e *fsm.Event) {
			//	e.Err = c.approveTicket(ctx)
			//},
			//fmt.Sprintf("before_%s", message.Type_TYPE_EVENT_TICKET_CREATED.String()): func(ctx context.Context, e *fsm.Event) {
			//	log.Println("before_AuthorizingCard")
			//	m, ok := e.Args[0].(*message.Message)
			//	if !ok {
			//		e.Err = fmt.Errorf("invalid message type")
			//		return
			//	}
			//	var v message.EventTicketCreated
			//	if err := m.Envelope.UnmarshalTo(&v); err != nil {
			//		e.Err = fmt.Errorf("unmarshal failed: %w", err)
			//		return
			//	}
			//	id, err := model.TicketIdParse(v.TicketId)
			//	if err != nil {
			//		e.Err = fmt.Errorf("parse ticket id failed: %w", err)
			//		return
			//	}
			//
			//	e.Err = c.storeTicketID(ctx, *id)
			//},
			//"enter_AuthorizingCard": func(ctx context.Context, e *fsm.Event) {
			//	e.Err = c.authorizeCard(ctx)
			//},
			//"enter_ApprovingOrder": func(ctx context.Context, e *fsm.Event) {
			//	e.Err = c.approveOrder(ctx)
			//},
			//"enter_RejectingTicket": func(ctx context.Context, e *fsm.Event) {
			//	e.Err = c.rejectTicket(ctx)
			//},
			//"enter_RejectingOrder": func(ctx context.Context, e *fsm.Event) {
			//	e.Err = c.rejectOrder(ctx)
			//},
		},
	)

	ms.SetState(currentState.Current)
	c.fsm = ms

	log.Println("CreateOrderSaga initialized and available transitions are:", c.fsm.AvailableTransitions())

	return c, nil
}

func (c *CancelOrderSaga) GetFSMVisualize() (string, error) {
	return fsm.VisualizeWithType(c.fsm, fsm.GRAPHVIZ)
}

func (c *CancelOrderSaga) CurrentStep() string {
	return c.fsm.Current()
}

func (c *CancelOrderSaga) Event(ctx context.Context, causeEvent *message.Message) error {
	if err := c.fsm.Event(ctx, causeEvent.Subject.Type.String(), causeEvent); err != nil {
		return err
	}
	return nil
}

//func (c *CancelOrderSaga) ExportState() *CancelOrderSagaState {
//	return &CancelOrderSagaState{
//		Current:  c.fsm.Current(),
//		OrderID:  c.orderID,
//		TicketID: c.ticketID,
//	}
//}

//func (c *CreateOrderSaga) createTicket(ctx context.Context) error {
//	return c.kitchenAPI.CreateTicket(ctx, c.orderID)
//}
//
//func (c *CreateOrderSaga) approveTicket(ctx context.Context) error {
//	return c.kitchenAPI.ApproveTicket(ctx, c.orderID, c.ticketID)
//}
//
//func (c *CreateOrderSaga) rejectTicket(ctx context.Context) error {
//	return c.kitchenAPI.RejectTicket(ctx, c.orderID, c.ticketID)
//}
//
//func (c *CreateOrderSaga) storeTicketID(ctx context.Context, ticketID model.TicketID) error {
//	c.ticketID = ticketID
//	return nil
//}
//
//func (c *CreateOrderSaga) authorizeCard(ctx context.Context) error {
//	return c.billingAPI.AuthorizeCard(ctx, c.orderID)
//}
//
//func (c *CreateOrderSaga) approveOrder(ctx context.Context) error {
//	_, err := c.orderSVC.ApproveOrder(ctx, c.orderID)
//	if err != nil {
//		return fmt.Errorf("approve order failed: %w", err)
//	}
//
//	return nil
//}
//
//func (c *CreateOrderSaga) rejectOrder(ctx context.Context) error {
//	_, err := c.orderSVC.RejectOrder(ctx, c.orderID)
//	if err != nil {
//		return fmt.Errorf("reject order failed: %w", err)
//	}
//
//	return nil
//}
