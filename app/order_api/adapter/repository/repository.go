package repository

type Repository struct {
	Order
}

func NewRepository(order Order) *Repository {
	return &Repository{Order: order}
}
