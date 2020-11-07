.PHONY: build
build:
	go build -v -o hammurabi ./cmd/hammurabi

.DEFAULT_GOAL: build