// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: e2e.proto

package e2e

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *Basic) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *Basic) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
