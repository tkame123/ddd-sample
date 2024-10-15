package event_handler

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/model"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type CardAuthorizeWhenCardAuthorizeHandler struct {
	svc service.CreateBill
}

func NewCardAuthorizeWhenCardAuthorizeHandler(svc service.CreateBill) domain_event.EventHandler {
	return &CardAuthorizeWhenCardAuthorizeHandler{svc: svc}
}

func (h *CardAuthorizeWhenCardAuthorizeHandler) Handler(ctx context.Context, mes *message.Message) error {
	if mes.Subject.Type != message.Type_TYPE_COMMAND_CARD_AUTHORIZE {
		return fmt.Errorf("invalid event type: %v", mes.Subject.Type)
	}

	var v message.CommandCardAuthorize
	err := mes.Envelope.UnmarshalTo(&v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	id, err := model.OrderIdParse(v.OrderId)
	if err != nil {
		return fmt.Errorf("failed to parse order id: %w", err)
	}

	if err := h.svc.AuthorizeCard(ctx, *id, nil); err != nil {
		return err
	}

	return nil
}
