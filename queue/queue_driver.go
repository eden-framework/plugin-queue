package queue

import (
	"context"
	"github.com/eden-framework/common"
)

//go:generate eden generate enum --type-name=QueueDriver
// api:enum
type QueueDriver uint8

// queue driver type
const (
	QUEUE_DRIVER_UNKNOWN  QueueDriver = iota
	QUEUE_DRIVER__BUILDIN             // buildin
	QUEUE_DRIVER__KAFKA               // kafka
	QUEUE_DRIVER__REDIS               // redis
)

type consumerDriver interface {
	Consume(ctx context.Context, topic string, handler func(m common.QueueMessage) error) error
}

type producerDriver interface {
	Produce(ctx context.Context, messages ...common.QueueMessage) error
}
