package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/createordersagastate"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
)

func (r *repo) CreateOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*create_order_saga.CreateOrderSagaState, error) {
	state, err := r.db.CreateOrderSagaState.Query().
		Where(createordersagastate.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return toModelOrderSagaState(state), nil
}

func (r *repo) CreateOrderSagaStateSave(ctx context.Context, state *create_order_saga.CreateOrderSagaState) error {
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

func fromModelCurrentState(state create_order_saga.CreateOrderSagaStep) createordersagastate.Current {
	switch state {
	case create_order_saga.CreateOrderSagaStep_ApprovalPending:
		return createordersagastate.CurrentApprovalPending
	case create_order_saga.CreateOrderSagaStep_CreatingTicket:
		return createordersagastate.CurrentCreatingTicket
	case create_order_saga.CreateOrderSagaStep_AuthorizingCard:
		return createordersagastate.CurrentAuthorizingCard
	case create_order_saga.CreateOrderSagaStep_ApprovingTicket:
		return createordersagastate.CurrentApprovingTicket
	case create_order_saga.CreateOrderSagaStep_ApprovingOrder:
		return createordersagastate.CurrentApprovingOrder
	case create_order_saga.CreateOrderSagaStep_OrderApproved:
		return createordersagastate.CurrentOrderApproved
	case create_order_saga.CreateOrderSagaStep_RejectingTicket:
		return createordersagastate.CurrentRejectingTicket
	case create_order_saga.CreateOrderSagaStep_RejectingOrder:
		return createordersagastate.CurrentRejectingOrder
	case create_order_saga.CreateOrderSagaStep_OrderRejected:
		return createordersagastate.CurrentOrderRejected

	default:
		panic("unexpected state")
	}
}

func toModelCurrentState(state createordersagastate.Current) create_order_saga.CreateOrderSagaStep {
	switch state {
	case createordersagastate.CurrentApprovalPending:
		return create_order_saga.CreateOrderSagaStep_ApprovalPending
	case createordersagastate.CurrentCreatingTicket:
		return create_order_saga.CreateOrderSagaStep_CreatingTicket
	case createordersagastate.CurrentAuthorizingCard:
		return create_order_saga.CreateOrderSagaStep_AuthorizingCard
	case createordersagastate.CurrentApprovingTicket:
		return create_order_saga.CreateOrderSagaStep_ApprovingTicket
	case createordersagastate.CurrentApprovingOrder:
		return create_order_saga.CreateOrderSagaStep_ApprovingOrder
	case createordersagastate.CurrentOrderApproved:
		return create_order_saga.CreateOrderSagaStep_OrderApproved
	case createordersagastate.CurrentRejectingTicket:
		return create_order_saga.CreateOrderSagaStep_RejectingTicket
	case createordersagastate.CurrentRejectingOrder:
		return create_order_saga.CreateOrderSagaStep_RejectingOrder
	case createordersagastate.CurrentOrderRejected:
		return create_order_saga.CreateOrderSagaStep_OrderRejected

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

func toModelOrderSagaState(state *ent.CreateOrderSagaState) *create_order_saga.CreateOrderSagaState {
	return &create_order_saga.CreateOrderSagaState{
		OrderID:  state.ID,
		Current:  toModelCurrentState(state.Current),
		TicketID: toModelTicketID(state.TicketID),
	}
}
