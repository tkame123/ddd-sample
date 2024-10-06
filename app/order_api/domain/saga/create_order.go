package saga

import (
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrder struct {
	fsm *fsm.FSM
}

func NewCreateOrder() *CreateOrder {
	eventTree := fsm.Events{
		{
			Name: "CreteTicket",
			Src:  []string{model.ApprovalPending.String()},
			Dst:  model.CreatingTicket.String(),
		},
		{
			Name: "AuthorizeCard",
			Src:  []string{model.CreatingTicket.String()},
			Dst:  model.AuthorizingCard.String(),
		},
		{
			Name: "ApproveTicket",
			Src:  []string{model.AuthorizingCard.String()},
			Dst:  model.ApprovingTicket.String(),
		},
		{
			Name: "ApproveOrder",
			Src:  []string{model.ApprovingTicket.String()},
			Dst:  model.ApprovingOrder.String(),
		},
		{
			Name: "OrderApprove",
			Src:  []string{model.ApprovingOrder.String()},
			Dst:  model.OrderApproved.String(),
		},

		// Ticketの作成が失敗した場合
		{
			Name: "TicketCreationFailed",
			Src:  []string{model.CreatingTicket.String()},
			Dst:  model.OrderRejected.String(),
		},

		// オーソリ（仮売上）が失敗した場合
		{
			Name: "AuthorizeCardFailed",
			Src:  []string{model.AuthorizingCard.String()},
			Dst:  model.RejectingTicket.String(),
		},
		{
			Name: "RejectOrder",
			Src:  []string{model.RejectingTicket.String()},
			Dst:  model.RejectingOrder.String(),
		},
		{
			Name: "RejectedOrder",
			Src:  []string{model.RejectingOrder.String()},
			Dst:  model.OrderRejected.String(),
		},
	}

	sm := fsm.NewFSM(
		model.ApprovalPending.String(),
		eventTree,
		fsm.Callbacks{},
	)

	return &CreateOrder{
		fsm: sm,
	}
}

func (c *CreateOrder) GetFSM() *fsm.FSM {
	return c.fsm
}
