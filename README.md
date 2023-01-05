# protoc-gen-go-binary

This is a plugin for the Google Protocol Buffers compiler
[`protoc`](https://github.com/protocolbuffers/protobuf) that generates
code to implement [`encoding.BinaryMarshaler`](https://golang.org/pkg/encoding/#BinaryMarshaler)
and [`encoding.BinaryUnmarshaler`](https://golang.org/pkg/encoding/#BinaryUnmarshaler)
by just calling the `Marshal` and `Unmarshal` functions already generated for the types.

This enables Go-generated protobuf messages to be used in situations where the code
already supports using the binary marshaling interfaces.

The code heavily relies on google.golang.org/protobuf/compiler/protogen and is mostly boilerplate. 

## Install

```
go get github.com/hashicorp/protoc-gen-go-binary
```

Also required:

- [buf](https://github.com/bufbuild/buf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

## Usage

Define your messages like normal:

```proto
syntax = "proto3";

message Request {
  oneof kind {
    string name = 1;
    int32  code = 2;
  }
}
```

The example message purposely uses a `oneof` since this won't work by
default with `encoding/json`. Next, generate the code:

```
protoc --go_out=. --go-binary_out=. request.proto
```

Your output should contain a file `request.pb.binary.go` which contains
the implementation of `encoding.BinaryMarshal/BinaryUnmarshal` for all your message types.
You can then encode binary encode your message as protobufs.

```go
import (
  "bytes"
  "encoding/gob"
)

var buf bytes.Buffer
encoder := gob.NewEncoder(&buf)

// Marshal
err := encoder.Encode(&Request{
  Kind: &Kind_Name{
    Name: "alice",
  },
}

// Unmarshal
var result Request
decoder := gob.NewDecoder(&buf)
err := decoder.Decode(&result)
```
