package auth_intercepter

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/lib/auth"
	"github.com/tkame123/ddd-sample/lib/metadata"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type authStrategyJWT struct {
	cfg         *provider.AuthConfig
	accessToken string
	obj         string
	enforcer    *casbin.Enforcer
}

func NewAuthStrategyJWT(
	cfg *provider.AuthConfig,
	enforcer *casbin.Enforcer,
	accessToken string,
	obj string,
) auth.Strategy {
	return &authStrategyJWT{
		cfg:         cfg,
		enforcer:    enforcer,
		accessToken: accessToken,
		obj:         obj,
	}
}

func (a *authStrategyJWT) Authenticate(ctx context.Context) (*metadata.UserInfo, error) {
	//MEMO: 厳密にはAccessTokenの検証であり認証といっていいのかが疑問ではある。。

	valid8r, err := getJWTValidator(a.cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot get JWT validator: %w", err)
	}
	validatedToken, err := valid8r.ValidateToken(ctx, a.accessToken)
	if err != nil {
		return nil, &AuthenticationError{cause: err}
	}

	token, ok := validatedToken.(*validator.ValidatedClaims)
	if !ok {
		return nil, fmt.Errorf("cannot convert token to ValidatedClaims")
	}
	claims, ok := token.CustomClaims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("cannot convert token to CustomClaims")
	}
	return &metadata.UserInfo{
		// auth0 sample
		// 認可コード（openID scope 有り）：auth0|6715cc8e39951d1cd61461ed
		// client-credentialsFlow: e0tpe0rHkPJ40AHB3M0HuXlv995CPwwq@clients
		ID:           token.RegisteredClaims.Subject,
		AccessPolicy: &metadata.AccessPolicy{Permissions: claims.Permissions},
	}, nil
}

func (a *authStrategyJWT) Authorize(ctx context.Context) error {
	userInfo, ok := metadata.GetUserInfo(ctx)
	if !ok {
		return fmt.Errorf("cannot get user info from context")
	}

	for _, permission := range userInfo.AccessPolicy.Permissions {
		ok, err := a.enforcer.Enforce(
			auth.Permission(permission).Sub(),
			a.obj,
			auth.Permission(permission).Act(),
		)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
	}
	return &PermissionError{}
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
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	return jwtValidator, nil
}
