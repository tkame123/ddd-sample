package intercepter

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/lib/auth"
	"github.com/tkame123/ddd-sample/lib/connect/intercepter/auth_intercepter"
	"github.com/tkame123/ddd-sample/lib/metadata"
	"strings"
)

const authTokenHeader = "authorization"
const developModeHeader = "x-develop-mode"

func NewAuthInterceptor(env *provider.EnvConfig, cfg *provider.AuthConfig, enforcer *casbin.Enforcer) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			authProvider := auth.Context{}
			var strategy auth.Strategy
			if env.ENV != provider.EnvProduction &&
				req.Header().Get(developModeHeader) == "true" {
				strategy = auth_intercepter.NewAuthStrategyNop()
			} else {
				accessToken := req.Header().Get(authTokenHeader)
				if accessToken == "" {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: invalid token"))
				}
				accessToken = strings.TrimPrefix(accessToken, "Bearer")
				accessToken = strings.TrimSpace(accessToken)

				strategy = auth_intercepter.NewAuthStrategyOIDC(cfg, enforcer, accessToken, req.Spec().Procedure)
			}
			authProvider.SetAuthStrategy(strategy)

			userInfo, err := authProvider.Authenticate(ctx)
			if auth_intercepter.IsAuthenticationError(err) {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			} else if err != nil {
				return nil, err
			}

			ctx = metadata.WithUserInfo(ctx, userInfo)

			err = authProvider.Authorize(ctx)
			if auth_intercepter.IsPermissionError(err) {
				return nil, connect.NewError(connect.CodePermissionDenied, err)
			} else if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
	return interceptor
}
