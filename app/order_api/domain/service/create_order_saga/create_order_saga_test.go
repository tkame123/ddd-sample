package create_order_saga_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	mockAPI "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/fake"
	mockRp "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/repository"
	mockOSVS "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/service"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	event_handler2 "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
	"github.com/tkame123/ddd-sample/lib/event"
	"testing"
)

func TestCreateOrderSaga_ShouldCreateOrder(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mockRp.NewMockRepository(mockCtrl)
	mockKitchenAPI := mockAPI.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mockAPI.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mockOSVS.NewMockCreateOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	initialStep := model.NewCreateOrderSagaState(orderID, model.CreateOrderSagaStep_ApprovalPending)
	saga := servive.NewCreateOrderSaga(
		initialStep,
		mockRepo,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockRepo.EXPECT().CreateOrderSagaStateSave(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes()
	mockKitchenAPI.EXPECT().ApproveTicket(gomock.Any(), gomock.Any()).AnyTimes()
	mockBillingAPI.EXPECT().AuthorizeCard(gomock.Any(), gomock.Any()).AnyTimes()
	mockOrderSVC.EXPECT().ApproveOrder(gomock.Any(), gomock.Any()).AnyTimes()

	err = event_handler2.NewNextStepSagaWhenOrderCreatedHandler(saga).
		Handler(ctx, &model.OrderCreatedEvent{OrderID: orderID})
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenTicketCreatedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_TicketCreated))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenCardAuthorizedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_CardAuthorized))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenTicketApprovedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_TicketApproved))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenOrderApprovedHandler(saga).
		Handler(ctx, &model.OrderApprovedEvent{OrderID: orderID})
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != model.CreateOrderSagaStep_OrderApproved {
		t.Errorf("CurrentStep() = %v, want %v", got, model.CreateOrderSagaStep_OrderApproved)
	}
}

func TestCreateOrderSaga_OrderRejectedDutToTicketCreationFailed(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mockRp.NewMockRepository(mockCtrl)
	mockKitchenAPI := mockAPI.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mockAPI.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mockOSVS.NewMockCreateOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	initialStep := model.NewCreateOrderSagaState(orderID, model.CreateOrderSagaStep_ApprovalPending)
	saga := servive.NewCreateOrderSaga(
		initialStep,
		mockRepo,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockRepo.EXPECT().CreateOrderSagaStateSave(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes()

	err = event_handler2.NewNextStepSagaWhenOrderCreatedHandler(saga).
		Handler(ctx, &model.OrderCreatedEvent{OrderID: orderID})
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenTicketCreationFailedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_TicketCreationFailed))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != model.CreateOrderSagaStep_OrderRejected {
		t.Errorf("CurrentStep() = %v, want %v", got, model.CreateOrderSagaStep_OrderRejected)
	}
}

func TestCreateOrderSaga_OrderRejectedDutToCardAuthorizeFailed(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mockRp.NewMockRepository(mockCtrl)
	mockKitchenAPI := mockAPI.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mockAPI.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mockOSVS.NewMockCreateOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	initialStep := model.NewCreateOrderSagaState(orderID, model.CreateOrderSagaStep_ApprovalPending)
	saga := servive.NewCreateOrderSaga(
		initialStep,
		mockRepo,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockRepo.EXPECT().CreateOrderSagaStateSave(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes()
	mockKitchenAPI.EXPECT().RejectTicket(gomock.Any(), gomock.Any()).AnyTimes()
	mockBillingAPI.EXPECT().AuthorizeCard(gomock.Any(), gomock.Any()).AnyTimes()
	mockOrderSVC.EXPECT().RejectOrder(gomock.Any(), gomock.Any()).AnyTimes()

	err = event_handler2.NewNextStepSagaWhenOrderCreatedHandler(saga).
		Handler(ctx, &model.OrderCreatedEvent{OrderID: orderID})
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenTicketCreatedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_TicketCreated))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenCardAuthorizeFailedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_CardAuthorizeFailed))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenTicketRejectedHandler(saga).
		Handler(ctx, fake.NewGeneralEvent(orderID, event.EventName_TicketRejected))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler2.NewNextStepSagaWhenOrderRejectedHandler(saga).
		Handler(ctx, &model.OrderRejectedEvent{OrderID: orderID})
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != model.CreateOrderSagaStep_OrderRejected {
		t.Errorf("CurrentStep() = %v, want %v", got, model.CreateOrderSagaStep_OrderRejected)
	}
}