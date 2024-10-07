package servive

import (
	"context"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type CreateOrderSaga struct {
	orderID     model.OrderID
	fsm         *fsm.FSM
	rep         *repository.Repository
	pub         domain_event.Publisher
	externalAPI *external_service.ExternalAPI
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
	rep *repository.Repository,
	pub domain_event.Publisher,
	externalAPI *external_service.ExternalAPI,
) *CreateOrderSaga {
	c := &CreateOrderSaga{
		orderID:     currentState.OrderID(),
		rep:         rep,
		pub:         pub,
		externalAPI: externalAPI,
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
			"enter_CreteTicket": func(ctx context.Context, e *fsm.Event) {
				c.createTicket(ctx)
			},
			"enter_ApproveTicket": func(ctx context.Context, e *fsm.Event) {
				c.approveTicket(ctx)
			},
			"enter_AuthorizeCard": func(ctx context.Context, e *fsm.Event) {
				c.authorizeCard(ctx)
			},
			"enter_ApproveOrder": func(ctx context.Context, e *fsm.Event) {
				c.approveOrder(ctx)
			},
			"enter_RejectOrder": func(ctx context.Context, e *fsm.Event) {
				c.rejectOrder(ctx)
			},
		},
	)

	ms.SetState(currentState.Current())
	c.fsm = ms

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
	if err := c.rep.CreateOrderSagaState.Save(ctx, model.NewCreateOrderSagaState(c.orderID, c.fsm.Current())); err != nil {
		return err
	}
	return nil
}

func (c *CreateOrderSaga) createTicket(ctx context.Context) {
	c.externalAPI.KitchenAPI.CreateTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) approveTicket(ctx context.Context) {
	c.externalAPI.KitchenAPI.ApproveTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) authorizeCard(ctx context.Context) {
	c.externalAPI.KitchenAPI.ApproveTicket(ctx, c.orderID)
}

func (c *CreateOrderSaga) approveOrder(ctx context.Context) {
	c.pub.PublishMessages(ctx, []model.OrderEvent{model.NewApproveOrderCommand(c.orderID)})
}

func (c *CreateOrderSaga) rejectOrder(ctx context.Context) {
	c.pub.PublishMessages(ctx, []model.OrderEvent{model.NewRejectOrderCommand(c.orderID)})
}
