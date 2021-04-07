default: build

.PHONY: build
build: 
	go build -o ./build/justgo
	@echo "Successfully built new binary under ./build"
