package cancel_order_saga

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

type CancelOrderSaga struct {
	fsm        *fsm.FSM
	orderSVC   service.OrderService
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI

	orderID  model.OrderID
	ticketID model.TicketID
}

func NewCancelOrderSaga(
	currentState *CancelOrderSagaState,
	orderSVC service.OrderService,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) (*CancelOrderSaga, error) {
	if currentState == nil {
		return nil, fmt.Errorf("current state is nil")
	}

	c := &CancelOrderSaga{
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,

		orderID:  currentState.OrderID,
		ticketID: currentState.TicketID,
	}

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
			"enter_CancelingTicket": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.cancelTicket(ctx)
			},
			"enter_CancelingCard": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.cancelCard(ctx)
			},
			"enter_CancellationConfirmingOrder": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.cancelConfirmOrder(ctx)
			},

			"enter_CancellationRejectingOrder": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.cancelRejectOrder(ctx)
			},
		},
	)

	ms.SetState(currentState.Current)
	c.fsm = ms

	log.Println("CancelOrderSaga initialized and available transitions are:", c.fsm.AvailableTransitions())

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

func (c *CancelOrderSaga) ExportState() *CancelOrderSagaState {
	return &CancelOrderSagaState{
		Current:  c.fsm.Current(),
		OrderID:  c.orderID,
		TicketID: c.ticketID,
	}
}

func (c *CancelOrderSaga) cancelConfirmOrder(ctx context.Context) error {
	_, err := c.orderSVC.CancelConfirmOrder(ctx, c.orderID)
	if err != nil {
		return fmt.Errorf("cancel confirm order failed: %w", err)
	}

	return nil
}

func (c *CancelOrderSaga) cancelRejectOrder(ctx context.Context) error {
	_, err := c.orderSVC.CancelRejectOrder(ctx, c.orderID)
	if err != nil {
		return fmt.Errorf("cancel reject order failed: %w", err)
	}

	return nil
}

func (c *CancelOrderSaga) cancelTicket(ctx context.Context) error {
	return c.kitchenAPI.CancelTicket(ctx, c.orderID, c.ticketID)
}

func (c *CancelOrderSaga) cancelCard(ctx context.Context) error {
	return c.billingAPI.CancelCard(ctx, c.orderID)
}
