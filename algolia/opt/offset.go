package opt

import "encoding/json"

type OffsetOption struct {
	value int
}

func Offset(v int) OffsetOption {
	return OffsetOption{v}
}

func (o OffsetOption) Get() int {
	return o.value
}

func (o OffsetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *OffsetOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}