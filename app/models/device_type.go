package models

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type DeviceType int

const (
	Web DeviceType = iota
	Iphone
	Android
	Other
)

var toStringMap = map[DeviceType]string{
	Web:     "web",
	Iphone:  "iphone",
	Android: "android",
	Other:   "other",
}

var toEnumMap = map[string]DeviceType{
	"web":     Web,
	"iphone":  Iphone,
	"android": Android,
	"other":   Other,
}

func (d DeviceType) String() string {
	return toStringMap[d]
}

func (d *DeviceType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toStringMap[*d])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (d *DeviceType) UnmarshalJSON(b []byte) error {
	// unmarshal as string
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	if value, ok := toEnumMap[s]; ok {
		*d = value
	} else {
		return errors.Errorf("Failed to parse enum DeviceType").Wrap("value: %s", s)
	}
	return nil
}
