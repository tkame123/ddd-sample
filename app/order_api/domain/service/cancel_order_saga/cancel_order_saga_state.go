package cancel_order_saga

import "github.com/tkame123/ddd-sample/app/order_api/domain/model"

type CancelOrderSagaState struct {
	Current CancelOrderSagaStep

	OrderID  model.OrderID
	TicketID model.TicketID
}

type CancelOrderSagaStep = string

const (
	CancelOrderSagaStep_CancelPending               CancelOrderSagaStep = "CancelPending"
	CancelOrderSagaStep_CancelingTicket             CancelOrderSagaStep = "CancelingTicket"
	CancelOrderSagaStep_CancelingCard               CancelOrderSagaStep = "CancelingCard"
	CancelOrderSagaStep_CancellationConfirmingOrder CancelOrderSagaStep = "CancellationConfirmingOrder"
	CancelOrderSagaStep_OrderCanceled               CancelOrderSagaStep = "OrderCanceled"

	CancelOrderSagaStep_CancellationRejectingOrder CancelOrderSagaStep = "CancellationRejectingOrder"
	CancelOrderSagaStep_OrderCancellationRejected  CancelOrderSagaStep = "OrderCancellationRejected"
)
