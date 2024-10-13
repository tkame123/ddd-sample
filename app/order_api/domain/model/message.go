package model

import (
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
	pb "google.golang.org/protobuf/proto"
)

const myService message.Service = message.Service_SERVICE_ORDER

func CreateMessage(m pb.Message) (*message.Message, error) {
	return event_helper.CreateMessage(myService, m)
}
