package kafka

import (
	"context"

	"github.com/isOdin-l/TinderArt/pkg/configs"
	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Conn *kafka.Writer
}

func NewKafka(cfg *configs.ConfigKafka) (*Kafka, error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.DSN()),
		Topic:    cfg.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Kafka{Conn: writer}, nil
}

func (k *Kafka) WriteMessage(ctx context.Context, key, value []byte) error {
	return k.Conn.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}
