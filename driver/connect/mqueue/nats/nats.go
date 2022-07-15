package nats

import (
	"errors"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type NatsModule struct {
	conn              stan.Conn
	managementTimeOut time.Duration
}

func New() *NatsModule {
	module := &NatsModule{}

	return module
}

func (self *NatsModule) LaunchNats(addr, username, password, stanClusterID, clientID string) error {

	if addr == "" {
		return errors.New(".env NATS_ADDR is nil")
	}

	nats, err := nats.Connect(addr, nats.UserInfo(username, password))
	if err != nil {
		log.Fatalf("[Nats][LaunchNats] nats.Connect failed err: %v", err)
	}

	if stanClusterID == "" {
		return errors.New(".env NATS_CLUSTER_ID is nil")
	}

	conn, err := stan.Connect(
		stanClusterID,
		clientID,
		stan.NatsConn(nats),
		stan.Pings(3, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("[Nats][LaunchNats] Connection lost, reason: %v", reason)
		}))

	if err != nil {
		log.Fatalf("[Nats][LaunchNats] Connection lost, err: %v", err)
	}

	self.conn = conn
	return err
}
