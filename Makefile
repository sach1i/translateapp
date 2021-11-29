BINARY := translateapp

.PHONY: test
test:
	@go test \
		-shuffle=on \
		-count=1 \
		-short \
		-timeout=5m \
		./...

.PHONY: build
build:
	@go build \
		-o=$(BINARY) \
		./cmd/translateapp

.PHONY: clean
clean:
	rm -rf $(BINARY)
