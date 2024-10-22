package auth

import (
	"github.com/tkame123/ddd-sample/lib/metadata"
	"strings"
)

type Permission string

const (
	OrdersRead  metadata.Permission = "read:orders"
	OrdersWrite metadata.Permission = "write:orders"
)

func (p Permission) Sub() string {
	result := strings.Split(string(p), ":")
	if len(result) < 2 {
		return ""
	}
	return result[1]
}

func (p Permission) Act() string {
	result := strings.Split(string(p), ":")
	if len(result) < 2 {
		return ""
	}
	return result[0]
}
