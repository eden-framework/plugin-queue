package queue

import (
	"context"
	"fmt"
	"github.com/eden-framework/common"
	"testing"
)

func TestProduceAndConsumeWithKafka(t *testing.T) {
	producer := &Producer{
		Driver: DRIVER__KAFKA,
		Host:   "localhost",
		Topic:  "default",
	}
	producer.Init()

	producer.Produce(context.Background(), common.QueueMessage{
		Key: []byte("hello1"),
		Val: []byte("hello1"),
	}, common.QueueMessage{
		Key: []byte("hello2"),
		Val: []byte("hello2"),
	})

	consumer := &Consumer{
		Driver:  DRIVER__KAFKA,
		Brokers: []string{"localhost:9092"},
		Topic:   "default",
		GroupID: "group1",
	}
	consumer.Init()
	fmt.Println("start to consume...")
	consumer.Consume(context.Background(), func(m common.QueueMessage) error {
		t.Log(m)
		return nil
	})
}

func TestProduceAndConsumeWithRedis(t *testing.T) {
	producer := &Producer{
		Driver: DRIVER__REDIS,
		Host:   "localhost",
		Topic:  "default",
	}
	producer.Init()

	producer.Produce(context.Background(), common.QueueMessage{
		Key: []byte("hello1"),
		Val: []byte("hello1"),
	}, common.QueueMessage{
		Key: []byte("hello2"),
		Val: []byte("hello2"),
	})

	go func() {
		consumer1 := &Consumer{
			Driver:  DRIVER__REDIS,
			Brokers: []string{"localhost:6379"},
			Topic:   "default",
		}
		consumer1.Init()
		fmt.Println("consumer1 start to consume...")
		consumer1.Consume(context.Background(), func(m common.QueueMessage) error {
			fmt.Print("consumer1 ")
			t.Log(m)
			return nil
		})
	}()

	go func() {
		consumer2 := &Consumer{
			Driver:  DRIVER__REDIS,
			Brokers: []string{"localhost:6379"},
			Topic:   "default",
		}
		consumer2.Init()
		fmt.Println("consumer2 start to consume...")
		consumer2.Consume(context.Background(), func(m common.QueueMessage) error {
			fmt.Print("consumer2 ")
			t.Log(m)
			return nil
		})
	}()

	select {}
}
