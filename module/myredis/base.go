package myredis

import (
	"context"

	"github.com/YWJSonic/ycore/driver/database/redis"
	"github.com/YWJSonic/ycore/module/mylog"
)

// New redis manager
func New(addr, password string, poolSize int) (*Manager, error) {

	ctx, cancel := context.WithCancel(context.TODO())
	obj := &Manager{
		ctx:    ctx,
		cancel: cancel,
	}

	driver, err := redis.New(addr, password, poolSize)
	if err != nil {
		mylog.Errorf("[Redis][init] connection error, err: %v", err)
		return nil, err
	}

	obj.Cmdable = driver
	go redis.PingLoop(ctx, cancel, driver)

	mylog.Infof("[Redis][New] Connect success, address: %v", addr)
	return obj, nil
}
