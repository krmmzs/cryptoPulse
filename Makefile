.PHONY: build install uninstall clean help

# Build targets
build:
	go build -o bin/btcq ./cmd/btcq

# Install binary to GOPATH/bin (or GOBIN if set)
install:
	go install ./cmd/btcq

# Uninstall binary from GOPATH/bin (or GOBIN if set)
uninstall:
	@GOBIN_PATH=$$(go env GOBIN); \
	if [ -z "$$GOBIN_PATH" ]; then \
		GOBIN_PATH=$$(go env GOPATH)/bin; \
	fi; \
	if [ -f "$$GOBIN_PATH/btcq" ]; then \
		rm "$$GOBIN_PATH/btcq" && echo "Removed $$GOBIN_PATH/btcq"; \
	else \
		echo "btcq not found in $$GOBIN_PATH"; \
	fi

# Clean build artifacts
clean:
	rm -rf bin/

# Show help
help:
	@echo "Available targets:"
	@echo "  build     - Build binary to bin/ directory"
	@echo "  install   - Install binary to GOPATH/bin using go install"
	@echo "  uninstall - Remove installed binary from GOPATH/bin"
	@echo "  clean     - Remove build artifacts"
	@echo "  help      - Show this help message"