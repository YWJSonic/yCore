package nats

import (
	"errors"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func New() *NatsDriver {
	module := &NatsDriver{}

	return module
}

func (self *NatsDriver) LaunchNats(env *NatsEnv) error {

	if env.Addr == "" {
		return errors.New(".env NATS_ADDR is nil")
	}

	nats, err := nats.Connect(env.Addr, nats.UserInfo(env.Username, env.Password))
	if err != nil {
		log.Fatalf("[Nats][LaunchNats] nats.Connect failed err: %v", err)
	}

	if env.StanClusterID == "" {
		return errors.New(".env NATS_CLUSTER_ID is nil")
	}

	conn, err := stan.Connect(
		env.StanClusterID,
		env.ClientID,
		stan.NatsConn(nats),
		stan.Pings(env.PingsInterval, env.PingsMaxOut),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("[Nats][LaunchNats] Connection lost, reason: %v", reason)
		}))

	if err != nil {
		log.Fatalf("[Nats][LaunchNats] Connection lost, err: %v", err)
	}

	self.conn = conn
	return err
}
