package repository

type Repository struct {
	Shipment
}

func NewRepository(shipment Shipment) *Repository {
	return &Repository{Shipment: shipment}
}
