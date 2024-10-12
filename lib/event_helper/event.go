package event_helper

import (
	"github.com/tkame123/ddd-sample/proto/message"
	pb "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

//// Deprecated: use ProtoEvent
//type Name = string
//
//const (
//	// orderAPI
//	EventName_OrderCreated  Name = "event-order-order_created"
//	EventName_OrderApproved Name = "event-order-order_approved"
//	EventName_OrderRejected Name = "event-order-order_rejected"
//
//	// kitchenAPI
//	EventName_TicketCreated        Name = "event-kitchen-ticket_created"
//	EventName_TicketCreationFailed Name = "event-kitchen-ticket_creation_failed"
//	EventName_TicketApproved       Name = "event-kitchen-ticket_approved"
//	EventName_TicketRejected       Name = "event-kitchen-ticket_rejected"
//
//	CommandName_TicketCreate  Name = "command-kitchen-ticket_create"
//	CommandName_TicketApprove Name = "command-kitchen-ticket_approve"
//	CommandName_TicketReject  Name = "command-kitchen-ticket_reject"
//
//	// BillingAPI
//	EventName_CardAuthorized      Name = "event-billing-card_authorized"
//	EventName_CardAuthorizeFailed Name = "event-billing-card_authorize_failed"
//
//	CommandName_CardAuthorize Name = "command-billing-card_authorize"
//)
//
//// Deprecated: use ProtoEvent
//type Event interface {
//	ID() uuid.UUID
//	Name() Name
//	ToBody() (string, error)
//}
//
//// message送受信用の構造体
//// Deprecated: use ProtoEvent
//type RawEvent struct {
//	Type   string          `json:"type"`
//	ID     string          `json:"id"`
//	Origin json.RawMessage `json:"origin"`
//}

func CreateMessage(t message.Type, s message.Service, envelop pb.Message) (*message.Message, error) {
	v, err := anypb.New(envelop)
	if err != nil {
		return nil, err
	}

	return &message.Message{
		Subject: &message.Subject{
			Type:   t,
			Source: s,
		},
		Envelope: v,
	}, nil
}
