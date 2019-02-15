package opt

import "encoding/json"

type MaxValuesPerFacetOption struct {
	value int
}

func MaxValuesPerFacet(v int) MaxValuesPerFacetOption {
	return MaxValuesPerFacetOption{v}
}

func (o MaxValuesPerFacetOption) Get() int {
	return o.value
}

func (o MaxValuesPerFacetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *MaxValuesPerFacetOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 100
		return nil
	}
	return json.Unmarshal(data, &o.value)
}