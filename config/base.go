package config

import (
	"time"
	"ycore/driver/load/file/yamlloader"
	"ycore/module/mylog"
)

var EnvInfo *Env

type Env struct {
	GRPC      *GRPC      `yaml:"grpc,omitempty"`
	Websocket *Websocket `yaml:"websocket,omitempty"`
	Redis     *Redis     `yaml:"redis,omitempty"`
	ArangoDB  *ArangoDB  `yaml:"arangoDB,omitempty"`
	Nats      *Nats      `yaml:"nats,omitempty"`
}

type GRPC struct {
	Addr string `yaml:"addr,omitempty"`
}

type Websocket struct {
	Addr           string        `yaml:"addr,omitempty"`
	Path           string        `yaml:"path,omitempty"`
	RequsetTimeOut time.Duration `yaml:"requsetTimeOut,omitempty"`
}

type Redis struct {
	Addr          string        `yaml:"addr,omitempty"`
	Password      string        `yaml:"password,omitempty"`
	PoolSize      int           `yaml:"poolSize,omitempty"`
	RetryCount    int           `yaml:"retryCount,omitempty"`
	RetryInterval time.Duration `yaml:"retryInterval,omitempty"`
}

type ArangoDB struct {
	Addr         string `yaml:"addr,omitempty"`
	Database     string `yaml:"database,omitempty"`
	Username     string `yaml:"username,omitempty"`
	Password     string `yaml:"password,omitempty"`
	HttpProtocol string `yaml:"httpProtocol,omitempty"`
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

func Init(path string) error {
	if err := yamlloader.LoadYaml(path, &EnvInfo); err != nil {
		mylog.Errorf("[Config][Init] load Error err: %v", err)
		return err
	}
	return nil
}
