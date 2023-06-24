package adapters

import "github.com/otaviohenrique/go-app-example/pkg/ports"

type SqsProducerMessage struct{}

func NewSqsProducerMessage() *SqsProducerMessage {
	result := new(SqsProducerMessage)

	return result
}

func (result *SqsProducerMessage) Read() []byte {
	return []byte("Hello World")
}

type SqsProducer struct{}

func NewSqsProducer() *SqsProducer {
	sqs := new(SqsProducer)

	return sqs
}

func (sqs *SqsProducer) Produce(_ ports.ProducerMessage) error {
	return nil
}
