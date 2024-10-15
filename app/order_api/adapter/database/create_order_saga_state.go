package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/createordersagastate"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

func (r *repo) CreateOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*model.CreateOrderSagaState, error) {
	state, err := r.db.CreateOrderSagaState.Query().
		Where(createordersagastate.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return toModelOrderSagaState(state), nil
}

func (r *repo) CreateOrderSagaStateSave(ctx context.Context, state *model.CreateOrderSagaState) error {
	err := r.db.CreateOrderSagaState.Create().
		SetID(state.OrderID).
		SetCurrent(fromModelCurrentState(state.Current)).
		SetNillableTicketID(fromModelTicketID(state.TicketID)).
		OnConflictColumns("id").
		UpdateTicketID().
		UpdateCurrent().
		UpdateUpdatedAt().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func fromModelTicketID(ticketID uuid.NullUUID) *uuid.UUID {
	if ticketID.Valid {
		return &ticketID.UUID
	}
	return nil
}

func fromModelCurrentState(state model.CreateOrderSagaStep) createordersagastate.Current {
	switch state {
	case model.CreateOrderSagaStep_ApprovalPending:
		return createordersagastate.CurrentApprovalPending
	case model.CreateOrderSagaStep_CreatingTicket:
		return createordersagastate.CurrentCreatingTicket
	case model.CreateOrderSagaStep_AuthorizingCard:
		return createordersagastate.CurrentAuthorizingCard
	case model.CreateOrderSagaStep_ApprovingTicket:
		return createordersagastate.CurrentApprovingTicket
	case model.CreateOrderSagaStep_ApprovingOrder:
		return createordersagastate.CurrentApprovingOrder
	case model.CreateOrderSagaStep_OrderApproved:
		return createordersagastate.CurrentOrderApproved
	case model.CreateOrderSagaStep_RejectingTicket:
		return createordersagastate.CurrentRejectingTicket
	case model.CreateOrderSagaStep_RejectingOrder:
		return createordersagastate.CurrentRejectingOrder
	case model.CreateOrderSagaStep_OrderRejected:
		return createordersagastate.CurrentOrderRejected

	default:
		panic("unexpected state")
	}
}

func toModelCurrentState(state createordersagastate.Current) model.CreateOrderSagaStep {
	switch state {
	case createordersagastate.CurrentApprovalPending:
		return model.CreateOrderSagaStep_ApprovalPending
	case createordersagastate.CurrentCreatingTicket:
		return model.CreateOrderSagaStep_CreatingTicket
	case createordersagastate.CurrentAuthorizingCard:
		return model.CreateOrderSagaStep_AuthorizingCard
	case createordersagastate.CurrentApprovingTicket:
		return model.CreateOrderSagaStep_ApprovingTicket
	case createordersagastate.CurrentApprovingOrder:
		return model.CreateOrderSagaStep_ApprovingOrder
	case createordersagastate.CurrentOrderApproved:
		return model.CreateOrderSagaStep_OrderApproved
	case createordersagastate.CurrentRejectingTicket:
		return model.CreateOrderSagaStep_RejectingTicket
	case createordersagastate.CurrentRejectingOrder:
		return model.CreateOrderSagaStep_RejectingOrder
	case createordersagastate.CurrentOrderRejected:
		return model.CreateOrderSagaStep_OrderRejected

	default:
		panic("unexpected state")
	}
}

func toModelTicketID(ticketID *uuid.UUID) uuid.NullUUID {
	if ticketID == nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{
		UUID:  *ticketID,
		Valid: true,
	}
}

func toModelOrderSagaState(state *ent.CreateOrderSagaState) *model.CreateOrderSagaState {
	return &model.CreateOrderSagaState{
		OrderID:  state.ID,
		Current:  toModelCurrentState(state.Current),
		TicketID: toModelTicketID(state.TicketID),
	}
}
