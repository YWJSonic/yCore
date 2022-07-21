package myyaml

import "time"

type Parameter struct {
	Redis    *Redis    `yaml:"redis,omitempty"`
	ArangoDB *ArangoDB `yaml:"arangoDB,omitempty"`
	Nats     *Nats     `yaml:"nats,omitempty"`
}

type Redis struct {
	Addr          string        `yaml:"addr,omitempty"`
	Password      string        `yaml:"password,omitempty"`
	PoolSize      int           `yaml:"poolSize,omitempty"`
	RetryCount    int           `yaml:"retryCount,omitempty"`
	RetryInterval time.Duration `yaml:"retryInterval,omitempty"`
}

type ArangoDB struct {
	Addr          string        `yaml:"addr,omitempty"`
	Database      string        `yaml:"database,omitempty"`
	Username      string        `yaml:"username,omitempty"`
	Password      string        `yaml:"password,omitempty"`
	RetryCount    int           `yaml:"retryCount,omitempty"`
	RetryInterval time.Duration `yaml:"retryInterval,omitempty"`
	HttpProtocol  string        `yaml:"httpProtocol,omitempty"`
}

type Nats struct {
	Addr              string        `yaml:"addr,omitempty"`
	Username          string        `yaml:"username,omitempty"`
	Password          string        `yaml:"password,omitempty"`
	ClusterID         string        `yaml:"clusterID,omitempty"`
	ReconnInterval    time.Duration `yaml:"reconnInterval,omitempty"`
	ConnectTimeOut    time.Duration `yaml:"connectTimeOut,omitempty"`
	StanPingsInterval time.Duration `yaml:"stanPingsInterval,omitempty"`
	StanPingsMaxOut   int           `yaml:"stanPingsMaxOut,omitempty"`
}
