package queue

import (
	"github.com/eden-framework/plugin-kafka/kafka"
	"github.com/eden-framework/plugin-redis/redis"
	"github.com/profzone/envconfig"
)

type Producer struct {
	Driver   Driver
	Host     string
	Port     int
	User     string
	Password envconfig.Password
	Topic    string
	producerDriver
}

func (p *Producer) SetDefault() {
	if p.Driver == DRIVER_UNKNOWN {
		p.Driver = DRIVER__BUILDIN
	}
	if p.Driver == DRIVER__REDIS || p.Driver == DRIVER__KAFKA {
		if p.Host == "" {
			panic("[Producer] must specify Host and Port when use REDIS or KAFKA drivers")
		}
	}
	if p.Topic == "" {
		p.Topic = "default"
	}
}

func (p *Producer) Init() {
	p.SetDefault()
	switch p.Driver {
	case DRIVER__BUILDIN:
	case DRIVER__REDIS:
		driver := &redis.Redis{
			Host:     p.Host,
			Port:     p.Port,
			User:     p.User,
			Password: p.Password,
			Topic:    p.Topic,
		}
		driver.Init()
		p.producerDriver = driver
	case DRIVER__KAFKA:
		driver := &kafka.Producer{
			Host:  p.Host,
			Port:  p.Port,
			Topic: p.Topic,
		}
		driver.Init()
		p.producerDriver = driver
	default:
		panic("[Producer] unsupported driver")
	}
}
