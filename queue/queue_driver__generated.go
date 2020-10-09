package queue

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidQueueDriver = errors.New("invalid QueueDriver")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("QueueDriver", map[string]string{
		"REDIS":   "redis",
		"KAFKA":   "kafka",
		"BUILDIN": "buildin",
	})
}

func ParseQueueDriverFromString(s string) (QueueDriver, error) {
	switch s {
	case "":
		return QUEUE_DRIVER_UNKNOWN, nil
	case "REDIS":
		return QUEUE_DRIVER__REDIS, nil
	case "KAFKA":
		return QUEUE_DRIVER__KAFKA, nil
	case "BUILDIN":
		return QUEUE_DRIVER__BUILDIN, nil
	}
	return QUEUE_DRIVER_UNKNOWN, InvalidQueueDriver
}

func ParseQueueDriverFromLabelString(s string) (QueueDriver, error) {
	switch s {
	case "":
		return QUEUE_DRIVER_UNKNOWN, nil
	case "redis":
		return QUEUE_DRIVER__REDIS, nil
	case "kafka":
		return QUEUE_DRIVER__KAFKA, nil
	case "buildin":
		return QUEUE_DRIVER__BUILDIN, nil
	}
	return QUEUE_DRIVER_UNKNOWN, InvalidQueueDriver
}

func (QueueDriver) EnumType() string {
	return "QueueDriver"
}

func (QueueDriver) Enums() map[int][]string {
	return map[int][]string{
		int(QUEUE_DRIVER__REDIS):   {"REDIS", "redis"},
		int(QUEUE_DRIVER__KAFKA):   {"KAFKA", "kafka"},
		int(QUEUE_DRIVER__BUILDIN): {"BUILDIN", "buildin"},
	}
}

func (v QueueDriver) String() string {
	switch v {
	case QUEUE_DRIVER_UNKNOWN:
		return ""
	case QUEUE_DRIVER__REDIS:
		return "REDIS"
	case QUEUE_DRIVER__KAFKA:
		return "KAFKA"
	case QUEUE_DRIVER__BUILDIN:
		return "BUILDIN"
	}
	return "UNKNOWN"
}

func (v QueueDriver) Label() string {
	switch v {
	case QUEUE_DRIVER_UNKNOWN:
		return ""
	case QUEUE_DRIVER__REDIS:
		return "redis"
	case QUEUE_DRIVER__KAFKA:
		return "kafka"
	case QUEUE_DRIVER__BUILDIN:
		return "buildin"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*QueueDriver)(nil)

func (v QueueDriver) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidQueueDriver
	}
	return []byte(str), nil
}

func (v *QueueDriver) UnmarshalText(data []byte) (err error) {
	*v, err = ParseQueueDriverFromString(string(bytes.ToUpper(data)))
	return
}
