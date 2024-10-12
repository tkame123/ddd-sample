package domain_event

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
)

type SagaFactory = func(ctx context.Context, rep repository.Repository, id model.OrderID) (*servive.CreateOrderSaga, error)

type CreateOrderSagaEventHandler interface {
	Handler(ctx context.Context, sagaFactory SagaFactory, mes *message.Message) error
}
