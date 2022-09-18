package app

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

func ConsumeKafkaStream() error {
	kafkaConfig := createKafkaConfig()
	topics := []string{"kafkaTest"}

	consumer, err := kafka.NewConsumer(&kafkaConfig)

	if err != nil {
		return fmt.Errorf("error while creating kafka config %v", err)
	}

	defer consumer.Close()
	consumer.SubscribeTopics(topics, nil)

	fmt.Printf("Listening....")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s:\n%s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func createKafkaConfig() kafka.ConfigMap {
	return kafka.ConfigMap{
		"bootstrap.servers": viper.Get("bootstrap_servers"),
		"group.id":          viper.GetString("group_id"),
		"auto.offset.reset": "earliest",
	}
}
