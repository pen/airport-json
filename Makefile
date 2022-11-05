lint:
	golangci-lint run

build: airport-json

airport-json: cmd/airport-json/*
	go build -v -ldflags '-X main.version=0.0.0-localbuild' -o $@ cmd/airport-json/main.go
