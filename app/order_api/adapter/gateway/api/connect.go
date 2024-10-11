package connect

import (
	"connectrpc.com/connect"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/proto/order_api/v1/order_apiv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"

	_ "github.com/lib/pq"
)

const address = "localhost:8080"

type server struct {
	rep repository.Repository
}

func NewServer(
	rep repository.Repository,
) *server {
	return &server{
		rep: rep,
	}
}

func (s *server) Run() {
	mux := http.NewServeMux()
	s.applyHandlers(mux)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func (s *server) applyHandlers(mux *http.ServeMux) {
	interceptors := s.mustInterceptors()
	// MEMO: Add handlers here.
	mux.Handle(order_apiv1connect.NewOrderServiceHandler(&orderServiceServer{rep: s.rep}, interceptors))
}

func (s *server) mustInterceptors() connect.Option {
	return connect.WithInterceptors(
	// MEMO: Add must interceptors here.
	)
}
