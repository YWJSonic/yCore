package myredis

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/YWJSonic/ycore/driver/database/redis"
	"github.com/YWJSonic/ycore/module/mylog"
)

// New redis manager
func New(addrT any, password string, poolSize int) (*Manager, error) {

	ctx, cancel := context.WithCancel(context.TODO())
	obj := &Manager{
		ctx:    ctx,
		cancel: cancel,
	}

	switch iaddr := addrT.(type) {
	case string:
		redisIPs := strings.Split(iaddr, ",")
		for _, ip := range redisIPs {
			_, err := url.Parse("http://" + ip)
			if err != nil {
				mylog.Errorf("[Redis][New] url Parse error: %v", err)
				return nil, err
			}
		}
		driver, err := redis.New(redisIPs, password, poolSize)
		if err != nil {
			mylog.Errorf("[Redis][init] connection error, err: %v", err)
			return nil, err
		}

		obj.Cmdable = driver
	case []string:
		driver, err := redis.New(iaddr, password, poolSize)
		if err != nil {
			mylog.Errorf("[Redis][init] connection error, err: %v", err)
			return nil, err
		}

		obj.Cmdable = driver
	default:
		return nil, fmt.Errorf("[Redis][init] addr type error addr: %v", addrT)
	}

	go redis.PingLoop(ctx, cancel, obj.Cmdable)

	mylog.Infof("[Redis][New] Connect success, address: %v", addrT)
	return obj, nil
}
