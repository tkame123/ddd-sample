package create_order_saga

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
)

type CreateOrderSaga struct {
	fsm        *fsm.FSM
	orderSVC   service.CreateOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI

	orderID  model.OrderID
	ticketID model.TicketID
}

func NewCreateOrderSaga(
	currentState *CreateOrderSagaState,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) (*CreateOrderSaga, error) {
	if currentState == nil {
		return nil, fmt.Errorf("current state is nil")
	}

	c := &CreateOrderSaga{
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,

		orderID:  currentState.OrderID,
		ticketID: currentState.TicketID,
	}

	ms := fsm.NewFSM(
		CreateOrderSagaStep_ApprovalPending,
		fsm.Events{
			{
				Name: message.Type_TYPE_EVENT_ORDER_CREATED.String(),
				Src:  []string{CreateOrderSagaStep_ApprovalPending},
				Dst:  CreateOrderSagaStep_CreatingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_CREATED.String(),
				Src:  []string{CreateOrderSagaStep_CreatingTicket},
				Dst:  CreateOrderSagaStep_AuthorizingCard,
			},
			{
				Name: message.Type_TYPE_EVENT_CARD_AUTHORIZED.String(),
				Src:  []string{CreateOrderSagaStep_AuthorizingCard},
				Dst:  CreateOrderSagaStep_ApprovingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_APPROVED.String(),
				Src:  []string{CreateOrderSagaStep_ApprovingTicket},
				Dst:  CreateOrderSagaStep_ApprovingOrder,
			},
			{
				Name: message.Type_TYPE_EVENT_ORDER_APPROVED.String(),
				Src:  []string{CreateOrderSagaStep_ApprovingOrder},
				Dst:  CreateOrderSagaStep_OrderApproved,
			},

			// Ticketの作成が失敗した場合
			{
				Name: message.Type_TYPE_EVENT_TICKET_CREATION_FAILED.String(),
				Src:  []string{CreateOrderSagaStep_CreatingTicket},
				Dst:  CreateOrderSagaStep_OrderRejected,
			},

			// オーソリ（仮売上）が失敗した場合
			{
				Name: message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED.String(),
				Src:  []string{CreateOrderSagaStep_AuthorizingCard},
				Dst:  CreateOrderSagaStep_RejectingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_REJECTED.String(),
				Src:  []string{CreateOrderSagaStep_RejectingTicket},
				Dst:  CreateOrderSagaStep_RejectingOrder,
			},
			{
				Name: message.Type_TYPE_EVENT_ORDER_REJECTED.String(),
				Src:  []string{CreateOrderSagaStep_RejectingOrder},
				Dst:  CreateOrderSagaStep_OrderRejected,
			},
		},
		fsm.Callbacks{
			"enter_CreatingTicket": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.createTicket(ctx)
			},
			"enter_ApprovingTicket": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.approveTicket(ctx)
			},
			fmt.Sprintf("before_%s", message.Type_TYPE_EVENT_TICKET_CREATED.String()): func(ctx context.Context, e *fsm.Event) {
				m, ok := e.Args[0].(*message.Message)
				if !ok {
					e.Err = fmt.Errorf("invalid message type")
					return
				}
				var v message.EventTicketCreated
				if err := m.Envelope.UnmarshalTo(&v); err != nil {
					e.Err = fmt.Errorf("unmarshal failed: %w", err)
					return
				}
				id, err := model.TicketIdParse(v.TicketId)
				if err != nil {
					e.Err = fmt.Errorf("parse ticket id failed: %w", err)
					return
				}

				e.Err = c.storeTicketID(ctx, *id)
			},
			"enter_AuthorizingCard": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.authorizeCard(ctx)
			},
			"enter_ApprovingOrder": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.approveOrder(ctx)
			},
			"enter_RejectingTicket": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.rejectTicket(ctx)
			},
			"enter_RejectingOrder": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.rejectOrder(ctx)
			},
		},
	)

	ms.SetState(currentState.Current)
	c.fsm = ms

	log.Println("CreateOrderSaga initialized and available transitions are:", c.fsm.AvailableTransitions())

	return c, nil
}

func (c *CreateOrderSaga) GetFSMVisualize() (string, error) {
	return fsm.VisualizeWithType(c.fsm, fsm.GRAPHVIZ)
}

func (c *CreateOrderSaga) CurrentStep() string {
	return c.fsm.Current()
}

func (c *CreateOrderSaga) Event(ctx context.Context, causeEvent *message.Message) error {
	if err := c.fsm.Event(ctx, causeEvent.Subject.Type.String(), causeEvent); err != nil {
		return err
	}
	return nil
}

func (c *CreateOrderSaga) ExportState() *CreateOrderSagaState {
	return &CreateOrderSagaState{
		Current:  c.fsm.Current(),
		OrderID:  c.orderID,
		TicketID: c.ticketID,
	}
}

func (c *CreateOrderSaga) createTicket(ctx context.Context) error {
	return c.kitchenAPI.CreateTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) approveTicket(ctx context.Context) error {
	return c.kitchenAPI.ApproveTicket(ctx, c.orderID, c.ticketID)
}

func (c *CreateOrderSaga) rejectTicket(ctx context.Context) error {
	return c.kitchenAPI.RejectTicket(ctx, c.orderID, c.ticketID)
}

func (c *CreateOrderSaga) storeTicketID(ctx context.Context, ticketID model.TicketID) error {
	c.ticketID = ticketID
	return nil
}

func (c *CreateOrderSaga) authorizeCard(ctx context.Context) error {
	return c.billingAPI.AuthorizeCard(ctx, c.orderID)
}

func (c *CreateOrderSaga) approveOrder(ctx context.Context) error {
	_, err := c.orderSVC.ApproveOrder(ctx, c.orderID)
	if err != nil {
		return fmt.Errorf("approve order failed: %w", err)
	}

	return nil
}

func (c *CreateOrderSaga) rejectOrder(ctx context.Context) error {
	_, err := c.orderSVC.RejectOrder(ctx, c.orderID)
	if err != nil {
		return fmt.Errorf("reject order failed: %w", err)
	}

	return nil
}
