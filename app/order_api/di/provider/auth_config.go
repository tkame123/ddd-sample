package provider

type AuthConfig struct {
	DomainName   string
	AudienceName string
}

func NewAuthConfig(envCfg *EnvConfig) *AuthConfig {
	return &AuthConfig{
		DomainName:   envCfg.AuthDomainName,
		AudienceName: envCfg.AuthAudienceName,
	}
}
