package event_helper

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func ParseID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

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
