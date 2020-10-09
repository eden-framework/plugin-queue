package queue

import (
	"fmt"
	"github.com/eden-framework/plugin-kafka/kafka"
	"github.com/eden-framework/plugin-redis/redis"
	"github.com/profzone/envconfig"
	"net"
	"strconv"
)

type Consumer struct {
	Driver QueueDriver
	// The list of broker addresses used to connect to the kafka cluster.
	Brokers []string

	// GroupID holds the optional consumer group id.  If GroupID is specified, then
	// Partition should NOT be specified e.g. 0
	GroupID string

	// Partition to read messages from.  Either Partition or GroupID may
	// be assigned, but not both
	Partition int
	User      string
	Password  envconfig.Password
	consumerDriver
}

func (c *Consumer) SetDefault() {
	if c.Driver == QUEUE_DRIVER_UNKNOWN {
		c.Driver = QUEUE_DRIVER__BUILDIN
	}
	if c.Driver == QUEUE_DRIVER__REDIS || c.Driver == QUEUE_DRIVER__KAFKA {
		if len(c.Brokers) == 0 {
			panic("[Consumer] must specify Broker list when use REDIS or KAFKA drivers")
		}
	}
	if c.Driver == QUEUE_DRIVER__KAFKA && c.GroupID == "" {
		panic("[Consumer] must specify GroupID when use KAFKA driver")
	}
}

func (c *Consumer) Init() {
	c.SetDefault()
	switch c.Driver {
	case QUEUE_DRIVER__BUILDIN:
		c.consumerDriver = newMemoryQueue(100)
	case QUEUE_DRIVER__REDIS:
		host, port, err := net.SplitHostPort(c.Brokers[0])
		if err != nil {
			panic(fmt.Sprintf("[Consumer] SplitHostPort from broker %s failed: %v", c.Brokers[0], err))
		}
		portInteger, err := strconv.ParseInt(port, 10, 32)
		if err != nil {
			panic(fmt.Sprintf("[Consumer] Parse port %s to integer failed: %v", port, err))
		}
		driver := &redis.Redis{
			Host:     host,
			Port:     int(portInteger),
			User:     c.User,
			Password: c.Password,
		}
		driver.Init()
		c.consumerDriver = driver
	case QUEUE_DRIVER__KAFKA:
		driver := &kafka.Consumer{
			Brokers:   c.Brokers,
			GroupID:   c.GroupID,
			Partition: c.Partition,
		}
		driver.Init()
		c.consumerDriver = driver
	default:
		panic("[Producer] unsupported driver")
	}
}
