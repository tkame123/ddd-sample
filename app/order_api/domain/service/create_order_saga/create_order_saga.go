package create_order_saga

import (
	"context"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
)

type CreateOrderSaga struct {
	currentState *model.CreateOrderSagaState
	fsm          *fsm.FSM
	rep          repository.Repository
	orderSVC     service.CreateOrder
	kitchenAPI   external_service.KitchenAPI
	billingAPI   external_service.BillingAPI
}

func NewCreateOrderSaga(
	currentState *model.CreateOrderSagaState,
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *CreateOrderSaga {
	c := &CreateOrderSaga{
		currentState: currentState,
		rep:          rep,
		orderSVC:     orderSVC,
		kitchenAPI:   kitchenAPI,
		billingAPI:   billingAPI,
	}

	ms := fsm.NewFSM(
		model.CreateOrderSagaStep_ApprovalPending,
		fsm.Events{
			{
				Name: message.Type_TYPE_EVENT_ORDER_CREATED.String(),
				Src:  []string{model.CreateOrderSagaStep_ApprovalPending},
				Dst:  model.CreateOrderSagaStep_CreatingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_CREATED.String(),
				Src:  []string{model.CreateOrderSagaStep_CreatingTicket},
				Dst:  model.CreateOrderSagaStep_AuthorizingCard,
			},
			{
				Name: message.Type_TYPE_EVENT_CARD_AUTHORIZED.String(),
				Src:  []string{model.CreateOrderSagaStep_AuthorizingCard},
				Dst:  model.CreateOrderSagaStep_ApprovingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_APPROVED.String(),
				Src:  []string{model.CreateOrderSagaStep_ApprovingTicket},
				Dst:  model.CreateOrderSagaStep_ApprovingOrder,
			},
			{
				Name: message.Type_TYPE_EVENT_ORDER_APPROVED.String(),
				Src:  []string{model.CreateOrderSagaStep_ApprovingOrder},
				Dst:  model.CreateOrderSagaStep_OrderApproved,
			},

			// Ticketの作成が失敗した場合
			{
				Name: message.Type_TYPE_EVENT_TICKET_CREATION_FAILED.String(),
				Src:  []string{model.CreateOrderSagaStep_CreatingTicket},
				Dst:  model.CreateOrderSagaStep_OrderRejected,
			},

			// オーソリ（仮売上）が失敗した場合
			{
				Name: message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED.String(),
				Src:  []string{model.CreateOrderSagaStep_AuthorizingCard},
				Dst:  model.CreateOrderSagaStep_RejectingTicket,
			},
			{
				Name: message.Type_TYPE_EVENT_TICKET_REJECTED.String(),
				Src:  []string{model.CreateOrderSagaStep_RejectingTicket},
				Dst:  model.CreateOrderSagaStep_RejectingOrder,
			},
			{
				Name: message.Type_TYPE_EVENT_ORDER_REJECTED.String(),
				Src:  []string{model.CreateOrderSagaStep_RejectingOrder},
				Dst:  model.CreateOrderSagaStep_OrderRejected,
			},
		},
		fsm.Callbacks{
			"enter_CreatingTicket": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.createTicket(ctx)
			},
			"enter_ApprovingTicket": func(ctx context.Context, e *fsm.Event) {
				e.Err = c.approveTicket(ctx)
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

	return c
}

func (c *CreateOrderSaga) GetFSMVisualize() (string, error) {
	return fsm.VisualizeWithType(c.fsm, fsm.GRAPHVIZ)
}

func (c *CreateOrderSaga) CurrentStep() string {
	return c.fsm.Current()
}

func (c *CreateOrderSaga) Event(ctx context.Context, causeEvent *message.Message) error {
	if err := c.fsm.Event(ctx, causeEvent.Subject.Type.String()); err != nil {
		return err
	}
	c.currentState.ApplyStep(c.fsm.Current())
	if err := c.rep.CreateOrderSagaStateSave(ctx, c.currentState); err != nil {
		return err
	}
	log.Printf("CreateOrderSaga steped: OrderID: %s, CurrentStep: %s", c.currentState.OrderID, c.fsm.Current())
	return nil
}

func (c *CreateOrderSaga) createTicket(ctx context.Context) error {
	c.kitchenAPI.CreateTicket(ctx, c.currentState.OrderID)
	return nil
}

func (c *CreateOrderSaga) approveTicket(ctx context.Context) error {
	c.kitchenAPI.ApproveTicket(ctx, c.currentState.OrderID)
	return nil
}

func (c *CreateOrderSaga) rejectTicket(ctx context.Context) error {
	c.kitchenAPI.RejectTicket(ctx, c.currentState.OrderID)
	return nil
}

func (c *CreateOrderSaga) authorizeCard(ctx context.Context) error {
	c.billingAPI.AuthorizeCard(ctx, c.currentState.OrderID)
	return nil
}

func (c *CreateOrderSaga) approveOrder(ctx context.Context) error {
	_, err := c.orderSVC.ApproveOrder(ctx, c.currentState.OrderID)

	if err != nil {
		// TODO: 原則再実行で成功が保証されているはずなので、通知だけ行う
		panic("approve order failed")
	}

	return nil
}

func (c *CreateOrderSaga) rejectOrder(ctx context.Context) error {
	_, err := c.orderSVC.RejectOrder(ctx, c.currentState.OrderID)

	if err != nil {
		// TODO: 原則再実行で成功が保証されているはずなので、通知だけ行う
		panic("reject order failed")
	}

	return nil
}
