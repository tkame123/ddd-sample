package provider

import "github.com/casbin/casbin/v2"

func NewCasbinEnforcer() (*casbin.Enforcer, error) {
	e, err := casbin.NewEnforcer("app/order_api/casbin_model.conf", "app/order_api/casbin_policy.csv")
	if err != nil {
		return nil, err
	}

	return e, nil
}
