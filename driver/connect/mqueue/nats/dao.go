package nats

type NatsEnv struct {
	Addr          string
	Username      string
	Password      string
	StanClusterID string
	ClientID      string
	PingsInterval int
	PingsMaxOut   int
}
