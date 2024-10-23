package database

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/cancelordersagastate"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
)

func (r *repo) CancelOrderSagaStateFindOne(ctx context.Context, id model.OrderID) (*cancel_order_saga.CancelOrderSagaState, error) {
	state, err := r.db.CancelOrderSagaState.Query().
		Where(cancelordersagastate.ID(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return toModelCancelOrderSagaState(state), nil
}

func (r *repo) CancelOrderSagaStateSave(ctx context.Context, state *cancel_order_saga.CancelOrderSagaState) error {
	err := r.db.CancelOrderSagaState.Create().
		SetID(state.OrderID).
		SetCurrent(fromModelCancelOrderSagaCurrentState(state.Current)).
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

func fromModelCancelOrderSagaCurrentState(state cancel_order_saga.CancelOrderSagaStep) cancelordersagastate.Current {
	switch state {
	case cancel_order_saga.CancelOrderSagaStep_CancelPending:
		return cancelordersagastate.CurrentCancelPending
	case cancel_order_saga.CancelOrderSagaStep_CancelingTicket:
		return cancelordersagastate.CurrentCancelingTicket
	case cancel_order_saga.CancelOrderSagaStep_CancelingCard:
		return cancelordersagastate.CurrentCancelingCard
	case cancel_order_saga.CancelOrderSagaStep_CancellationConfirmingOrder:
		return cancelordersagastate.CurrentCancellationConfirmingOrder
	case cancel_order_saga.CancelOrderSagaStep_OrderCanceled:
		return cancelordersagastate.CurrentOrderCanceled
	case cancel_order_saga.CancelOrderSagaStep_CancellationRejectingOrder:
		return cancelordersagastate.CurrentCancellationRejectingOrder
	case cancel_order_saga.CancelOrderSagaStep_OrderCancellationRejected:
		return cancelordersagastate.CurrentOrderCancellationRejected

	default:
		panic("invalid state")
	}

}

func toModelCancelOrderSagaCurrentState(state cancelordersagastate.Current) cancel_order_saga.CancelOrderSagaStep {
	switch state {
	case cancelordersagastate.CurrentCancelPending:
		return cancel_order_saga.CancelOrderSagaStep_CancelPending
	case cancelordersagastate.CurrentCancelingTicket:
		return cancel_order_saga.CancelOrderSagaStep_CancelingTicket
	case cancelordersagastate.CurrentCancelingCard:
		return cancel_order_saga.CancelOrderSagaStep_CancelingCard
	case cancelordersagastate.CurrentCancellationConfirmingOrder:
		return cancel_order_saga.CancelOrderSagaStep_CancellationConfirmingOrder
	case cancelordersagastate.CurrentOrderCanceled:
		return cancel_order_saga.CancelOrderSagaStep_OrderCanceled
	case cancelordersagastate.CurrentCancellationRejectingOrder:
		return cancel_order_saga.CancelOrderSagaStep_CancellationRejectingOrder
	case cancelordersagastate.CurrentOrderCancellationRejected:
		return cancel_order_saga.CancelOrderSagaStep_OrderCancellationRejected

	default:
		panic("invalid state")
	}
}

func toModelCancelOrderSagaState(state *ent.CancelOrderSagaState) *cancel_order_saga.CancelOrderSagaState {
	return &cancel_order_saga.CancelOrderSagaState{
		OrderID:  state.ID,
		Current:  toModelCancelOrderSagaCurrentState(state.Current),
		TicketID: toModelTicketID(state.TicketID),
	}
}
