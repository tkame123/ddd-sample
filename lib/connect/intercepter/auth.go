package intercepter

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/lib/metadata"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

const authTokenHeader = "authorization"

func NewAuthInterceptor(cfg *provider.AuthConfig) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			accessToken := req.Header().Get(authTokenHeader)
			if accessToken == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: invalid token"))
			}
			accessToken = strings.TrimPrefix(accessToken, "Bearer")
			accessToken = strings.TrimSpace(accessToken)

			// 認証の実行
			valid8r, err := getJWTValidator(cfg)
			if err != nil {
				return nil, fmt.Errorf("cannot get JWT validator: %w", err)
			}
			validatedToken, err := valid8r.ValidateToken(ctx, accessToken)
			if err != nil {
				return nil, fmt.Errorf("cannot validate token (unauthenticated): %w", err)
			}
			token, ok := validatedToken.(*validator.ValidatedClaims)
			if !ok {
				return nil, fmt.Errorf("cannot convert token to ValidatedClaims")
			}
			ctx = metadata.WithUserInfo(ctx, &metadata.UserInfo{
				// sample
				// 認可コード（openID scope 有り）：auth0|6715cc8e39951d1cd61461ed
				// client-credentialsFlow: e0tpe0rHkPJ40AHB3M0HuXlv995CPwwq@clients
				ID: token.RegisteredClaims.Subject,
			})

			// TODO: 認可の用意　アクセストークンからのScopeの対応

			// TODO: 認可の実行

			return next(ctx, req)
		}
	}
	return interceptor
}

func getJWTValidator(cfg *provider.AuthConfig) (*validator.Validator, error) {
	issuerURL, err := url.Parse("https://" + cfg.DomainName + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	jwtProvider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		jwtProvider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{cfg.AudienceName},
	)
	if err != nil {
		return nil, err
	}

	return jwtValidator, nil
}
