package common

import (
	"encoding/json"
)

const (
	ResultFailed uint8 = iota
	ResultSuccess
)

type Result struct {
	Type uint8
	Info string
}

func (this *Result) ToJSON() []byte {
	bytes, err := json.Marshal(this)
	if err != nil {
		return []byte(`{"Type":0,"Info":"ToJSON failed"}`)
	}
	return bytes
}
