package opt

import "encoding/json"

type MaxFacetHitsOption struct {
	value int
}

func MaxFacetHits(v int) MaxFacetHitsOption {
	return MaxFacetHitsOption{v}
}

func (o MaxFacetHitsOption) Get() int {
	return o.value
}

func (o MaxFacetHitsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *MaxFacetHitsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 10
		return nil
	}
	return json.Unmarshal(data, &o.value)
}