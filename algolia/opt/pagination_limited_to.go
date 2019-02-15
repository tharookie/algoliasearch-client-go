package opt

import "encoding/json"

type PaginationLimitedToOption struct {
	value int
}

func PaginationLimitedTo(v int) PaginationLimitedToOption {
	return PaginationLimitedToOption{v}
}

func (o PaginationLimitedToOption) Get() int {
	return o.value
}

func (o PaginationLimitedToOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *PaginationLimitedToOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 1000
		return nil
	}
	return json.Unmarshal(data, &o.value)
}