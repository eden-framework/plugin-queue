package queue

import (
	"context"
	"fmt"
	"github.com/eden-framework/common"
	"testing"
)

func TestProduceAndConsumeWithBuildin(t *testing.T) {
	producer := &Producer{
		Driver: DRIVER__BUILDIN,
	}
	producer.Init()

	messages := make([]common.QueueMessage, 0)
	for i := 0; i < 100; i++ {
		messages = append(messages, common.QueueMessage{
			Topic: "default",
			Key:   []byte(fmt.Sprintf("hello%d", i+1)),
			Val:   []byte(fmt.Sprintf("hello%d", i+1)),
		})
	}
	go func() {
		err := producer.Produce(context.Background(), messages...)
		if err != nil {
			t.Fatal(err)
		}
	}()

	consumer := &Consumer{
		Driver: DRIVER__BUILDIN,
	}
	consumer.Init()
	go func() {
		fmt.Println("consumer1 start to consume...")
		consumer.Consume(context.Background(), "default", func(m common.QueueMessage) error {
			fmt.Printf("consumer1 %s=%s\n", m.Key, m.Val)
			//t.Log(m)
			return nil
		})
	}()

	go func() {
		fmt.Println("consumer2 start to consume...")
		consumer.Consume(context.Background(), "default", func(m common.QueueMessage) error {
			fmt.Printf("consumer2 %s=%s\n", m.Key, m.Val)
			//t.Log(m)
			return nil
		})
	}()

	select {}
}

func TestProduceAndConsumeWithKafka(t *testing.T) {
	producer := &Producer{
		Driver: DRIVER__KAFKA,
		Host:   "localhost",
	}
	producer.Init()

	err := producer.Produce(context.Background(), common.QueueMessage{
		Topic: "default",
		Key:   []byte("hello1"),
		Val:   []byte("hello1"),
	}, common.QueueMessage{
		Topic: "default",
		Key:   []byte("hello2"),
		Val:   []byte("hello2"),
	})
	if err != nil {
		t.Fatal(err)
	}

	consumer := &Consumer{
		Driver:  DRIVER__KAFKA,
		Brokers: []string{"localhost:9092"},
		GroupID: "group1",
	}
	consumer.Init()

	go func() {
		fmt.Println("consumer1 start to consume...")
		consumer.Consume(context.Background(), "default", func(m common.QueueMessage) error {
			fmt.Print("consumer1 ")
			t.Log(m)
			return nil
		})
	}()

	go func() {
		fmt.Println("consumer2 start to consume...")
		consumer.Consume(context.Background(), "default", func(m common.QueueMessage) error {
			fmt.Print("consumer2 ")
			t.Log(m)
			return nil
		})
	}()

	select {}
}

func TestProduceAndConsumeWithRedis(t *testing.T) {
	producer := &Producer{
		Driver: DRIVER__REDIS,
		Host:   "localhost",
	}
	producer.Init()

	err := producer.Produce(context.Background(), common.QueueMessage{
		Topic: "default",
		Key:   []byte("hello1"),
		Val:   []byte("hello1"),
	}, common.QueueMessage{
		Topic: "default",
		Key:   []byte("hello2"),
		Val:   []byte("hello2"),
	})
	if err != nil {
		t.Fatal(err)
	}

	consumer := &Consumer{
		Driver:  DRIVER__REDIS,
		Brokers: []string{"localhost:6379"},
	}
	consumer.Init()

	go func() {
		fmt.Println("consumer1 start to consume...")
		consumer.Consume(context.Background(), "default", func(m common.QueueMessage) error {
			fmt.Print("consumer1 ")
			t.Log(m)
			return nil
		})
	}()

	go func() {
		fmt.Println("consumer2 start to consume...")
		consumer.Consume(context.Background(), "default", func(m common.QueueMessage) error {
			fmt.Print("consumer2 ")
			t.Log(m)
			return nil
		})
	}()

	select {}
}
