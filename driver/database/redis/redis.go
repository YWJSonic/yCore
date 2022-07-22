package redis

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// New redis manager
func New(addr, password string, poolSize int) (*Manager, error) {

	ctx, cancel := context.WithCancel(context.TODO())
	obj := &Manager{
		redisClient: nil,
		ctx:         ctx,
	}

	redisClient, err := obj.init(addr, password, poolSize, ctx, cancel)
	if err != nil {
		fmt.Printf("[Redis][New] Init error: %v", err)
		return nil, err
	}

	obj.redisClient = redisClient

	fmt.Printf("[Redis][New] Connect success, address: %v", addr)
	return obj, nil
}

type Manager struct {
	redisClient redis.Cmdable
	ctx         context.Context
	// mu          sync.RWMutex
}

func (mgr *Manager) init(addr, password string, poolSize int, ctx context.Context, cancel context.CancelFunc) (redis.Cmdable, error) {
	redisIPs := strings.Split(addr, ",")
	for _, ip := range redisIPs {
		_, err := url.Parse("http://" + ip)
		if err != nil {
			fmt.Printf("[Redis][init] URL parse error: %v", err)
			<-ctx.Done()
			cancel()
			return nil, err
		}
	}

	client, err := mgr.connection(redisIPs, password, poolSize)
	if err != nil {
		fmt.Printf("[Redis][init] connection error, err: %v", err)
		return nil, err
	}

	go pingLoop(ctx, client, cancel)

	return client, nil
}

func (mgr *Manager) connection(redisIPs []string, password string, poolSize int) (redis.Cmdable, error) {
	switch len(redisIPs) {
	case 0: // error
		return nil, errors.New("[Redis][connection] IP address null")
	case 1: // single
		redisClient := redis.NewClient(&redis.Options{
			Addr:       redisIPs[0],
			Password:   password,
			PoolSize:   poolSize,
			MaxConnAge: 1 * time.Hour,
		})
		// redisClient.AddHook(apmgoredis.NewHook())

		return redisClient, nil
	default: // cluster
		redisClient := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      redisIPs,
			Password:   password,
			PoolSize:   poolSize,
			MaxConnAge: 1 * time.Hour,
		})
		// redisClient.AddHook(apmgoredis.NewHook())

		return redisClient, nil

	}
}

func pingLoop(ctx context.Context, redisClient redis.Cmdable, cancel context.CancelFunc) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if _, err := redisClient.Ping(ctx).Result(); err != nil {
				fmt.Printf("[Redis][pingLoop] ping error, err: %v", err.Error())
				return
			}
		case <-ctx.Done():
			cancel()
			return
		}
	}
}
