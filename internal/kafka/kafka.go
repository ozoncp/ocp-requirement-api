package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-requirement-api/internal/config"
	"github.com/rs/zerolog/log"
	"time"
)

type EventType int

const (
	Create EventType = iota
	Update
	Remove
)

func (e EventType) String() string {
	switch e {
	case Create:
		return "Create"
	case Update:
		return "Update"
	case Remove:
		return "Remove"
	default:
		return "Undefined"
	}
}

type Message struct {
	EventType    string       `json:"event_type"`
	EventContent EventContent `json:"event_content"`
}

type EventContent struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	Text   string `json:"text"`
}

func CreateMessage(e string, eventMessage EventContent) Message {
	return Message{
		EventType:    e,
		EventContent: eventMessage,
	}
}

type Producer interface {
	Send(message Message) error
}

type producer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewProducer(cfg config.Config) (*producer, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.Kafka.Addresses, kafkaConfig)
	if err != nil {
		return nil, err
	}

	p := &producer{
		producer: syncProducer,
		topic:    cfg.Kafka.Topic,
	}
	return p, nil
}

func (p *producer) Send(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Err(err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(p.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.producer.SendMessage(msg)
	return err
}
