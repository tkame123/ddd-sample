package connect

import (
	"connectrpc.com/connect"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/idempotency"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/lib/connect/intercepter"
	"github.com/tkame123/ddd-sample/proto/order_api/v1/order_apiv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"

	_ "github.com/lib/pq"
)

const address = "localhost:8080"

type Server struct {
	authCfg        *provider.AuthConfig
	rep            repository.Repository
	pub            domain_event.Publisher
	repIdempotency *idempotency.Repository
	enforcer       *casbin.Enforcer
}

func NewServer(
	authCfg *provider.AuthConfig,
	rep repository.Repository,
	pub domain_event.Publisher,
	repIdempotency *idempotency.Repository,
	enforcer *casbin.Enforcer,
) Server {
	return Server{
		authCfg:        authCfg,
		rep:            rep,
		pub:            pub,
		repIdempotency: repIdempotency,
		enforcer:       enforcer,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()
	s.applyHandlers(mux)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func (s *Server) applyHandlers(mux *http.ServeMux) {
	mustInterceptors := s.mustInterceptors()
	// MEMO: Add handlers here.
	mux.Handle(order_apiv1connect.NewOrderServiceHandler(&orderServiceServer{rep: s.rep, pub: s.pub}, mustInterceptors))
}

func (s *Server) mustInterceptors() connect.Option {
	return connect.WithInterceptors(
		// MEMO: Add must interceptors here.
		intercepter.NewAuthInterceptor(s.authCfg, s.enforcer),
		intercepter.NewIdempotencyCheckInterceptor(s.repIdempotency),
	)
}
