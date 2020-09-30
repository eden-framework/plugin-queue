package queue

import (
	"context"
	"github.com/derekparker/trie"
	"github.com/eden-framework/common"
)

var defaultMemoryQueue *memoryQueue

type memoryQueue struct {
	tree *trie.Trie
}

func newMemoryQueue() *memoryQueue {
	if defaultMemoryQueue == nil {
		defaultMemoryQueue = &memoryQueue{
			tree: trie.New(),
		}
	}
	return defaultMemoryQueue
}

func (m *memoryQueue) Consume(ctx context.Context, topic string, handler func(m common.QueueMessage) error) error {
	panic("implement me")
}

func (m *memoryQueue) Produce(ctx context.Context, messages ...common.QueueMessage) error {
	panic("implement me")
}
