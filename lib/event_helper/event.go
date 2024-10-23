package event_helper

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

var EnvelopNameMap = map[protoreflect.Name]message.Type{
	"EventOrderCreated":               message.Type_TYPE_EVENT_ORDER_CREATED,
	"EventOrderApproved":              message.Type_TYPE_EVENT_ORDER_APPROVED,
	"EventOrderRejected":              message.Type_TYPE_EVENT_ORDER_REJECTED,
	"EventOrderCanceled":              message.Type_TYPE_EVENT_ORDER_CANCELED,
	"EventOrderCancellationConfirmed": message.Type_TYPE_EVENT_ORDER_CANCELLATION_CONFIRMED,
	"EventOrderCancellationRejected":  message.Type_TYPE_EVENT_ORDER_CANCELLATION_REJECTED,
	"EventTicketCreated":              message.Type_TYPE_EVENT_TICKET_CREATED,
	"EventTicketApproved":             message.Type_TYPE_EVENT_TICKET_APPROVED,
	"EventTicketRejected":             message.Type_TYPE_EVENT_TICKET_REJECTED,
	"EventTicketCreationFailed":       message.Type_TYPE_EVENT_TICKET_CREATION_FAILED,
	"EventTicketCanceled":             message.Type_TYPE_EVENT_TICKET_CANCELED,
	"EventTicketCancellationRejected": message.Type_TYPE_EVENT_TICKET_CANCELLATION_REJECTED,
	"EventCardAuthorized":             message.Type_TYPE_EVENT_CARD_AUTHORIZED,
	"EventCardAuthorizationFailed":    message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED,
	"EventCardCanceled":               message.Type_TYPE_EVENT_CARD_CANCELED,
	"CommandTicketCreate":             message.Type_TYPE_COMMAND_TICKET_CREATE,
	"CommandTicketApprove":            message.Type_TYPE_COMMAND_TICKET_APPROVE,
	"CommandTicketReject":             message.Type_TYPE_COMMAND_TICKET_REJECT,
	"CommandCardAuthorize":            message.Type_TYPE_COMMAND_CARD_AUTHORIZE,
	"CommandCardCancel":               message.Type_TYPE_COMMAND_CARD_CANCEL,
}

func ParseID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func CreateMessage(s message.Service, envelop pb.Message) (*message.Message, error) {
	name := envelop.ProtoReflect().Type().Descriptor().Name()
	messageType, ok := EnvelopNameMap[name]
	if !ok {
		return nil, fmt.Errorf("message type not found: %s", name)
	}

	v, err := anypb.New(envelop)
	if err != nil {
		return nil, err
	}

	return &message.Message{
		Subject: &message.Subject{
			Type:   messageType,
			Source: s,
		},
		Envelope: v,
	}, nil
}

func ParseMessageFromSQS(msg *types.Message) (*message.Message, error) {
	type Body struct {
		Message string `json:"message"`
	}
	var body Body
	if err := json.Unmarshal([]byte(*msg.Body), &body); err != nil {
		return nil, err
	}
	var m message.Message
	if err := protojson.Unmarshal([]byte(body.Message), &m); err != nil {
		return nil, err
	}

	return &m, nil
}
