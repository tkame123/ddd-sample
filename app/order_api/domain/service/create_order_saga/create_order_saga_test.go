package create_order_saga_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/mock"
	servive "github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
	pb "google.golang.org/protobuf/proto"
	"testing"
)

func eventCreateHelper(envelop pb.Message) *message.Message {
	ev, _ := model.CreateMessage(envelop)
	return ev
}

func TestCreateOrderSaga_ShouldCreateOrder(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockRepository(mockCtrl)
	mockKitchenAPI := mock.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mock.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mock.NewMockCreateOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	ticketID := uuid.New()
	initialStep := &model.CreateOrderSagaState{OrderID: orderID, Current: model.CreateOrderSagaStep_ApprovalPending}
	saga, _ := servive.NewCreateOrderSaga(
		initialStep,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockRepo.EXPECT().CreateOrderSagaStateSave(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().ApproveTicket(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockBillingAPI.EXPECT().AuthorizeCard(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockOrderSVC.EXPECT().ApproveOrder(gomock.Any(), gomock.Any()).AnyTimes()

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCreated{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketCreated{
				OrderId:  orderID.String(),
				TicketId: ticketID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	state := saga.ExportState()
	if state.TicketID.Valid != true {
		t.Errorf("TicketID is not valid")
	}
	if state.TicketID.UUID != ticketID {
		t.Errorf("TicketID is not equal")
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventCardAuthorized{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketApproved{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderApproved{
				OrderId: orderID.String(),
			},
		))
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
	mockRepo := mock.NewMockRepository(mockCtrl)
	mockKitchenAPI := mock.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mock.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mock.NewMockCreateOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	initialStep := &model.CreateOrderSagaState{OrderID: orderID, Current: model.CreateOrderSagaStep_ApprovalPending}
	saga, _ := servive.NewCreateOrderSaga(
		initialStep,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockRepo.EXPECT().CreateOrderSagaStateSave(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCreated{
				OrderId: orderID.String(),
			}))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketCreationFailed{
				OrderId: orderID.String(),
			}))
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
	mockRepo := mock.NewMockRepository(mockCtrl)
	mockKitchenAPI := mock.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mock.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mock.NewMockCreateOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	initialStep := &model.CreateOrderSagaState{OrderID: orderID, Current: model.CreateOrderSagaStep_ApprovalPending}
	saga, _ := servive.NewCreateOrderSaga(
		initialStep,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockRepo.EXPECT().CreateOrderSagaStateSave(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().CreateTicket(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockKitchenAPI.EXPECT().RejectTicket(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockBillingAPI.EXPECT().AuthorizeCard(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockOrderSVC.EXPECT().RejectOrder(gomock.Any(), gomock.Any()).AnyTimes()

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCreated{
				OrderId: orderID.String(),
			}))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketCreated{
				OrderId: orderID.String(),
			}))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventCardAuthorizationFailed{
				OrderId: orderID.String(),
			}))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketRejected{
				OrderId: orderID.String(),
			}))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderRejected{
				OrderId: orderID.String(),
			}))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != model.CreateOrderSagaStep_OrderRejected {
		t.Errorf("CurrentStep() = %v, want %v", got, model.CreateOrderSagaStep_OrderRejected)
	}
}
