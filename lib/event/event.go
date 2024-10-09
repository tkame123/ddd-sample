package event

type Name = string

const (
	// orderAPI
	EventName_OrderCreated  Name = "event.order.order_created"
	EventName_OrderApproved Name = "event.order.order_approved"
	EventName_OrderRejected Name = "event.order.order_rejected"

	// kitchenAPI
	EventName_TicketCreated        Name = "event.kitchen.ticket_created"
	EventName_TicketCreationFailed Name = "event.kitchen.ticket_creation_failed"
	EventName_TicketApproved       Name = "event.kitchen.ticket_approved"
	EventName_TicketRejected       Name = "event.kitchen.ticket_rejected"

	// BillingAPI
	EventName_CardAuthorized      Name = "event.billing.card_authorized"
	EventName_CardAuthorizeFailed Name = "event.billing.card_authorize_failed"
)

type Event interface {
	Name() Name
}

type GeneralEvent struct {
	id   int
	name string
}

func NewGeneralEvent(id int, name string) *GeneralEvent {
	return &GeneralEvent{
		id:   id,
		name: name,
	}
}

func (e *GeneralEvent) Name() Name {
	return e.name
}

func (e *GeneralEvent) ID() int {
	return e.id
}
