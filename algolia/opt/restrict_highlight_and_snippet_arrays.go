package opt

import "encoding/json"

type RestrictHighlightAndSnippetArraysOption struct {
	value bool
}

func RestrictHighlightAndSnippetArrays(v bool) RestrictHighlightAndSnippetArraysOption {
	return RestrictHighlightAndSnippetArraysOption{v}
}

func (o RestrictHighlightAndSnippetArraysOption) Get() bool {
	return o.value
}

func (o RestrictHighlightAndSnippetArraysOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RestrictHighlightAndSnippetArraysOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}