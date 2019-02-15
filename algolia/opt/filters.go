package opt

import "encoding/json"

type FiltersOption struct {
	value string
}

func Filters(v string) FiltersOption {
	return FiltersOption{v}
}

func (o FiltersOption) Get() string {
	return o.value
}

func (o FiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *FiltersOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;attribute&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}