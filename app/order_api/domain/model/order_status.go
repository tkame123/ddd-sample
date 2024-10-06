package model

type OrderStatus int

const (
	ApprovalPending OrderStatus = iota
	CreatingTicket
	AuthorizingCard
	ApprovingTicket
	ApprovingOrder
	OrderApproved

	RejectingTicket
	RejectingOrder
	OrderRejected
)

func (s OrderStatus) String() string {
	switch s {
	case ApprovalPending:
		return "ApprovalPending"
	case CreatingTicket:
		return "CreatingTicket"
	case AuthorizingCard:
		return "AuthorizingCard"
	case ApprovingTicket:
		return "ApprovingTicket"
	case ApprovingOrder:
		return "ApprovingOrder"
	case OrderApproved:
		return "OrderApproved"
	case RejectingTicket:
		return "RejectingTicket"
	case RejectingOrder:
		return "RejectingOrder"
	case OrderRejected:
		return "OrderRejected"
	default:
		return "Unknown"
	}
}
