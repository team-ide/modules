package kafka_service

import "github.com/team-ide/framework/kafka"

var (
	kafkaTest1 kafka.IService
)

func InitKafkaTest1(name string, cfg *kafka.Config) (err error) {
	// 如果 已存在 则 直接 返回
	if kafkaTest1 != nil {
		return
	}
	// 创建 服务
	kafkaTest1, err = NewKafkaService(name, cfg)
	if err != nil {
		return
	}

	return
}

func GetKafkaTest1() (res kafka.IService) {
	return kafkaTest1
}
