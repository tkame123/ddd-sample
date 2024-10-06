package repository

type Repository struct {
	Order
	CreateOrderSagaState
}

func NewRepository(order Order, cos CreateOrderSagaState) *Repository {
	return &Repository{Order: order, CreateOrderSagaState: cos}
}
