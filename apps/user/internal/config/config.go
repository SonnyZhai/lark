package config

import (
	"flag"
	"github.com/spf13/cast"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name          string              `yaml:"name"`
	ServerID      int                 `yaml:"server_id"`
	Log           string              `yaml:"log"`
	GrpcServer    *conf.Grpc          `yaml:"grpc_server"`
	Etcd          *conf.Etcd          `yaml:"etcd"`
	Mysql         *conf.Mysql         `yaml:"mysql"`
	Redis         *conf.Redis         `yaml:"redis"`
	Elasticsearch *conf.Elasticsearch `yaml:"elasticsearch"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/user.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	grpcPort = flag.Int("gp", 7500, "grpc server port")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId
	config.GrpcServer.ServerID = config.ServerID
	config.GrpcServer.Port = *grpcPort

	xsnowflake.NewSnowflake(config.ServerID)
	xlog.Shared(config.Log, config.Name+cast.ToString(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
