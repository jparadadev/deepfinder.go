# Build project
.PHONY = build
build:
	go build

# Run tests
.PHONY = test
test:
	go test ./tests -v