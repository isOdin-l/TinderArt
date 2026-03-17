package kafka

import (
	"context"

	"github.com/isOdin-l/TinderArt/pkg/configs"
	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Conn *kafka.Conn
}

func NewKafka(cfg *configs.ConfigKafka) (*Kafka, error) {
	conn, errCn := kafka.DialLeader(context.Background(), "tcp", cfg.DSN(), cfg.KafkaTopic, cfg.KafkaPartition)
	if errCn != nil {
		return nil, errCn
	}

	return &Kafka{Conn: conn}, nil
}

func (k *Kafka) WriteMessage(ctx context.Context, key, value []byte) error {
	_, err := k.Conn.WriteMessages(
		kafka.Message{Key: key, Value: value},
	)
	return err
}

func (k *Kafka) ReadMessage(ctx context.Context) (kafka.Message, error) {
	msg, err := k.Conn.ReadMessage(10e6) // 10MB max size
	return msg, err
}
