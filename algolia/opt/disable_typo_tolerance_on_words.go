package opt

import "encoding/json"

type DisableTypoToleranceOnWordsOption struct {
	value []string
}

func DisableTypoToleranceOnWords(v []string) DisableTypoToleranceOnWordsOption {
	return DisableTypoToleranceOnWordsOption{v}
}

func (o DisableTypoToleranceOnWordsOption) Get() []string {
	return o.value
}

func (o DisableTypoToleranceOnWordsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *DisableTypoToleranceOnWordsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}