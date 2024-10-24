// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: order_api/v1/order.proto

package order_apiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/tkame123/ddd-sample/proto/order_api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// OrderServiceName is the fully-qualified name of the OrderService service.
	OrderServiceName = "order_api.v1.OrderService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OrderServiceFindOrderProcedure is the fully-qualified name of the OrderService's FindOrder RPC.
	OrderServiceFindOrderProcedure = "/order_api.v1.OrderService/FindOrder"
	// OrderServiceCreateOrderProcedure is the fully-qualified name of the OrderService's CreateOrder
	// RPC.
	OrderServiceCreateOrderProcedure = "/order_api.v1.OrderService/CreateOrder"
	// OrderServiceCancelOrderProcedure is the fully-qualified name of the OrderService's CancelOrder
	// RPC.
	OrderServiceCancelOrderProcedure = "/order_api.v1.OrderService/CancelOrder"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	orderServiceServiceDescriptor           = v1.File_order_api_v1_order_proto.Services().ByName("OrderService")
	orderServiceFindOrderMethodDescriptor   = orderServiceServiceDescriptor.Methods().ByName("FindOrder")
	orderServiceCreateOrderMethodDescriptor = orderServiceServiceDescriptor.Methods().ByName("CreateOrder")
	orderServiceCancelOrderMethodDescriptor = orderServiceServiceDescriptor.Methods().ByName("CancelOrder")
)

// OrderServiceClient is a client for the order_api.v1.OrderService service.
type OrderServiceClient interface {
	FindOrder(context.Context, *connect.Request[v1.FindOrderRequest]) (*connect.Response[v1.FindOrderResponse], error)
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error)
	CancelOrder(context.Context, *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error)
}

// NewOrderServiceClient constructs a client for the order_api.v1.OrderService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrderServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OrderServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &orderServiceClient{
		findOrder: connect.NewClient[v1.FindOrderRequest, v1.FindOrderResponse](
			httpClient,
			baseURL+OrderServiceFindOrderProcedure,
			connect.WithSchema(orderServiceFindOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createOrder: connect.NewClient[v1.CreateOrderRequest, v1.CreateOrderResponse](
			httpClient,
			baseURL+OrderServiceCreateOrderProcedure,
			connect.WithSchema(orderServiceCreateOrderMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
		cancelOrder: connect.NewClient[v1.CancelOrderRequest, v1.CancelOrderResponse](
			httpClient,
			baseURL+OrderServiceCancelOrderProcedure,
			connect.WithSchema(orderServiceCancelOrderMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyIdempotent),
			connect.WithClientOptions(opts...),
		),
	}
}

// orderServiceClient implements OrderServiceClient.
type orderServiceClient struct {
	findOrder   *connect.Client[v1.FindOrderRequest, v1.FindOrderResponse]
	createOrder *connect.Client[v1.CreateOrderRequest, v1.CreateOrderResponse]
	cancelOrder *connect.Client[v1.CancelOrderRequest, v1.CancelOrderResponse]
}

// FindOrder calls order_api.v1.OrderService.FindOrder.
func (c *orderServiceClient) FindOrder(ctx context.Context, req *connect.Request[v1.FindOrderRequest]) (*connect.Response[v1.FindOrderResponse], error) {
	return c.findOrder.CallUnary(ctx, req)
}

// CreateOrder calls order_api.v1.OrderService.CreateOrder.
func (c *orderServiceClient) CreateOrder(ctx context.Context, req *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error) {
	return c.createOrder.CallUnary(ctx, req)
}

// CancelOrder calls order_api.v1.OrderService.CancelOrder.
func (c *orderServiceClient) CancelOrder(ctx context.Context, req *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error) {
	return c.cancelOrder.CallUnary(ctx, req)
}

// OrderServiceHandler is an implementation of the order_api.v1.OrderService service.
type OrderServiceHandler interface {
	FindOrder(context.Context, *connect.Request[v1.FindOrderRequest]) (*connect.Response[v1.FindOrderResponse], error)
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error)
	CancelOrder(context.Context, *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error)
}

// NewOrderServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrderServiceHandler(svc OrderServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	orderServiceFindOrderHandler := connect.NewUnaryHandler(
		OrderServiceFindOrderProcedure,
		svc.FindOrder,
		connect.WithSchema(orderServiceFindOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceCreateOrderHandler := connect.NewUnaryHandler(
		OrderServiceCreateOrderProcedure,
		svc.CreateOrder,
		connect.WithSchema(orderServiceCreateOrderMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceCancelOrderHandler := connect.NewUnaryHandler(
		OrderServiceCancelOrderProcedure,
		svc.CancelOrder,
		connect.WithSchema(orderServiceCancelOrderMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyIdempotent),
		connect.WithHandlerOptions(opts...),
	)
	return "/order_api.v1.OrderService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OrderServiceFindOrderProcedure:
			orderServiceFindOrderHandler.ServeHTTP(w, r)
		case OrderServiceCreateOrderProcedure:
			orderServiceCreateOrderHandler.ServeHTTP(w, r)
		case OrderServiceCancelOrderProcedure:
			orderServiceCancelOrderHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOrderServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrderServiceHandler struct{}

func (UnimplementedOrderServiceHandler) FindOrder(context.Context, *connect.Request[v1.FindOrderRequest]) (*connect.Response[v1.FindOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order_api.v1.OrderService.FindOrder is not implemented"))
}

func (UnimplementedOrderServiceHandler) CreateOrder(context.Context, *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order_api.v1.OrderService.CreateOrder is not implemented"))
}

func (UnimplementedOrderServiceHandler) CancelOrder(context.Context, *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order_api.v1.OrderService.CancelOrder is not implemented"))
}
