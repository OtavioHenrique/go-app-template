package ports

type ConsumerResult interface {
	Read() []byte
}

type Consumer interface {
	Consume() (ConsumerResult, error)
}
