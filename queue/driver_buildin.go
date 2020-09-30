package queue

import (
	"context"
	"fmt"
	"github.com/cornelk/hashmap"
	"github.com/eden-framework/common"
	"time"
)

var defaultMemoryQueue *memoryQueue

type memoryQueue struct {
	container *hashmap.HashMap
	batchSize int
}

func newMemoryQueue(batchSize int) *memoryQueue {
	if defaultMemoryQueue == nil {
		defaultMemoryQueue = &memoryQueue{
			container: hashmap.New(10),
			batchSize: batchSize,
		}
	}
	return defaultMemoryQueue
}

func (m *memoryQueue) Consume(ctx context.Context, topic string, handler func(m common.QueueMessage) error) error {
	prepared := make(chan common.QueueMessage, m.batchSize)
	actual, _ := m.container.GetOrInsert(topic, &prepared)
	ch := actual.(*chan common.QueueMessage)

Run:
	for {
		select {
		case <-ctx.Done():
			break Run
		case m := <-*ch:
			err := handler(m)
			if err != nil {
				*ch <- m
			}
		}
	}

	return nil
}

func (m *memoryQueue) Produce(ctx context.Context, messages ...common.QueueMessage) error {
	for _, msg := range messages {
		if msg.Topic == "" {
			return fmt.Errorf("[memoryQueue] topic of message is not specified")
		}
	}

	for _, msg := range messages {
		msg.Time = time.Now()
		prepared := make(chan common.QueueMessage, m.batchSize)
		actual, _ := m.container.GetOrInsert(msg.Topic, &prepared)
		ch := actual.(*chan common.QueueMessage)

		*ch <- msg
	}

	return nil
}
