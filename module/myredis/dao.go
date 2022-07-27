package myredis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Manager struct {
	driver redis.Cmdable
	ctx    context.Context
	cancel context.CancelFunc
}
