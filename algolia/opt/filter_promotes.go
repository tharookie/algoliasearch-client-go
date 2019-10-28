// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// FilterPromotesOption is a wrapper for an FilterPromotes option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type FilterPromotesOption struct {
	value bool
}

// FilterPromotes wraps the given value into a FilterPromotesOption.
func FilterPromotes(v bool) *FilterPromotesOption {
	return &FilterPromotesOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *FilterPromotesOption) Get() bool {
	if o == nil {
		return false
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// FilterPromotesOption.
func (o FilterPromotesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// FilterPromotesOption.
func (o *FilterPromotesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *FilterPromotesOption) Equal(o2 *FilterPromotesOption) bool {
	if o == nil {
		return o2 == nil || o2.value == false
	}
	if o2 == nil {
		return o == nil || o.value == false
	}
	return o.value == o2.value
}

// FilterPromotesEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func FilterPromotesEqual(o1, o2 *FilterPromotesOption) bool {
	return o1.Equal(o2)
}