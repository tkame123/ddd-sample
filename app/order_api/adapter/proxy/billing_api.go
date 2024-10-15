package proxy

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
)

type BillingAPI struct {
	pub domain_event.Publisher
}

func NewBillingAPI(pub domain_event.Publisher) external_service.BillingAPI {
	return &BillingAPI{pub: pub}
}

func (k *BillingAPI) AuthorizeCard(ctx context.Context, orderID model.OrderID) error {
	//	TODO: Implement this(仮置き）
	log.Println("AuthorizeCard")
	command, err := model.CreateMessage(
		&message.CommandCardAuthorize{
			OrderId: orderID.String(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	k.pub.PublishMessages(ctx, []*message.Message{command})

	return nil
}
