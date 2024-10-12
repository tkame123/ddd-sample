package domain_event

import (
	"context"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type CreateOrderSagaEventHandler interface {
	Handler(ctx context.Context, saga *servive.CreateOrderSaga, mes *message.Message) error
}
