package app

import (
	"context"
	"fmt"
)

type Auth interface {
	GetUID(token string) (string, error)
}

// unexported key type prevents collisions
type key int

const (
	userIdKey key = 0
)

// Возвращает новый контекст с User Id.
func WithUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, userIdKey, uid)
}

// Извлекает User Id из контекста.
func UserIdFromContext(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(userIdKey).(string)
	if !ok {
		return "", fmt.Errorf("Context missing User ID")
	}
	return userId, nil
}
