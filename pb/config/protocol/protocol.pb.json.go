// Code generated by protoc-gen-jsonify. DO NOT EDIT.
// source: config/protocol/protocol.proto

package protocol
import (
	"bytes"
	"encoding/json"
	"github.com/gogo/protobuf/jsonpb"
)

// MySQLOptionJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of MySQLOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var MySQLOptionJSONMarshaler = new(jsonpb.Marshaler)
// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *MySQLOption) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := MySQLOptionJSONMarshaler.Marshal(buf, m); err != nil {
	  return nil, err
	}
	return buf.Bytes(), nil
}
var _ json.Marshaler = (*MySQLOption)(nil)
// MySQLOptionJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of MySQLOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var MySQLOptionJSONUnmarshaler = new(jsonpb.Unmarshaler)
// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *MySQLOption) UnmarshalJSON(b []byte) error {
	return MySQLOptionJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}
var _ json.Unmarshaler = (*MySQLOption)(nil)

// TCPOptionJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of TCPOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var TCPOptionJSONMarshaler = new(jsonpb.Marshaler)
// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *TCPOption) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := TCPOptionJSONMarshaler.Marshal(buf, m); err != nil {
	  return nil, err
	}
	return buf.Bytes(), nil
}
var _ json.Marshaler = (*TCPOption)(nil)
// TCPOptionJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of TCPOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var TCPOptionJSONUnmarshaler = new(jsonpb.Unmarshaler)
// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *TCPOption) UnmarshalJSON(b []byte) error {
	return TCPOptionJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}
var _ json.Unmarshaler = (*TCPOption)(nil)

// RedisOptionJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of RedisOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var RedisOptionJSONMarshaler = new(jsonpb.Marshaler)
// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *RedisOption) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := RedisOptionJSONMarshaler.Marshal(buf, m); err != nil {
	  return nil, err
	}
	return buf.Bytes(), nil
}
var _ json.Marshaler = (*RedisOption)(nil)
// RedisOptionJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of RedisOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var RedisOptionJSONUnmarshaler = new(jsonpb.Unmarshaler)
// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *RedisOption) UnmarshalJSON(b []byte) error {
	return RedisOptionJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}
var _ json.Unmarshaler = (*RedisOption)(nil)

