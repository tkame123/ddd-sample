package repository

type Repository interface {
	Order
	CreateOrderSagaState
	ProcessedMessage
}
