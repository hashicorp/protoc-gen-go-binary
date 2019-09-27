.PHONY: proto
proto:
	cd e2e && protoc --gofast_out=. --go-binary_out=logtostderr=true:. e2e.proto
