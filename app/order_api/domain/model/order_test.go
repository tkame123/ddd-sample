package model_test

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"testing"
)

func TestOrder_ShouldApproved(t *testing.T) {
	ctx := context.Background()

	order := &model.Order{
		Status: model.OrderStatus_ApprovalPending,
	}

	sm := order.StatusFSM()

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(fsm.VisualizeWithType(sm, fsm.GRAPHVIZ))

	if got := sm.Current(); got != model.OrderStatus_ApprovalPending {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_ApprovalPending)
	}

	err := sm.Event(ctx, "authorized")
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := sm.Current(); got != model.OrderStatus_Approved {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_Approved)
	}
}

func TestOrder_ShouldReject(t *testing.T) {
	ctx := context.Background()

	order := &model.Order{
		Status: model.OrderStatus_ApprovalPending,
	}

	sm := order.StatusFSM()

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(fsm.VisualizeWithType(sm, fsm.GRAPHVIZ))

	if got := sm.Current(); got != model.OrderStatus_ApprovalPending {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_ApprovalPending)
	}

	err := sm.Event(ctx, "rejected")
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := sm.Current(); got != model.OrderStatus_Rejected {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_Rejected)
	}
}

func TestOrder_ShouldCanceled(t *testing.T) {
	ctx := context.Background()

	order := &model.Order{
		Status: model.OrderStatus_Approved,
	}

	sm := order.StatusFSM()

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(fsm.VisualizeWithType(sm, fsm.GRAPHVIZ))

	if got := sm.Current(); got != model.OrderStatus_Approved {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_Approved)
	}

	err := sm.Event(ctx, "cancel")
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = sm.Event(ctx, "cancelConfirmed")
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := sm.Current(); got != model.OrderStatus_Canceled {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_Canceled)
	}
}

func TestOrder_CancelRejected(t *testing.T) {
	ctx := context.Background()

	order := &model.Order{
		Status: model.OrderStatus_Approved,
	}

	sm := order.StatusFSM()

	// 下記で絵になるので、迷ったら出力して比較する
	// http://www.webgraphviz.com/
	fmt.Println(fsm.VisualizeWithType(sm, fsm.GRAPHVIZ))

	if got := sm.Current(); got != model.OrderStatus_Approved {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_Approved)
	}

	err := sm.Event(ctx, "cancel")
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	err = sm.Event(ctx, "cancelRejected")
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	if got := sm.Current(); got != model.OrderStatus_Approved {
		t.Errorf("CurrentStep() = %v, want %v", got, model.OrderStatus_Approved)
	}
}
