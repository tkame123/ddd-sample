package auth_intercepter

import (
	"context"
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/lib/auth"
	"github.com/tkame123/ddd-sample/lib/metadata"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type Auth struct {
	cfg         *provider.AuthConfig
	accessToken string
}

func NewAuth(cfg *provider.AuthConfig, accessToken string) (auth.Strategy, error) {
	return &Auth{
		cfg:         cfg,
		accessToken: accessToken,
	}, nil
}

func (a *Auth) Authenticate(ctx context.Context) error {
	valid8r, err := getJWTValidator(a.cfg)
	if err != nil {
		return fmt.Errorf("cannot get JWT validator: %w", err)
	}
	validatedToken, err := valid8r.ValidateToken(ctx, a.accessToken)
	if err != nil {
		return fmt.Errorf("cannot validate token (unauthenticated): %w", err)
	}
	token, ok := validatedToken.(*validator.ValidatedClaims)
	if !ok {
		return fmt.Errorf("cannot convert token to ValidatedClaims")
	}

	// 認証結果のContextへの格納
	ctx = metadata.WithUserInfo(ctx, &metadata.UserInfo{
		// sample
		// 認可コード（openID scope 有り）：auth0|6715cc8e39951d1cd61461ed
		// client-credentialsFlow: e0tpe0rHkPJ40AHB3M0HuXlv995CPwwq@clients
		ID: token.RegisteredClaims.Subject,
	})
	return nil
}

func (a *Auth) Authorize(ctx context.Context) error {
	//TODO implement me
	return nil
}

func getJWTValidator(cfg *provider.AuthConfig) (*validator.Validator, error) {
	issuerURL, err := url.Parse("https://" + cfg.DomainName + "/")
	if err != nil {
		return nil, fmt.Errorf("Failed to parse the issuer url: %v", err)
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
