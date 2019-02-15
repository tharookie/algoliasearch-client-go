package opt

import "encoding/json"

type AroundLatLngOption struct {
	value string
}

func AroundLatLng(v string) AroundLatLngOption {
	return AroundLatLngOption{v}
}

func (o AroundLatLngOption) Get() string {
	return o.value
}

func (o AroundLatLngOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AroundLatLngOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}