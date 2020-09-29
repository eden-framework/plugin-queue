package queue

import (
	"context"
	"github.com/eden-framework/common"
)

//go:generate eden generate enum --type-name=Driver
// api:enum
type Driver uint8

// queue driver type
const (
	DRIVER_UNKNOWN  Driver = iota
	DRIVER__BUILDIN        // buildin
	DRIVER__KAFKA          // kafka
	DRIVER__REDIS          // redis
)

type ConsumerDriver interface {
	Consume(ctx context.Context, handler func(m common.QueueMessage) error) error
}

type ProducerDriver interface {
	Produce(ctx context.Context, messages ...common.QueueMessage) error
}