.PHONY: proto

proto: build
	cd e2e && buf generate

build: 
	go build .
	