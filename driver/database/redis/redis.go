package redis

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"
	"ycore/module/mylog"

	"github.com/go-redis/redis/v8"
)

func New(addr, password string, poolSize int) (redis.Cmdable, error) {
	redisIPs := strings.Split(addr, ",")
	for _, ip := range redisIPs {
		_, err := url.Parse("http://" + ip)
		if err != nil {
			mylog.Errorf("[Redis][New] url Parse error: %v", err)
			return nil, err
		}
	}

	driver, err := connection(redisIPs, password, poolSize)
	if err != nil {
		mylog.Errorf("[RedisDriver][New] connection error: %v", err)
		return nil, err
	}

	mylog.Infof("[RedisDriver][New] Connect success, address: %v", addr)
	return driver, nil
}

func connection(redisIPs []string, password string, poolSize int) (redis.Cmdable, error) {
	switch len(redisIPs) {
	case 0: // error
		return nil, errors.New("[RedisDriver][connection] IP address null")

	case 1: // single
		redisClient := redis.NewClient(&redis.Options{
			Addr:       redisIPs[0],
			Password:   password,
			PoolSize:   poolSize,
			MaxConnAge: 1 * time.Hour,
		})
		return redisClient, nil

	default: // cluster
		redisClient := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      redisIPs,
			Password:   password,
			PoolSize:   poolSize,
			MaxConnAge: 1 * time.Hour,
		})
		return redisClient, nil

	}
}

func PingLoop(ctx context.Context, cancel context.CancelFunc, redisClient redis.Cmdable) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			if _, err := redisClient.Ping(ctx).Result(); err != nil {
				cancel()
				mylog.Errorf("[Redis][pingLoop] ping error, err: %v", err.Error())
				return
			}
		}
	}
}
