package repository

type Repository interface {
	Order
	CreateOrderSagaState
	CancelOrderSagaState
	ProcessedMessage
}
