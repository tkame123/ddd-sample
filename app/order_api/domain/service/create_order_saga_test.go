package servive_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	mockMes "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/domain_event"
	mockAPI "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/external_service"
	mockRp "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/event_handler"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/order"
	"github.com/tkame123/ddd-sample/lib/event"
	"testing"
)

func TestCreateOrderSaga_ShouldCreateOrder(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrder := mockRp.NewMockOrder(mockCtrl)
	mockCOS := mockRp.NewMockCreateOrderSagaState(mockCtrl)
	mockPub := mockMes.NewMockPublisher(mockCtrl)
	mockKitchenAPI := mockAPI.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mockAPI.NewMockBillingAPI(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderSvc := order.NewService(repository.Repository{Order: mockOrder, CreateOrderSagaState: mockCOS}, mockPub)
	orderID := o.OrderID()
	initialStep := model.NewCreateOrderSagaState(orderID, model.CreateOrderSagaStep_ApprovalPending)
	saga := servive.NewCreateOrderSaga(
		initialStep,
		&repository.Repository{Order: mockOrder, CreateOrderSagaState: mockCOS},
		orderSvc,
		&external_service.ExternalAPI{KitchenAPI: mockKitchenAPI, BillingAPI: mockBillingAPI},
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockOrder.EXPECT().FindOne(gomock.Any(), gomock.Any()).AnyTimes().Return(o, nil)
	mockCOS.EXPECT().Save(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes()
	mockKitchenAPI.EXPECT().ApproveTicket(gomock.Any(), gomock.Any()).AnyTimes()
	mockBillingAPI.EXPECT().AuthorizeCard(gomock.Any(), gomock.Any()).AnyTimes()
	mockPub.EXPECT().PublishMessages(gomock.Any(), gomock.Any()).AnyTimes()

	err = event_handler.NewNextStepSagaWhenOrderCreatedHandler(saga).
		Handler(ctx, *model.NewOrderCreatedEvent(orderID))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenTicketCreatedHandler(saga).
		Handler(ctx, event.NewGeneralEvent(orderID, event.EventName_TicketCreated))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenCardAuthorizedHandler(saga).
		Handler(ctx, event.NewGeneralEvent(orderID, event.EventName_CardAuthorized))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenTicketApprovedHandler(saga).
		Handler(ctx, event.NewGeneralEvent(orderID, event.EventName_TicketApproved))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenOrderApprovedHandler(saga).
		Handler(ctx, *model.NewOrderApprovedEvent(orderID))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != model.CreateOrderSagaStep_OrderApproved {
		t.Errorf("CurrentStep() = %v, want %v", got, model.CreateOrderSagaStep_OrderApproved)
	}
}
