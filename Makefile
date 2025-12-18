.PHONY: all ray rs clean

# Default target
all: rs

# Go configuration
GO := go
MOD := go.mod

ray:
	$(GO) run . ray

rs:
	$(GO) run . rs

# Build for release
build-ray:
	$(GO) build -o raytracer . && ./raytracer rs

clean:
	$(GO) clean
	rm -f raytracer

# Format and check
fmt:
	$(GO) fmt ./...

test:
	$(GO) test ./...
