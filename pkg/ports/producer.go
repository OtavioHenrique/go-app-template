package ports

type ProducerMessage interface {
	Read() []byte
}

type Producer interface {
	Produce(ProducerMessage) error
}
