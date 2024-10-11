package create_order_saga

import (
	"context"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"log"
)

type CreateOrderSaga struct {
	orderID    model.OrderID
	fsm        *fsm.FSM
	rep        repository.Repository
	orderSVC   service.CreateOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

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

func NewCreateOrderSaga(
	currentState *model.CreateOrderSagaState,
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *CreateOrderSaga {
	c := &CreateOrderSaga{
		orderID:    currentState.OrderID,
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}

	ms := fsm.NewFSM(
		model.CreateOrderSagaStep_ApprovalPending,
		fsm.Events{
			{
				Name: CreateOrderSagaEvent_CreteTicket,
				Src:  []string{model.CreateOrderSagaStep_ApprovalPending},
				Dst:  model.CreateOrderSagaStep_CreatingTicket,
			},
			{
				Name: CreateOrderSagaEvent_AuthorizeCard,
				Src:  []string{model.CreateOrderSagaStep_CreatingTicket},
				Dst:  model.CreateOrderSagaStep_AuthorizingCard,
			},
			{
				Name: CreateOrderSagaEvent_ApproveTicket,
				Src:  []string{model.CreateOrderSagaStep_AuthorizingCard},
				Dst:  model.CreateOrderSagaStep_ApprovingTicket,
			},
			{
				Name: CreateOrderSagaEvent_ApproveOrder,
				Src:  []string{model.CreateOrderSagaStep_ApprovingTicket},
				Dst:  model.CreateOrderSagaStep_ApprovingOrder,
			},
			{
				Name: CreateOrderSagaEvent_OrderApprove,
				Src:  []string{model.CreateOrderSagaStep_ApprovingOrder},
				Dst:  model.CreateOrderSagaStep_OrderApproved,
			},

			// Ticketの作成が失敗した場合
			{
				Name: CreateOrderSagaEvent_TicketCreationFailed,
				Src:  []string{model.CreateOrderSagaStep_CreatingTicket},
				Dst:  model.CreateOrderSagaStep_OrderRejected,
			},

			// オーソリ（仮売上）が失敗した場合
			{
				Name: CreateOrderSagaEvent_AuthorizeCardFailed,
				Src:  []string{model.CreateOrderSagaStep_AuthorizingCard},
				Dst:  model.CreateOrderSagaStep_RejectingTicket,
			},
			{
				Name: CreateOrderSagaEvent_RejectOrder,
				Src:  []string{model.CreateOrderSagaStep_RejectingTicket},
				Dst:  model.CreateOrderSagaStep_RejectingOrder,
			},
			{
				Name: CreateOrderSagaEvent_RejectedOrder,
				Src:  []string{model.CreateOrderSagaStep_RejectingOrder},
				Dst:  model.CreateOrderSagaStep_OrderRejected,
			},
		},
		fsm.Callbacks{
			"enter_CreatingTicket": func(ctx context.Context, e *fsm.Event) {
				c.createTicket(ctx)
			},
			"enter_ApprovingTicket": func(ctx context.Context, e *fsm.Event) {
				c.approveTicket(ctx)
			},
			"enter_AuthorizingCard": func(ctx context.Context, e *fsm.Event) {
				c.authorizeCard(ctx)
			},
			"enter_ApprovingOrder": func(ctx context.Context, e *fsm.Event) {
				c.approveOrder(ctx)
			},
			"enter_RejectingTicket": func(ctx context.Context, e *fsm.Event) {
				c.rejectTicket(ctx)
			},
			"enter_RejectingOrder": func(ctx context.Context, e *fsm.Event) {
				c.rejectOrder(ctx)
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

func (c *CreateOrderSaga) Event(ctx context.Context, causeEvent CreateOrderSagaEvent) error {
	if err := c.fsm.Event(ctx, causeEvent); err != nil {
		return err
	}
	if err := c.rep.CreateOrderSagaStateSave(ctx, model.NewCreateOrderSagaState(c.orderID, c.fsm.Current())); err != nil {
		return err
	}
	return nil
}

func (c *CreateOrderSaga) createTicket(ctx context.Context) {
	c.kitchenAPI.CreateTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) approveTicket(ctx context.Context) {
	c.kitchenAPI.ApproveTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) rejectTicket(ctx context.Context) {
	c.kitchenAPI.RejectTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) authorizeCard(ctx context.Context) {
	c.billingAPI.AuthorizeCard(ctx, c.orderID)
}

func (c *CreateOrderSaga) approveOrder(ctx context.Context) {
	_, err := c.orderSVC.ApproveOrder(ctx, c.orderID)

	if err != nil {
		// TODO: 原則再実行で成功が保証されているはずなので、通知だけ行う
		panic("approve order failed")
	}
}

func (c *CreateOrderSaga) rejectOrder(ctx context.Context) {
	_, err := c.orderSVC.RejectOrder(ctx, c.orderID)

	if err != nil {
		// TODO: 原則再実行で成功が保証されているはずなので、通知だけ行う
		panic("reject order failed")
	}
}
