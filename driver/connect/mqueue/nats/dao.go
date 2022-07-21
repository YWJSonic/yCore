package nats

import (
	"github.com/nats-io/stan.go"
)

type NatsDriver struct {
	conn stan.Conn
}

type NatsEnv struct {
	Addr          string
	Username      string
	Password      string
	StanClusterID string
	ClientID      string
	PingsInterval int
	PingsMaxOut   int
}
