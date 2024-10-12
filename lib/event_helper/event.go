package event_helper

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/proto/message"
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
