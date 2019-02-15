package opt

import "encoding/json"

type OptionalWordsOption struct {
	value []string
}

func OptionalWords(v []string) OptionalWordsOption {
	return OptionalWordsOption{v}
}

func (o OptionalWordsOption) Get() []string {
	return o.value
}

func (o OptionalWordsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *OptionalWordsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}