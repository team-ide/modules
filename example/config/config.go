package config

import (
	"example/component"
	"flag"
	"path/filepath"
	"strings"

	"github.com/team-ide/framework"
	"github.com/team-ide/framework/db"
	"github.com/team-ide/framework/es"
	"github.com/team-ide/framework/kafka"
	"github.com/team-ide/framework/mongodb"
	"github.com/team-ide/framework/redis"
	"github.com/team-ide/framework/util"
	"github.com/team-ide/framework/web"
	"github.com/team-ide/framework/zookeeper"
)

var (
	defaultConfig = "conf/config.yml"
	argConfig     = flag.String("c", "", "配置文件")
)

func GetConfigPath() string {
	if argConfig != nil && *argConfig != "" {
		return *argConfig
	}
	return defaultConfig
}

var (
	config     *Config
	configPath string
)

type Config struct {
	WebApi    *web.Config `json:"web_api,omitempty" yaml:"web_api,omitempty"`
	WebManage *web.Config `json:"web_manage,omitempty" yaml:"web_manage,omitempty"`

	DbTest1 *db.Config `json:"db_test_1,omitempty" yaml:"db_test_1,omitempty"`
	DbTest2 *db.Config `json:"db_test_2,omitempty" yaml:"db_test_2,omitempty"`

	RedisTest1 *redis.Config `json:"redis_test_1,omitempty" yaml:"redis_test_1,omitempty"`

	ZkTest1 *zookeeper.Config `json:"zk_test_1,omitempty" yaml:"zk_test_1,omitempty"`

	EsTest1 *es.Config `json:"es_test_1,omitempty" yaml:"es_test_1,omitempty"`

	KafkaTest1 *kafka.Config `json:"kafka_test_1,omitempty" yaml:"kafka_test_1,omitempty"`

	MongodbTest1 *mongodb.Config `json:"mongodb_test_1,omitempty" yaml:"mongodb_test_1,omitempty"`

	Log *framework.LogConfig `json:"log,omitempty" yaml:"log,omitempty"`
}

func GetConfig() *Config {
	return config
}
func SetConfig(c *Config) {
	config = c
}

func init() {
	component.Starter.AddInitConfigFunc("init config", 0, initConfig)
	component.Starter.AddInitConfigFunc("init logger config", 100, initLoggerConfig)
}

func initConfig() (err error) {
	if config != nil {
		return
	}
	conf := GetConfigPath()
	if !strings.HasPrefix(conf, "/") {
		conf = filepath.Join(framework.GetPwdDir(), conf)
	}

	cnf, err := framework.ReadConfig[*Config](util.FormatPath(conf))
	if err != nil {
		return
	}

	config = cnf

	return
}

func initLoggerConfig() (err error) {
	c := GetConfig().Log
	if c == nil {
		return
	}
	framework.LoggerInit(c)
	return
}
