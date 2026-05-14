.PHONY: build test vet run clean

BIN_DIR  := bin
BINARY   := $(BIN_DIR)/stringsplit

build:
	go build -o $(BINARY) ./cmd/stringsplit/

test:
	go test -v ./...

vet:
	go vet ./...

run: build
	$(BINARY)

clean:
	rm -rf $(BIN_DIR)
