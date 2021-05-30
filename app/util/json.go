package util

import (
	"bytes"
	"encoding/json"
	"io"
)

// Converts the argument to a string containing pretty JSON, logging errors
func ToJSON(x interface{}) string {
	return string(ToJSONBytes(x, true))
}

// Converts the argument to a string containing compact JSON, logging errors
func ToJSONCompact(x interface{}) string {
	return string(ToJSONBytes(x, false))
}

// Converts the argument to an optionally indented byte array, logging errors
func ToJSONBytes(x interface{}, indent bool) []byte {
	if indent {
		b, _ := json.MarshalIndent(x, "", "  ")
		return b
	}
	b, _ := json.Marshal(x)
	return b
}

// Parses the provided JSON to the provided interface
func FromJSON(msg json.RawMessage, tgt interface{}) error {
	return json.Unmarshal(msg, tgt)
}

func FromJSONReader(r io.Reader, tgt interface{}) error {
	return json.NewDecoder(r).Decode(tgt)
}

// Parses the provided JSON to the provided interface, validating that all fields are used
func FromJSONStrict(msg json.RawMessage, tgt interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(msg))
	dec.DisallowUnknownFields()
	return dec.Decode(tgt)
}

// Parses the provided JSON as a string
func FromJSONString(msg json.RawMessage) (string, error) {
	tgt := ""
	err := json.Unmarshal(msg, &tgt)
	if err != nil {
		return "", err
	}
	return tgt, nil
}
