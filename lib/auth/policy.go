package auth

import (
	"strings"
)

type Permission string

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
