package opt

import "encoding/json"

type RestrictSearchableAttributesOption struct {
	value []string
}

func RestrictSearchableAttributes(v []string) RestrictSearchableAttributesOption {
	return RestrictSearchableAttributesOption{v}
}

func (o RestrictSearchableAttributesOption) Get() []string {
	return o.value
}

func (o RestrictSearchableAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RestrictSearchableAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}