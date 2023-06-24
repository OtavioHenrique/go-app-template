package adapters

import "github.com/otaviohenrique/go-app-example/pkg/ports"

type SqsConsumerResult struct{}

func NewSqsConsumerResult() *SqsConsumerResult {
	result := new(SqsConsumerResult)

	return result
}

func (result *SqsConsumerResult) Read() []byte {
	return []byte("Hello World")
}

type SqsConsumer struct{}

func NewSqsConsumer() *SqsConsumer {
	sqs := new(SqsConsumer)

	return sqs
}

func (sqs *SqsConsumer) Consume() (ports.ConsumerResult, error) {
	return NewSqsConsumerResult(), nil
}
