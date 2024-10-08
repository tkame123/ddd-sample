package servive_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	mockAPI "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/external_service"
	mockRp "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/repository"
	mockOSVS "github.com/tkame123/ddd-sample/app/order_api/domain/port/mock/service"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/event_handler"
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
	orderID := o.OrderID()
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
	orderID := o.OrderID()
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

	err = event_handler.NewNextStepSagaWhenOrderCreatedHandler(saga).
		Handler(ctx, *model.NewOrderCreatedEvent(orderID))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenTicketCreationFailedHandler(saga).
		Handler(ctx, event.NewGeneralEvent(orderID, event.EventName_TicketCreationFailed))
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
	orderID := o.OrderID()
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

	err = event_handler.NewNextStepSagaWhenCardAuthorizeFailedHandler(saga).
		Handler(ctx, event.NewGeneralEvent(orderID, event.EventName_CardAuthorizeFailed))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenTicketRejectedHandler(saga).
		Handler(ctx, event.NewGeneralEvent(orderID, event.EventName_TicketRejected))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = event_handler.NewNextStepSagaWhenOrderRejectedHandler(saga).
		Handler(ctx, *model.NewOrderRejectedEvent(orderID))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != model.CreateOrderSagaStep_OrderRejected {
		t.Errorf("CurrentStep() = %v, want %v", got, model.CreateOrderSagaStep_OrderRejected)
	}
}
