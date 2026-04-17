package start

import (
	"example/component"
	"example/component/db_service"
	"example/component/es_service"
	"example/component/kafka_service"
	"example/component/mongodb_service"
	"example/component/redis_service"
	"example/component/zk_service"
	"example/config"
	"github.com/team-ide/framework"
)

func init() {
	component.Starter.AddInitComponentFunc("init component", 100, InitComponent)
}

func InitComponent() (err error) {
	framework.Info("组件 初始化 开始")

	cfg := config.GetConfig()

	if err = db_service.InitDbTest1("db_test_1", cfg.DbTest1); err != nil {
		return
	}

	if err = db_service.InitDbTest1("db_test_2", cfg.DbTest2); err != nil {
		return
	}

	if err = redis_service.InitRedisTest1("redis_test_1", cfg.RedisTest1); err != nil {
		return
	}

	if err = zk_service.InitZkTest1("zk_test_1", cfg.ZkTest1); err != nil {
		return
	}

	if err = kafka_service.InitKafkaTest1("kafka_test_1", cfg.KafkaTest1); err != nil {
		return
	}

	if err = es_service.InitEsTest1("es_test_1", cfg.EsTest1); err != nil {
		return
	}

	if err = mongodb_service.InitMongodbTest1("mongodb_test_1", cfg.MongodbTest1); err != nil {
		return
	}

	framework.Info("组件 初始化 结束")
	return
}
