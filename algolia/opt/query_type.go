package opt

import "encoding/json"

type QueryTypeOption struct {
	value string
}

func QueryType(v string) QueryTypeOption {
	return QueryTypeOption{v}
}

func (o QueryTypeOption) Get() string {
	return o.value
}

func (o QueryTypeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *QueryTypeOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;prefixLast&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}