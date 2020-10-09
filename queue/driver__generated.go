package queue

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidDriver = errors.New("invalid Driver")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("Driver", map[string]string{
		"REDIS":   "redis",
		"KAFKA":   "kafka",
		"BUILDIN": "buildin",
	})
}

func ParseDriverFromString(s string) (Driver, error) {
	switch s {
	case "":
		return DRIVER_UNKNOWN, nil
	case "REDIS":
		return DRIVER__REDIS, nil
	case "KAFKA":
		return DRIVER__KAFKA, nil
	case "BUILDIN":
		return DRIVER__BUILDIN, nil
	}
	return DRIVER_UNKNOWN, InvalidDriver
}

func ParseDriverFromLabelString(s string) (Driver, error) {
	switch s {
	case "":
		return DRIVER_UNKNOWN, nil
	case "redis":
		return DRIVER__REDIS, nil
	case "kafka":
		return DRIVER__KAFKA, nil
	case "buildin":
		return DRIVER__BUILDIN, nil
	}
	return DRIVER_UNKNOWN, InvalidDriver
}

func (Driver) EnumType() string {
	return "Driver"
}

func (Driver) Enums() map[int][]string {
	return map[int][]string{
		int(DRIVER__REDIS):   {"REDIS", "redis"},
		int(DRIVER__KAFKA):   {"KAFKA", "kafka"},
		int(DRIVER__BUILDIN): {"BUILDIN", "buildin"},
	}
}

func (v Driver) String() string {
	switch v {
	case DRIVER_UNKNOWN:
		return ""
	case DRIVER__REDIS:
		return "REDIS"
	case DRIVER__KAFKA:
		return "KAFKA"
	case DRIVER__BUILDIN:
		return "BUILDIN"
	}
	return "UNKNOWN"
}

func (v Driver) Label() string {
	switch v {
	case DRIVER_UNKNOWN:
		return ""
	case DRIVER__REDIS:
		return "redis"
	case DRIVER__KAFKA:
		return "kafka"
	case DRIVER__BUILDIN:
		return "buildin"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*Driver)(nil)

func (v Driver) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDriver
	}
	return []byte(str), nil
}

func (v *Driver) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDriverFromString(string(bytes.ToUpper(data)))
	return
}
