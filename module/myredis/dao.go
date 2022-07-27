package myredis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Manager struct {
	redis.Cmdable
	ctx    context.Context
	cancel context.CancelFunc
}
