// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package e2e

import (
	"testing"
)

func TestInterfaceImpl(t *testing.T) {
	b := &Basic{A: "foo"}

	var b2 Basic
	buf, err := b.MarshalBinary()
	if err != nil {
		t.Fatalf("MarshalBinary errored: %v", err)
	}

	if err := b2.UnmarshalBinary(buf); err != nil {
		t.Fatalf("UnmarshalBinary errored: %v", err)
	}

	if b.A != b2.A {
		t.Fatalf("Binary Marshal + Unmarshal isnt lossless")
	}
}
