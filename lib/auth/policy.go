package auth

import "github.com/tkame123/ddd-sample/lib/metadata"

type Permission string

const (
	OrdersRead  metadata.Permission = "read:orders"
	OrdersWrite metadata.Permission = "write:orders"
)
