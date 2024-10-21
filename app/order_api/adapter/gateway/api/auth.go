package connect

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

const authTokenHeader = "authorization"

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
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
			valid8r, err := getJWTValidator()
			if err != nil {
				return nil, errors.New("cannot get JWT validator")
			}

			_, err = valid8r.ValidateToken(ctx, accessToken)
			if err != nil {
				return nil, fmt.Errorf("cannot validate token (unauthenticated): %w", err)
			}

			// TODO: ID tokenの習得及びUser情報のContextへの格納

			// TODO: 認可の用意　アクセストークンからのScopeの対応

			// TODO: 認可の実行

			return next(ctx, req)
		}
	}
	return interceptor
}

func getJWTValidator() (*validator.Validator, error) {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
	)
	if err != nil {
		return nil, err
	}

	return jwtValidator, nil
}
