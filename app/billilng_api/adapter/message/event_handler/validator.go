package event_handler

import (
	"github.com/tkame123/ddd-sample/proto/message"
)

func IsCreateBillEvent(tp message.Type) bool {
	_, ok := EventMap[tp]
	return ok
}
