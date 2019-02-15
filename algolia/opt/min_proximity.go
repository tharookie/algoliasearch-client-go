package opt

import "encoding/json"

type MinProximityOption struct {
	value int
}

func MinProximity(v int) MinProximityOption {
	return MinProximityOption{v}
}

func (o MinProximityOption) Get() int {
	return o.value
}

func (o MinProximityOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *MinProximityOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 1
		return nil
	}
	return json.Unmarshal(data, &o.value)
}