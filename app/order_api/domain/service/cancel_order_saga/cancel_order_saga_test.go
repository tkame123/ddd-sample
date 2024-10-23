package cancel_order_saga_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/mock"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
	"github.com/tkame123/ddd-sample/proto/message"
	pb "google.golang.org/protobuf/proto"
	"testing"
)

func eventCreateHelper(envelop pb.Message) *message.Message {
	ev, _ := model.CreateMessage(envelop)
	return ev
}

func TestCancelOrderSaga_ShouldCancelOrder(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockKitchenAPI := mock.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mock.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mock.NewMockCancelOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	ticketID := uuid.New()
	initialStep := &cancel_order_saga.CancelOrderSagaState{OrderID: orderID, Current: cancel_order_saga.CancelOrderSagaStep_CancelPending}
	saga, _ := cancel_order_saga.NewCancelOrderSaga(
		initialStep,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockKitchenAPI.EXPECT().CancelTicket(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(nil)
	mockBillingAPI.EXPECT().CancelCard(gomock.Any(), gomock.Any()).MinTimes(1).Return(nil)
	mockOrderSVC.EXPECT().CancelConfirmOrder(gomock.Any(), gomock.Any()).MinTimes(1)

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCanceled{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketCanceled{
				OrderId:  orderID.String(),
				TicketId: ticketID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventCardCanceled{
				OrderId: orderID.String(),
				// TODO: BillIDを追加する
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCancellationConfirmed{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != cancel_order_saga.CancelOrderSagaStep_OrderCanceled {
		t.Errorf("CurrentStep() = %v, want %v", got, cancel_order_saga.CancelOrderSagaStep_OrderCanceled)
	}
}

func TestCancelOrderSaga_OrderCancelRejectDutToTicketCancelReject(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockKitchenAPI := mock.NewMockKitchenAPI(mockCtrl)
	mockBillingAPI := mock.NewMockBillingAPI(mockCtrl)
	mockOrderSVC := mock.NewMockCancelOrder(mockCtrl)

	o, _, err := model.NewOrder(nil)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	orderID := o.OrderID
	ticketID := uuid.New()
	initialStep := &cancel_order_saga.CancelOrderSagaState{OrderID: orderID, Current: cancel_order_saga.CancelOrderSagaStep_CancelPending}
	saga, _ := cancel_order_saga.NewCancelOrderSaga(
		initialStep,
		mockOrderSVC,
		mockKitchenAPI,
		mockBillingAPI,
	)

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(saga.GetFSMVisualize())

	mockKitchenAPI.EXPECT().CancelTicket(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(nil)
	mockOrderSVC.EXPECT().CancelRejectOrder(gomock.Any(), gomock.Any()).MinTimes(1)

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCanceled{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventTicketCancellationRejected{
				OrderId:  orderID.String(),
				TicketId: ticketID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = saga.Event(ctx,
		eventCreateHelper(
			&message.EventOrderCancellationRejected{
				OrderId: orderID.String(),
			},
		))
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := saga.CurrentStep(); got != cancel_order_saga.CancelOrderSagaStep_OrderCancellationRejected {
		t.Errorf("CurrentStep() = %v, want %v", got, cancel_order_saga.CancelOrderSagaStep_OrderCancellationRejected)
	}
}
