package repository

type Repository struct {
	Ticket
}

func NewRepository(ticket Ticket) *Repository {
	return &Repository{Ticket: ticket}
}
