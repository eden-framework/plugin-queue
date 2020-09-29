package queue

type Producer struct {
	Driver Driver
	Host   string
	Port   int
	Topic  string
	p      ProducerDriver
}
