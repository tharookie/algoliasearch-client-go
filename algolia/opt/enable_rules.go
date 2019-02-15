package opt

import "encoding/json"

type EnableRulesOption struct {
	value bool
}

func EnableRules(v bool) EnableRulesOption {
	return EnableRulesOption{v}
}

func (o EnableRulesOption) Get() bool {
	return o.value
}

func (o EnableRulesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *EnableRulesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}