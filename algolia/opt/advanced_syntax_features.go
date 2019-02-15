package opt

import "encoding/json"

type AdvancedSyntaxFeaturesOption struct {
	value []string
}

func AdvancedSyntaxFeatures(v []string) AdvancedSyntaxFeaturesOption {
	return AdvancedSyntaxFeaturesOption{v}
}

func (o AdvancedSyntaxFeaturesOption) Get() []string {
	return o.value
}

func (o AdvancedSyntaxFeaturesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AdvancedSyntaxFeaturesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}