BINARY   := netmapper
MAIN     := ./main.go
VERSION  := 1.0.0
LDFLAGS  := -ldflags "-X main.version=$(VERSION)"

.PHONY: build test install clean lint tidy

build:
	go build $(LDFLAGS) -o $(BINARY) $(MAIN)

test:
	go test ./...

install:
	go install $(LDFLAGS) ./...

clean:
	rm -f $(BINARY) $(BINARY).exe

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

run-simple: build
	./$(BINARY) analyze --config examples/simple.yaml

run-complex: build
	./$(BINARY) analyze --config examples/complex.yaml

run-invalid: build
	./$(BINARY) validate --config examples/invalid.yaml

.DEFAULT_GOAL := build
